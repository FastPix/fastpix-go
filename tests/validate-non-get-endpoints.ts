#!/usr/bin/env tsx
/*
 * Non-GET endpoints validator (POST / PUT / PATCH / DELETE) using `openapi-response-validator`.
 *
 * Unlike the GET validator, these operations MUTATE live data, so we cannot hit
 * the raw API and the SDK separately (that would create/delete twice). Instead
 * this driver invokes the Go SDK once per operation (via tests/sdkharness) and:
 *  - captures the SDK's deserialized return value (for the diff + artifact)
 *  - captures the raw HTTP status + raw JSON body from the SDK's underlying
 *    response (for OpenAPI response-schema validation)
 *
 * No fixtures are required. The driver runs a create -> use -> delete lifecycle:
 *  1. CREATE phase  (POST)   - creates real resources, captures their IDs
 *  2. UPDATE phase  (PUT/PATCH) - exercises updates against the created IDs
 *  3. DELETE phase  (DELETE) - tears the resources down, LAST, so deletes only
 *     run after every POST/PUT/PATCH has completed.
 *
 * A step whose required IDs were never captured (because an upstream create
 * failed) is reported as SKIP rather than called with nulls.
 *
 * Output:
 *  - per-operation artifacts in `tests/artifacts-non-get/`
 *  - `tests/NON_GET_ENDPOINTS_VALIDATION_REPORT.md`
 *
 * Requirements:
 *  - FASTPIX_USERNAME / FASTPIX_PASSWORD env vars (Basic Auth)
 *  - optional FASTPIX_BASE_URL / FASTPIX_SERVER_URL (defaults to spec server)
 */

/// <reference path="./shims.d.ts" />

import { readFileSync, writeFileSync, existsSync, mkdirSync } from "node:fs";
import { spawnSync } from "node:child_process";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { createRequire } from "node:module";
import yaml from "js-yaml";

const require = createRequire(import.meta.url);
const openapiResponseValidatorMod = require("openapi-response-validator");
const OpenAPIResponseValidator =
  openapiResponseValidatorMod?.default ?? openapiResponseValidatorMod;

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const ARTIFACTS_DIRNAME = "artifacts-non-get";
const REPORT_MD = "NON_GET_ENDPOINTS_VALIDATION_REPORT.md";
const MAX_PREVIEW_CHARS = 4000;

type Phase = "CREATE" | "UPDATE" | "DELETE";

type EndpointInfo = {
  path: string;
  method: string;
  operationId: string;
  responses: Record<string, any>;
};

// Mutable context threaded through the lifecycle; populated by capture() hooks.
type Ctx = {
  signingKeyId?: string;
  playlistId?: string;
  streamId?: string;
  mediaId?: string;
  mediaPlaybackId?: string; // the media's default playback id (ready when media is Ready)
  createdPlaybackId?: string; // a playback id created via create-media-playback-id
  trackId?: string;
  streamPlaybackId?: string;
  simulcastId?: string;
  uploadId?: string;
};

type Step = {
  operationId: string;
  phase: Phase;
  // ctx keys that must be present, else the step is skipped
  needs?: (keyof Ctx)[];
  // build the per-call request (path params) from the current ctx
  request: (ctx: Ctx) => Record<string, any>;
  // extract a created id from the SDK response value into ctx
  capture?: (value: any, ctx: Ctx) => void;
  // if the call fails with an error containing this substring, wait and retry
  // (used to wait for an async resource — e.g. a playback id — to become ready)
  retryOn?: string;
};

type StepResult = {
  operationId: string;
  method: string;
  path: string;
  phase: Phase;
  status: "PASS" | "FAIL" | "SKIP";
  httpStatus: number | null;
  openapiValid: boolean | null;
  openapiErrors: any[];
  sdkOk: boolean;
  sdkError?: string;
  missingInSDK: string[];
  missingInAPI: string[];
  note?: string;
  capturedId?: string;
};

type GoSDKResult =
  | { ok: true; value: any; statusCode: number | null; rawBody: any }
  | { ok: false; error: { name?: string; message?: string; statusCode?: number; bodyJson?: any } };

function safeFileSlug(input: string): string {
  return input.replace(/[^a-zA-Z0-9_.-]+/g, "_");
}

function toPrettyJson(value: unknown): string {
  return JSON.stringify(value, null, 2);
}

function preview(text: string): string {
  return text.length > MAX_PREVIEW_CHARS ? `${text.slice(0, MAX_PREVIEW_CHARS)}\n... [truncated]` : text;
}

function basicAuthHeader(username: string, password: string): string {
  return "Basic " + Buffer.from(`${username}:${password}`).toString("base64");
}

const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

// A freshly-created media is "Preparing"; adding playback-ids / tracks to it
// returns 400 until it reaches "Ready". Poll the GET media endpoint so the
// dependent create steps have a usable resource.
async function waitForMediaReady(
  baseUrl: string,
  username: string,
  password: string,
  mediaId: string,
  timeoutMs = 180000,
  intervalMs = 5000,
): Promise<string> {
  const url = `${baseUrl.replace(/\/$/, "")}/on-demand/${mediaId}`;
  const deadline = Date.now() + timeoutMs;
  let last = "unknown";
  while (Date.now() < deadline) {
    try {
      const res = await fetch(url, { headers: { Accept: "application/json", Authorization: basicAuthHeader(username, password) } });
      const body: any = await res.json().catch(() => null);
      last = body?.data?.status ?? last;
      if (last === "Ready") return "Ready";
      if (last === "Errored" || last === "Failed") return last;
    } catch {
      /* transient; keep polling */
    }
    await sleep(intervalMs);
  }
  return last;
}

// ---------------------------------------------------------------------------
// Go SDK invocation (via the tests/sdkharness subprocess)
// ---------------------------------------------------------------------------

// Path to the Go harness package, relative to the SDK repo root (parent of tests/).
const GO_HARNESS_PKG = "./tests/sdkharness";

// Invoke the Go SDK by running the compiled harness as a subprocess, passing the
// operation + request as JSON on stdin. On success the harness returns the
// deserialized response body plus the raw wire status/body (captured by a
// wrapping HTTP client) so the response-schema validation stays faithful.
function invokeGoSDK(
  operationId: string,
  request: Record<string, any>,
  baseUrl: string,
  username: string,
  password: string,
): GoSDKResult {
  const repoRoot = join(__dirname, "..");

  const child = spawnSync("go", ["run", GO_HARNESS_PKG], {
    input: JSON.stringify({ operationId, request, baseUrl, username, password }),
    encoding: "utf-8",
    cwd: repoRoot,
    maxBuffer: 10 * 1024 * 1024,
  });

  if (child.error) {
    return { ok: false, error: { name: "GoSpawnError", message: child.error.message } };
  }

  const stdout = (child.stdout || "").trim();
  const stderr = (child.stderr || "").trim();
  if (stderr) console.error(`go stderr: ${stderr.split("\n").slice(0, 3).join(" ")}`);

  if (!stdout.startsWith("{") && !stdout.startsWith("[")) {
    return { ok: false, error: { name: "GoRuntimeError", message: (stderr || stdout).slice(0, 500) } };
  }

  try {
    const parsed = JSON.parse(stdout);
    if (parsed?.ok) {
      // The harness already returns rawBody as parsed JSON (object/array) or a
      // string when the body was not JSON; pass it through unchanged.
      return { ok: true, value: parsed.value, statusCode: parsed.statusCode ?? null, rawBody: parsed.rawBody ?? null };
    }
    return { ok: false, error: parsed?.error ?? { name: "GoSDKError", message: stdout.slice(0, 500) } };
  } catch (e: any) {
    return { ok: false, error: { name: "GoOutputParseError", message: `${e.message}: ${stdout.slice(0, 300)}` } };
  }
}

// A freshly-added track is fetched/processed asynchronously; generating
// subtitles before it exists returns 404 "track not found". Poll the media's
// track list until the track is present (and Ready when status is exposed).
async function waitForTrackReady(
  baseUrl: string,
  username: string,
  password: string,
  mediaId: string,
  trackId: string,
  timeoutMs = 180000,
  intervalMs = 5000,
): Promise<string> {
  const url = `${baseUrl.replace(/\/$/, "")}/on-demand/${mediaId}`;
  const deadline = Date.now() + timeoutMs;
  let last = "absent";
  while (Date.now() < deadline) {
    try {
      const res = await fetch(url, { headers: { Accept: "application/json", Authorization: basicAuthHeader(username, password) } });
      const body: any = await res.json().catch(() => null);
      const track = (body?.data?.tracks ?? []).find((t: any) => t?.id === trackId);
      if (track) {
        last = track.status ?? "present";
        if (last === "Ready" || last === "present") return last;
      }
    } catch {
      /* transient; keep polling */
    }
    await sleep(intervalMs);
  }
  return last;
}

// ---------------------------------------------------------------------------
// Spec + OpenAPI validation (shared with the GET validator)
// ---------------------------------------------------------------------------

function resolveSpecPath(): string {
  const candidates = [
    process.env.FASTPIX_OPENAPI_SPEC,
    join(__dirname, "../fixed 7.yaml"),
    join(__dirname, "../fastpix.yaml"),
    join(__dirname, "../fixed.yaml"),
    join(__dirname, "../fastpix-openapi.yaml"),
    join(__dirname, "../../fastpix-openapi.yaml"),
  ].filter((p): p is string => Boolean(p));
  for (const p of candidates) if (existsSync(p)) return p;
  throw new Error(`OpenAPI spec not found. Tried: ${candidates.join(", ")}`);
}

function loadOpenAPISpec(): any {
  return yaml.load(readFileSync(resolveSpecPath(), "utf-8"));
}

function extractNonGetEndpoints(spec: any): Map<string, EndpointInfo> {
  const out = new Map<string, EndpointInfo>();
  for (const [path, methods] of Object.entries(spec.paths || {})) {
    const m = methods as any;
    for (const method of ["post", "put", "patch", "delete"]) {
      if (!m[method]) continue;
      out.set(m[method].operationId, {
        path,
        method: method.toUpperCase(),
        operationId: m[method].operationId,
        responses: m[method].responses || {},
      });
    }
  }
  return out;
}

function convertRefsToDefinitions(node: any): any {
  if (node == null || typeof node !== "object") return node;
  if (Array.isArray(node)) return node.map(convertRefsToDefinitions);
  const out: any = {};
  for (const [k, v] of Object.entries(node)) {
    if (k === "$ref" && typeof v === "string") out[k] = v.replace("#/components/schemas/", "#/definitions/");
    else out[k] = convertRefsToDefinitions(v);
  }
  return out;
}

function makeOpenAPIResponseValidator(spec: any, endpoint: EndpointInfo) {
  const definitions = convertRefsToDefinitions(spec.components?.schemas || {});
  const responses: any = {};
  for (const [status, def] of Object.entries(endpoint.responses || {})) {
    const d = def as any;
    const schema = d?.content?.["application/json"]?.schema;
    if (!schema) continue;
    responses[status] = { description: d.description || "", schema: convertRefsToDefinitions(schema) };
  }
  if (Object.keys(responses).length === 0) return null;
  return new OpenAPIResponseValidator({ responses, definitions });
}

// ---------------------------------------------------------------------------
// JSON diff helpers (shared with the GET validator)
// ---------------------------------------------------------------------------

function collectJsonPaths(value: any, prefix = "", opts: { includeEmptyArrays?: boolean } = {}): Set<string> {
  const out = new Set<string>();
  const includeEmptyArrays = opts.includeEmptyArrays ?? true;
  if (value === null || value === undefined) return out;
  if (typeof value !== "object") {
    if (prefix) out.add(prefix);
    return out;
  }
  if (Array.isArray(value)) {
    if (!includeEmptyArrays && value.length === 0) return out;
    const arrayPrefix = prefix ? `${prefix}[]` : "[]";
    out.add(arrayPrefix);
    for (const item of value) for (const p of collectJsonPaths(item, arrayPrefix, opts)) out.add(p);
    return out;
  }
  for (const [k, v] of Object.entries(value)) {
    if (!includeEmptyArrays && Array.isArray(v) && v.length === 0) continue;
    if (!includeEmptyArrays && (v === null || v === undefined)) continue;
    if (!includeEmptyArrays && typeof v === "object" && v !== null && !Array.isArray(v) && Object.keys(v).length === 0) continue;
    const p = prefix ? `${prefix}.${k}` : k;
    out.add(p);
    for (const child of collectJsonPaths(v, p, opts)) out.add(child);
  }
  return out;
}

function canonicalizeKey(key: string): string {
  const camel = key.includes("_")
    ? key.toLowerCase().replace(/_([a-z0-9])/g, (_, c) => String(c).toUpperCase())
    : key;
  return camel.replaceAll("SDK", "Sdk").replaceAll("API", "Api");
}

function normalizeJsonForComparison(value: any): any {
  if (value === null || value === undefined) return value;
  if (Array.isArray(value)) return value.map(normalizeJsonForComparison);
  if (typeof value !== "object") return value;
  const out: any = {};
  for (const [k, v] of Object.entries(value)) out[canonicalizeKey(k)] = normalizeJsonForComparison(v);
  return out;
}

function sortUnique(arr: string[]) {
  return Array.from(new Set(arr)).sort((a, b) => a.localeCompare(b));
}

function jsonRoundTrip(value: any): any {
  return JSON.parse(JSON.stringify(value));
}

// ---------------------------------------------------------------------------
// Lifecycle definition: ordered so all DELETEs run after every POST/PUT/PATCH.
// ---------------------------------------------------------------------------

const STEPS: Step[] = [
  // ---- CREATE ----
  { operationId: "create_signing_key", phase: "CREATE", request: () => ({}), capture: (v, c) => { c.signingKeyId = v?.data?.id; } },
  { operationId: "create-a-playlist", phase: "CREATE", request: () => ({}), capture: (v, c) => { c.playlistId = v?.data?.id; } },
  { operationId: "create-new-stream", phase: "CREATE", request: () => ({}), capture: (v, c) => { c.streamId = v?.data?.streamId ?? v?.data?.id; } },
  { operationId: "create-media", phase: "CREATE", request: () => ({}), capture: (v, c) => { c.mediaId = v?.data?.id; c.mediaPlaybackId = v?.data?.playbackIds?.[0]?.id; } },
  { operationId: "create-media-playback-id", phase: "CREATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }), capture: (v, c) => { c.createdPlaybackId = v?.data?.playbackIds?.[0]?.id ?? v?.data?.id; } },
  { operationId: "Add-media-track", phase: "CREATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }), capture: (v, c) => { c.trackId = v?.data?.id; } },
  { operationId: "create-playbackId-of-stream", phase: "CREATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }), capture: (v, c) => { c.streamPlaybackId = v?.data?.playbackIds?.[0]?.id ?? v?.data?.id; } },
  { operationId: "create-simulcast-of-stream", phase: "CREATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }), capture: (v, c) => { c.simulcastId = v?.data?.simulcastId ?? v?.data?.id; } },
  { operationId: "direct-upload-video-media", phase: "CREATE", request: () => ({}), capture: (v, c) => { c.uploadId = v?.data?.uploadId ?? v?.data?.id; } },
  { operationId: "Generate-subtitle-track", phase: "CREATE", needs: ["mediaId", "trackId"], request: (c) => ({ mediaId: c.mediaId, trackId: c.trackId }) },

  // ---- UPDATE (PUT/PATCH) ----
  { operationId: "updated-media", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "updated-source-access", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "updated-mp4Support", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "update-media-summary", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "update-media-chapters", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "update-media-named-entities", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "update-media-moderation", phase: "UPDATE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "update-media-track", phase: "UPDATE", needs: ["mediaId", "trackId"], request: (c) => ({ mediaId: c.mediaId, trackId: c.trackId }) },
  { operationId: "update-domain-restrictions", phase: "UPDATE", needs: ["mediaId", "mediaPlaybackId"], retryOn: "not ready for updates", request: (c) => ({ mediaId: c.mediaId, playbackId: c.mediaPlaybackId }) },
  { operationId: "update-user-agent-restrictions", phase: "UPDATE", needs: ["mediaId", "mediaPlaybackId"], retryOn: "not ready for updates", request: (c) => ({ mediaId: c.mediaId, playbackId: c.mediaPlaybackId }) },
  { operationId: "update-a-playlist", phase: "UPDATE", needs: ["playlistId"], request: (c) => ({ playlistId: c.playlistId }) },
  { operationId: "add-media-to-playlist", phase: "UPDATE", needs: ["playlistId", "mediaId"], request: (c) => ({ playlistId: c.playlistId, mediaId: c.mediaId }) },
  { operationId: "change-media-order-in-playlist", phase: "UPDATE", needs: ["playlistId", "mediaId"], request: (c) => ({ playlistId: c.playlistId, mediaId: c.mediaId }) },
  { operationId: "update-live-stream", phase: "UPDATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }) },
  { operationId: "update-specific-simulcast-of-stream", phase: "UPDATE", needs: ["streamId", "simulcastId"], request: (c) => ({ streamId: c.streamId, simulcastId: c.simulcastId }) },
  // a freshly-created stream is already enabled, so disable first, then enable.
  { operationId: "disable-live-stream", phase: "UPDATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }) },
  { operationId: "enable-live-stream", phase: "UPDATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }) },
  // complete requires an actively-streaming encoder; with no ingest it is
  // expected to fail (the one allowed failure in a credentials-only run).
  { operationId: "complete-live-stream", phase: "UPDATE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }) },
  { operationId: "cancel-upload", phase: "UPDATE", needs: ["uploadId"], request: (c) => ({ uploadId: c.uploadId }) },

  // ---- DELETE (last) ----
  { operationId: "delete-media-from-playlist", phase: "DELETE", needs: ["playlistId", "mediaId"], request: (c) => ({ playlistId: c.playlistId, mediaId: c.mediaId }) },
  { operationId: "delete-a-playlist", phase: "DELETE", needs: ["playlistId"], request: (c) => ({ playlistId: c.playlistId }) },
  { operationId: "delete-media-track", phase: "DELETE", needs: ["mediaId", "trackId"], request: (c) => ({ mediaId: c.mediaId, trackId: c.trackId }) },
  { operationId: "delete-media-playback-id", phase: "DELETE", needs: ["mediaId", "createdPlaybackId"], request: (c) => ({ mediaId: c.mediaId, playbackId: c.createdPlaybackId }) },
  { operationId: "delete-simulcast-of-stream", phase: "DELETE", needs: ["streamId", "simulcastId"], request: (c) => ({ streamId: c.streamId, simulcastId: c.simulcastId }) },
  { operationId: "delete-playbackId-of-stream", phase: "DELETE", needs: ["streamId", "streamPlaybackId"], request: (c) => ({ streamId: c.streamId, playbackId: c.streamPlaybackId }) },
  { operationId: "delete-live-stream", phase: "DELETE", needs: ["streamId"], request: (c) => ({ streamId: c.streamId }) },
  { operationId: "delete-media", phase: "DELETE", needs: ["mediaId"], request: (c) => ({ mediaId: c.mediaId }) },
  { operationId: "delete_signing_key", phase: "DELETE", needs: ["signingKeyId"], request: (c) => ({ signingKeyId: c.signingKeyId }) },
];

// ---------------------------------------------------------------------------
// Artifacts + report
// ---------------------------------------------------------------------------

function writeArtifacts(operationId: string, rawBody: any, sdkValue: any) {
  const dir = join(__dirname, ARTIFACTS_DIRNAME);
  mkdirSync(dir, { recursive: true });
  const slug = safeFileSlug(operationId);
  writeFileSync(join(dir, `${slug}.raw.json`), toPrettyJson(rawBody ?? null));
  writeFileSync(join(dir, `${slug}.sdk.json`), toPrettyJson(sdkValue ?? null));
}

function writeReport(results: StepResult[], ctx: Ctx) {
  const total = results.length;
  const pass = results.filter((r) => r.status === "PASS").length;
  const fail = results.filter((r) => r.status === "FAIL").length;
  const skip = results.filter((r) => r.status === "SKIP").length;

  const lines: string[] = [];
  lines.push("# Non-GET endpoints validation report\n");
  lines.push(`Generated: ${new Date().toISOString()}\n`);
  lines.push("## Summary\n");
  lines.push(`- **Total**: ${total}`);
  lines.push(`- **PASS**: ${pass}`);
  lines.push(`- **FAIL**: ${fail}`);
  lines.push(`- **SKIP**: ${skip}\n`);

  lines.push("## Captured resources\n");
  for (const [k, v] of Object.entries(ctx)) lines.push(`- \`${k}\`: ${v ?? "(not created)"}`);
  lines.push("");

  lines.push("## Consolidated report\n");
  lines.push("| Phase | Method | OperationId | HTTP | OpenAPI valid | SDK | Missing in SDK | Missing in API | Status |");
  lines.push("|---|---|---|---:|:--:|:--:|---|---|:--:|");
  const phaseOrder: Phase[] = ["CREATE", "UPDATE", "DELETE"];
  for (const phase of phaseOrder) {
    for (const r of results.filter((x) => x.phase === phase)) {
      const ov = r.openapiValid === null ? "—" : r.openapiValid ? "✅" : "❌";
      const sdk = r.status === "SKIP" ? "—" : r.sdkOk ? "✅" : "❌";
      const mis = (a: string[]) => (a.length ? a.join(", ") : "None");
      const st = r.status === "PASS" ? "✅ PASS" : r.status === "SKIP" ? "⤳ SKIP" : "❌ FAIL";
      lines.push(`| ${r.phase} | ${r.method} | \`${r.operationId}\` | ${r.httpStatus ?? "—"} | ${ov} | ${sdk} | ${mis(r.missingInSDK)} | ${mis(r.missingInAPI)} | ${st} |`);
    }
  }
  lines.push("");

  lines.push("## Per-operation details\n");
  for (const r of results) {
    lines.push(`### ${r.operationId} (\`${r.method} ${r.path}\`)`);
    lines.push(`- **Phase**: ${r.phase}`);
    lines.push(`- **Status**: ${r.status}`);
    if (r.httpStatus !== null) lines.push(`- **HTTP status**: ${r.httpStatus}`);
    if (r.capturedId) lines.push(`- **Captured id**: \`${r.capturedId}\``);
    if (r.note) lines.push(`- **Note**: ${r.note}`);
    if (r.sdkError) lines.push(`- **SDK error**: ${preview(r.sdkError)}`);
    if (r.openapiErrors.length) {
      lines.push(`- **OpenAPI errors**:`);
      for (const e of r.openapiErrors.slice(0, 20)) lines.push(`  - \`${e.path ?? ""}\` ${e.message ?? JSON.stringify(e)}`);
    }
    if (r.missingInSDK.length) { lines.push(`- **Missing in SDK (present in API)**:`); for (const p of r.missingInSDK) lines.push(`  - \`${p}\``); }
    if (r.missingInAPI.length) { lines.push(`- **Missing in API (present in SDK)**:`); for (const p of r.missingInAPI) lines.push(`  - \`${p}\``); }
    lines.push("");
  }

  const reportPath = join(__dirname, REPORT_MD);
  writeFileSync(reportPath, lines.join("\n"));
  console.log(`Report generated: ${reportPath}`);
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

async function main(): Promise<void> {
  const spec = loadOpenAPISpec();
  const endpoints = extractNonGetEndpoints(spec);

  const baseUrl: string =
    process.env.FASTPIX_BASE_URL
    ?? process.env.FASTPIX_SERVER_URL
    ?? ((spec.servers?.[0]?.url as string | undefined) ?? "https://api.fastpix.com/v1/");

  const username = process.env.FASTPIX_USERNAME ?? "";
  const password = process.env.FASTPIX_PASSWORD ?? "";
  if (!username || !password) {
    throw new Error("Set FASTPIX_USERNAME and FASTPIX_PASSWORD env vars for BasicAuth (use real credentials for live API validation)");
  }

  const ctx: Ctx = {};
  const results: StepResult[] = [];

  for (let i = 0; i < STEPS.length; i++) {
    const step = STEPS[i];
    const ep = endpoints.get(step.operationId);
    const base = {
      operationId: step.operationId,
      method: ep?.method ?? "?",
      path: ep?.path ?? "?",
      phase: step.phase,
      openapiErrors: [] as any[],
      missingInSDK: [] as string[],
      missingInAPI: [] as string[],
    };

    console.log(`[${i + 1}/${STEPS.length}] (${step.phase}) ${step.operationId}`);

    if (!ep) {
      results.push({ ...base, status: "SKIP", httpStatus: null, openapiValid: null, sdkOk: false, note: "operationId not found in spec" });
      continue;
    }

    const missingDeps = (step.needs ?? []).filter((k) => !ctx[k]);
    if (missingDeps.length) {
      console.log(`  ⤳ SKIP (missing: ${missingDeps.join(", ")})`);
      results.push({ ...base, status: "SKIP", httpStatus: null, openapiValid: null, sdkOk: false, note: `missing dependency: ${missingDeps.join(", ")}` });
      continue;
    }

    // generating subtitles needs the just-added track to be fetched/ready first
    if (step.operationId === "Generate-subtitle-track" && ctx.mediaId && ctx.trackId) {
      process.stdout.write(`  ⏳ waiting for track ${ctx.trackId} to be ready...`);
      const tstatus = await waitForTrackReady(baseUrl, username, password, ctx.mediaId, ctx.trackId);
      console.log(` ${tstatus}`);
    }

    const request = step.request(ctx);
    let sdkRes = invokeGoSDK(step.operationId, request, baseUrl, username, password);

    // wait for an async-provisioning resource (e.g. a playback id transitioning
    // from "preparing" to "available") by retrying while the error still matches.
    if (step.retryOn) {
      let attempt = 0;
      const maxAttempts = 24; // ~2 min at 5s
      while (!sdkRes.ok && attempt < maxAttempts && JSON.stringify(sdkRes.error ?? {}).includes(step.retryOn)) {
        attempt++;
        if (attempt === 1) process.stdout.write(`  ⏳ resource not ready, retrying`);
        else process.stdout.write(".");
        await sleep(5000);
        sdkRes = invokeGoSDK(step.operationId, request, baseUrl, username, password);
      }
      if (attempt > 0) console.log("");
    }

    if (!sdkRes.ok) {
      const msg = `${sdkRes.error?.name ?? "Error"}: ${sdkRes.error?.message ?? "SDK call failed"}`;
      console.log(`  ❌ FAIL — ${msg.split("\n")[0].slice(0, 120)}`);
      writeArtifacts(step.operationId, sdkRes.error?.bodyJson ?? null, sdkRes.error ?? null);
      results.push({ ...base, status: "FAIL", httpStatus: sdkRes.error?.statusCode ?? null, openapiValid: null, sdkOk: false, sdkError: msg });
      continue;
    }

    // capture created ids for downstream steps
    let capturedId: string | undefined;
    if (step.capture) {
      step.capture(sdkRes.value, ctx);
    }

    // a just-created media must reach "Ready" before playback-ids / tracks can
    // be added, otherwise those create steps 400 and cascade into SKIPs.
    if (step.operationId === "create-media" && ctx.mediaId) {
      process.stdout.write(`  ⏳ waiting for media ${ctx.mediaId} to be Ready...`);
      const status = await waitForMediaReady(baseUrl, username, password, ctx.mediaId);
      console.log(` ${status}`);
    }
    // best-effort: surface whatever id this step just stored
    capturedId =
      ctx.signingKeyId && step.operationId === "create_signing_key" ? ctx.signingKeyId :
      ctx.playlistId && step.operationId === "create-a-playlist" ? ctx.playlistId :
      ctx.streamId && step.operationId === "create-new-stream" ? ctx.streamId :
      ctx.mediaId && step.operationId === "create-media" ? ctx.mediaId :
      ctx.createdPlaybackId && step.operationId === "create-media-playback-id" ? ctx.createdPlaybackId :
      ctx.trackId && step.operationId === "Add-media-track" ? ctx.trackId :
      ctx.streamPlaybackId && step.operationId === "create-playbackId-of-stream" ? ctx.streamPlaybackId :
      ctx.simulcastId && step.operationId === "create-simulcast-of-stream" ? ctx.simulcastId :
      ctx.uploadId && step.operationId === "direct-upload-video-media" ? ctx.uploadId :
      undefined;

    // OpenAPI response-schema validation against the raw wire body
    const validator = makeOpenAPIResponseValidator(spec, ep);
    let openapiValid: boolean | null = null;
    let openapiErrors: any[] = [];
    if (validator && sdkRes.statusCode) {
      const err = validator.validateResponse(String(sdkRes.statusCode), sdkRes.rawBody);
      openapiValid = !err;
      openapiErrors = err?.errors ?? [];
    }

    // path diff between raw API body and SDK value
    const apiNorm = normalizeJsonForComparison(sdkRes.rawBody);
    const sdkNorm = sdkRes.value && typeof sdkRes.value === "object" ? normalizeJsonForComparison(jsonRoundTrip(sdkRes.value)) : null;
    const apiPaths = collectJsonPaths(apiNorm, "", { includeEmptyArrays: false });
    const sdkPaths = sdkNorm ? collectJsonPaths(sdkNorm, "", { includeEmptyArrays: false }) : new Set<string>();
    const missingInSDK = sdkPaths.size ? sortUnique([...apiPaths].filter((p) => !sdkPaths.has(p))) : [];
    const missingInAPI = sdkPaths.size ? sortUnique([...sdkPaths].filter((p) => !apiPaths.has(p))) : [];

    writeArtifacts(step.operationId, sdkRes.rawBody, sdkRes.value);

    const status: StepResult["status"] =
      sdkRes.ok && (openapiValid === null || openapiValid) && missingInSDK.length === 0 && missingInAPI.length === 0
        ? "PASS"
        : "FAIL";

    console.log(`  ${status === "PASS" ? "✅ PASS" : "❌ FAIL"} (HTTP ${sdkRes.statusCode ?? "?"})${capturedId ? ` id=${capturedId}` : ""}`);

    results.push({
      ...base,
      status,
      httpStatus: sdkRes.statusCode,
      openapiValid,
      openapiErrors,
      sdkOk: true,
      missingInSDK,
      missingInAPI,
      capturedId,
    });
  }

  writeReport(results, ctx);

  const pass = results.filter((r) => r.status === "PASS").length;
  const fail = results.filter((r) => r.status === "FAIL").length;
  const skip = results.filter((r) => r.status === "SKIP").length;
  console.log(`Summary: total=${results.length} pass=${pass} fail=${fail} skip=${skip}`);
}

await main();
