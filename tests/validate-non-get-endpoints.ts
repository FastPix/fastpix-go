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

function sortKeysDeep(value: unknown): unknown {
  if (Array.isArray(value)) return value.map(sortKeysDeep);
  if (value && typeof value === "object") {
    return Object.keys(value as Record<string, unknown>)
      .sort((a, b) => a.localeCompare(b))
      .reduce((acc, key) => {
        acc[key] = sortKeysDeep((value as Record<string, unknown>)[key]);
        return acc;
      }, {} as Record<string, unknown>);
  }
  return value;
}

function toPrettyJson(value: unknown): string {
  return JSON.stringify(sortKeysDeep(value), null, 2);
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


// Resolve the absolute path to the go binary from fixed, known locations.
// Using an absolute path in spawnSync avoids PATH-based executable lookup entirely,
// fully resolving Sonar S4036.
function resolveAbsoluteGoBinary(): string {
  const candidates = process.platform === "win32"
    ? [
        String.raw`C:\Go\bin\go.exe`,
        String.raw`C:\Program Files\Go\bin\go.exe`,
      ]
    : [
        "/opt/homebrew/bin/go",        // Homebrew on Apple Silicon Mac
        "/usr/local/go/bin/go",        // Standard Go install on Linux/macOS
        "/usr/bin/go",
        "/usr/local/bin/go",
        "/home/linuxbrew/.linuxbrew/bin/go", // Homebrew on Linux
      ];
  for (const p of candidates) {
    if (existsSync(p)) return p;
  }
  return "/opt/homebrew/bin/go";
}

const GO_BINARY = resolveAbsoluteGoBinary();

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
  mkdirSync(join(repoRoot, ".gotmp"), { recursive: true });

  const child = spawnSync(GO_BINARY, ["run", GO_HARNESS_PKG], {
    input: JSON.stringify({ operationId, request, baseUrl, username, password }),
    encoding: "utf-8",
    cwd: repoRoot,
    maxBuffer: 10 * 1024 * 1024,
    env: {
      HOME: process.env.HOME ?? "",
      GOPATH: process.env.GOPATH ?? "",
      GOROOT: process.env.GOROOT ?? "",
      GOCACHE: process.env.GOCACHE ?? "",
      GOTMPDIR: join(repoRoot, ".gotmp"),
      USERPROFILE: process.env.USERPROFILE ?? "",
    },
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
        if (last === "Ready" || last === "present" || last === "available") return last;
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
    const schema = def?.content?.["application/json"]?.schema;
    if (!schema) continue;
    responses[status] = { description: def.description || "", schema: convertRefsToDefinitions(schema) };
  }
  if (Object.keys(responses).length === 0) return null;
  return new OpenAPIResponseValidator({ responses, definitions });
}

// ---------------------------------------------------------------------------
// JSON diff helpers (shared with the GET validator)
// ---------------------------------------------------------------------------

function shouldSkipValue(v: any, includeEmptyArrays: boolean): boolean {
  if (includeEmptyArrays) return false;
  if (Array.isArray(v) && v.length === 0) return true;
  if (v === null || v === undefined) return true;
  return typeof v === "object" && !Array.isArray(v) && Object.keys(v).length === 0;
}

function collectJsonPathsFromArray(arr: any[], prefix: string, opts: { includeEmptyArrays?: boolean }, out: Set<string>): void {
  const includeEmptyArrays = opts.includeEmptyArrays ?? true;
  if (!includeEmptyArrays && arr.length === 0) return;
  const arrayPrefix = prefix ? `${prefix}[]` : "[]";
  out.add(arrayPrefix);
  for (const item of arr) for (const p of collectJsonPaths(item, arrayPrefix, opts)) out.add(p);
}

function collectJsonPaths(value: any, prefix = "", opts: { includeEmptyArrays?: boolean } = {}): Set<string> {
  const out = new Set<string>();
  const includeEmptyArrays = opts.includeEmptyArrays ?? true;
  if (value === null || value === undefined) return out;
  if (typeof value !== "object") {
    if (prefix) out.add(prefix);
    return out;
  }
  if (Array.isArray(value)) {
    collectJsonPathsFromArray(value, prefix, opts, out);
    return out;
  }
  for (const [k, v] of Object.entries(value)) {
    if (shouldSkipValue(v, includeEmptyArrays)) continue;
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
  return structuredClone(value);
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

function pushStepDetail(lines: string[], r: StepResult): void {
  lines.push(
    `### ${r.operationId} (\`${r.method} ${r.path}\`)`,
    `- **Phase**: ${r.phase}`,
    `- **Status**: ${r.status}`,
  );
  if (r.httpStatus !== null) lines.push(`- **HTTP status**: ${r.httpStatus}`);
  if (r.capturedId) lines.push(`- **Captured id**: \`${r.capturedId}\``);
  if (r.note) lines.push(`- **Note**: ${r.note}`);
  if (r.sdkError) lines.push(`- **SDK error**: ${preview(r.sdkError)}`);
  if (r.openapiErrors.length) {
    lines.push("- **OpenAPI errors**:");
    for (const e of r.openapiErrors.slice(0, 20)) lines.push(`  - \`${e.path ?? ""}\` ${e.message ?? JSON.stringify(e)}`);
  }
  if (r.missingInSDK.length) {
    lines.push("- **Missing in SDK (present in API)**:");
    for (const p of r.missingInSDK) lines.push(`  - \`${p}\``);
  }
  if (r.missingInAPI.length) {
    lines.push("- **Missing in API (present in SDK)**:");
    for (const p of r.missingInAPI) lines.push(`  - \`${p}\``);
  }
  lines.push("");
}

function formatOpenAPICol(openapiValid: boolean | null): string {
  if (openapiValid === null) return "—";
  return openapiValid ? "✅" : "❌";
}

function formatSDKCol(status: StepResult["status"], sdkOk: boolean): string {
  if (status === "SKIP") return "—";
  return sdkOk ? "✅" : "❌";
}

function formatStatus(status: StepResult["status"]): string {
  if (status === "PASS") return "✅ PASS";
  if (status === "SKIP") return "⤳ SKIP";
  return "❌ FAIL";
}

function formatPaths(paths: string[]): string {
  return paths.length ? paths.join(", ") : "None";
}

function buildTableRow(r: StepResult): string {
  const ov = formatOpenAPICol(r.openapiValid);
  const sdk = formatSDKCol(r.status, r.sdkOk);
  const st = formatStatus(r.status);
  return `| ${r.phase} | ${r.method} | \`${r.operationId}\` | ${r.httpStatus ?? "—"} | ${ov} | ${sdk} | ${formatPaths(r.missingInSDK)} | ${formatPaths(r.missingInAPI)} | ${st} |`;
}

function writeReport(results: StepResult[], ctx: Ctx) {
  const total = results.length;
  const pass = results.filter((r) => r.status === "PASS").length;
  const fail = results.filter((r) => r.status === "FAIL").length;
  const skip = results.filter((r) => r.status === "SKIP").length;

  const lines: string[] = [];
  lines.push(
    "# Non-GET endpoints validation report\n",
    `Generated: ${new Date().toISOString()}\n`,
    "## Summary\n",
    `- **Total**: ${total}`,
    `- **PASS**: ${pass}`,
    `- **FAIL**: ${fail}`,
    `- **SKIP**: ${skip}\n`,
    "## Captured resources\n",
  );
  for (const [k, v] of Object.entries(ctx)) lines.push(`- \`${k}\`: ${v ?? "(not created)"}`);
  lines.push(
    "",
    "## Consolidated report\n",
    "| Phase | Method | OperationId | HTTP | OpenAPI valid | SDK | Missing in SDK | Missing in API | Status |",
    "|---|---|---|---:|:--:|:--:|---|---|:--:|",
  );
  const phaseOrder: Phase[] = ["CREATE", "UPDATE", "DELETE"];
  for (const phase of phaseOrder) {
    for (const r of results.filter((x) => x.phase === phase)) {
      lines.push(buildTableRow(r));
    }
  }
  lines.push("", "## Per-operation details\n");
  for (const r of results) {
    pushStepDetail(lines, r);
  }

  const reportPath = join(__dirname, REPORT_MD);
  writeFileSync(reportPath, lines.join("\n"));
  console.log(`Report generated: ${reportPath}`);
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

type RunStepDeps = { spec: any; baseUrl: string; username: string; password: string };

async function waitIfNeeded(step: Step, ctx: Ctx, baseUrl: string, username: string, password: string): Promise<boolean> {
  if (step.operationId === "Generate-subtitle-track" && ctx.mediaId && ctx.trackId) {
    process.stdout.write(`  ⏳ waiting for track ${ctx.trackId} to be ready...`);
    const tstatus = await waitForTrackReady(baseUrl, username, password, ctx.mediaId, ctx.trackId);
    console.log(` ${tstatus}`);
    if (tstatus !== "Ready" && tstatus !== "present" && tstatus !== "available") {
      console.log(`  ⤳ SKIP — track not ready after timeout (status: ${tstatus})`);
      return false;
    }
  }
  return true;
}

async function retryLoop(
  step: Step,
  request: Record<string, any>,
  baseUrl: string,
  username: string,
  password: string,
): Promise<GoSDKResult> {
  let sdkRes = invokeGoSDK(step.operationId, request, baseUrl, username, password);
  if (!step.retryOn) return sdkRes;
  let attempt = 0;
  const maxAttempts = 24;
  while (!sdkRes.ok && attempt < maxAttempts && JSON.stringify(sdkRes.error ?? {}).includes(step.retryOn)) {
    attempt++;
    if (attempt === 1) process.stdout.write(`  ⏳ resource not ready, retrying`);
    else process.stdout.write(".");
    await sleep(5000);
    sdkRes = invokeGoSDK(step.operationId, request, baseUrl, username, password);
  }
  if (attempt > 0) console.log("");
  return sdkRes;
}

async function waitForMediaIfNeeded(step: Step, ctx: Ctx, baseUrl: string, username: string, password: string): Promise<void> {
  if (step.operationId === "create-media" && ctx.mediaId) {
    process.stdout.write(`  ⏳ waiting for media ${ctx.mediaId} to be Ready...`);
    const status = await waitForMediaReady(baseUrl, username, password, ctx.mediaId);
    console.log(` ${status}`);
  }
}

function getCapturedId(operationId: string, ctx: Ctx): string | undefined {
  const captureMap: Record<string, string | undefined> = {
    "create_signing_key": ctx.signingKeyId,
    "create-a-playlist": ctx.playlistId,
    "create-new-stream": ctx.streamId,
    "create-media": ctx.mediaId,
    "create-media-playback-id": ctx.createdPlaybackId,
    "Add-media-track": ctx.trackId,
    "create-playbackId-of-stream": ctx.streamPlaybackId,
    "create-simulcast-of-stream": ctx.simulcastId,
    "direct-upload-video-media": ctx.uploadId,
  };
  return captureMap[operationId];
}

function validateStepResponse(
  spec: any,
  ep: EndpointInfo,
  sdkRes: Extract<GoSDKResult, { ok: true }>,
): { openapiValid: boolean | null; openapiErrors: any[] } {
  const validator = makeOpenAPIResponseValidator(spec, ep);
  if (!validator || !sdkRes.statusCode) return { openapiValid: null, openapiErrors: [] };
  const err = validator.validateResponse(String(sdkRes.statusCode), sdkRes.rawBody);
  return { openapiValid: !err, openapiErrors: err?.errors ?? [] };
}

function computeStepDiff(sdkRes: Extract<GoSDKResult, { ok: true }>): { missingInSDK: string[]; missingInAPI: string[] } {
  const apiNorm = normalizeJsonForComparison(sdkRes.rawBody);
  const sdkNorm = sdkRes.value && typeof sdkRes.value === "object"
    ? normalizeJsonForComparison(jsonRoundTrip(sdkRes.value))
    : null;
  const apiPaths = collectJsonPaths(apiNorm, "", { includeEmptyArrays: false });
  const sdkPaths = sdkNorm ? collectJsonPaths(sdkNorm, "", { includeEmptyArrays: false }) : new Set<string>();
  return {
    missingInSDK: sdkPaths.size ? sortUnique([...apiPaths].filter((p) => !sdkPaths.has(p))) : [],
    missingInAPI: sdkPaths.size ? sortUnique([...sdkPaths].filter((p) => !apiPaths.has(p))) : [],
  };
}

async function runStep(
  step: Step,
  ep: EndpointInfo | undefined,
  ctx: Ctx,
  deps: RunStepDeps,
): Promise<StepResult> {
  const { spec, baseUrl, username, password } = deps;
  const base = {
    operationId: step.operationId,
    method: ep?.method ?? "?",
    path: ep?.path ?? "?",
    phase: step.phase,
    openapiErrors: [] as any[],
    missingInSDK: [] as string[],
    missingInAPI: [] as string[],
  };

  if (!ep) {
    return { ...base, status: "SKIP", httpStatus: null, openapiValid: null, sdkOk: false, note: "operationId not found in spec" };
  }

  const missingDeps = (step.needs ?? []).filter((k) => !ctx[k]);
  if (missingDeps.length) {
    console.log(`  ⤳ SKIP (missing: ${missingDeps.join(", ")})`);
    return { ...base, status: "SKIP", httpStatus: null, openapiValid: null, sdkOk: false, note: `missing dependency: ${missingDeps.join(", ")}` };
  }

  const trackReady = await waitIfNeeded(step, ctx, baseUrl, username, password);
  if (!trackReady) {
    return { ...base, status: "SKIP", httpStatus: null, openapiValid: null, sdkOk: false, note: "track not ready after timeout" };
  }

  const request = step.request(ctx);
  const sdkRes = await retryLoop(step, request, baseUrl, username, password);

  if (!sdkRes.ok) {
    const msg = `${sdkRes.error?.name ?? "Error"}: ${sdkRes.error?.message ?? "SDK call failed"}`;
    console.log(`  ❌ FAIL — ${msg.split("\n")[0].slice(0, 120)}`);
    writeArtifacts(step.operationId, sdkRes.error?.bodyJson ?? null, sdkRes.error ?? null);
    return { ...base, status: "FAIL", httpStatus: sdkRes.error?.statusCode ?? null, openapiValid: null, sdkOk: false, sdkError: msg };
  }

  if (step.capture) step.capture(sdkRes.value, ctx);
  await waitForMediaIfNeeded(step, ctx, baseUrl, username, password);

  const capturedId = getCapturedId(step.operationId, ctx);
  const { openapiValid, openapiErrors } = validateStepResponse(spec, ep, sdkRes);
  const { missingInSDK, missingInAPI } = computeStepDiff(sdkRes);

  writeArtifacts(step.operationId, sdkRes.rawBody, sdkRes.value);

  const status: StepResult["status"] =
    openapiValid !== false && missingInSDK.length === 0 && missingInAPI.length === 0 ? "PASS" : "FAIL";

  const idSuffix = capturedId ? ` id=${capturedId}` : "";
  const statusLabel = status === "PASS" ? "✅ PASS" : "❌ FAIL";
  console.log(`  ${statusLabel} (HTTP ${sdkRes.statusCode ?? "?"})${idSuffix}`);

  return { ...base, status, httpStatus: sdkRes.statusCode, openapiValid, openapiErrors, sdkOk: true, missingInSDK, missingInAPI, capturedId };
}

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
    console.log(`[${i + 1}/${STEPS.length}] (${step.phase}) ${step.operationId}`);
    const result = await runStep(step, ep, ctx, { spec, baseUrl, username, password });
    if (result.capturedId) {
      // capturedId surfaced for logging; ctx already updated by capture()
    }
    results.push(result);
  }

  writeReport(results, ctx);

  const pass = results.filter((r) => r.status === "PASS").length;
  const fail = results.filter((r) => r.status === "FAIL").length;
  const skip = results.filter((r) => r.status === "SKIP").length;
  console.log(`Summary: total=${results.length} pass=${pass} fail=${fail} skip=${skip}`);
}

await main();