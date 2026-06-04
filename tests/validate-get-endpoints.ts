#!/usr/bin/env tsx
/*
 * GET endpoints validator using `openapi-response-validator`
 *
 * Per GET endpoint in OpenAPI spec:
 * - Calls the API to get the raw JSON response
 * - Validates the raw response against the OpenAPI response schema using `openapi-response-validator`
 * - Calls the Go SDK method for the same operationId (via the tests/sdkharness
 *   subprocess), capturing either the success object OR the returned error (normalized)
 * - Compares JSON paths between raw API JSON and SDK-parsed JSON (including the same normalization rules)
 * - Persists per-endpoint artifacts to disk (API and SDK)
 * - Generates two markdown reports in `tests/`:
 *   - `GET_ENDPOINTS_OPENAPI_RESPONSE_VALIDATION_REPORT.md`
 *   - `GET_ENDPOINTS_OPENAPI_RESPONSE_FIX_SUGGESTIONS.md`
 * - Updates `tests/README.md` by replacing the block between markers:
 *   - `<!-- BEGIN GET_ENDPOINTS_CONSOLIDATED -->`
 *   - `<!-- END GET_ENDPOINTS_CONSOLIDATED -->`
 *
 * Requirements:
 * - FASTPIX_USERNAME / FASTPIX_PASSWORD env vars (Basic Auth)
 * - `tests/get-endpoints-fixtures.json` for endpoints with required path params (optional but recommended)
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
// eslint-disable-next-line no-console

type Fixture = {
  operations: Record<
    string,
    {
      pathParams?: Record<string, string>;
      query?: Record<string, string | number | boolean | Array<string | number | boolean>>;
    }
  >;
};

type EndpointInfo = {
  path: string;
  method: "GET";
  operationId: string;
  responses: any;
  parameters: Array<any>;
};

type FixSuggestion = {
  title: string;
  why: string;
  where?: string;
  pasteYaml?: string;
};

type EndpointResult = {
  endpoint: string;
  operationId: string;
  method: "GET";
  openapiValid: boolean;
  openapiErrors: Array<{ path?: string; message?: string; errorCode?: string }>;
  sdkParseOk: boolean;
  sdkParseError?: string;
  missingInSDK: string[];
  missingInAPI: string[];
  emptyArraysOmittedInSDK: string[];
  emptyArraysOmittedInAPI: string[];
  apiResponseFile?: string;
  sdkResponseFile?: string;
  apiResponsePreview?: string;
  sdkResponsePreview?: string;
  status: "PASS" | "FAIL";
  note?: string;
  fixSuggestions?: FixSuggestion[];
};

const ARTIFACTS_DIRNAME = "artifacts";
const MAX_PREVIEW_CHARS = 4000;
const PLACEHOLDER_UUID = "00000000-0000-0000-0000-000000000000";
const FIX_SUGGESTIONS_MD = "GET_ENDPOINTS_OPENAPI_RESPONSE_FIX_SUGGESTIONS.md";

function safeFileSlug(input: string): string {
  return input.replace(/[^a-zA-Z0-9._-]+/g, "_");
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
  if (text.length <= MAX_PREVIEW_CHARS) return text;
  return text.slice(0, MAX_PREVIEW_CHARS) + "\n... (truncated)";
}

function writeArtifactFiles(
  operationId: string,
  rawBody: unknown,
  sdkBody: unknown,
): {
  apiPath: string;
  sdkPath: string;
  apiPreview: string;
  sdkPreview: string;
} {
  const artifactsDir = join(__dirname, ARTIFACTS_DIRNAME);
  mkdirSync(artifactsDir, { recursive: true });
  // eslint-disable-next-line no-console

  const slug = safeFileSlug(operationId);
  const apiFilename = `${slug}.api.json`;
  const sdkFilename = `${slug}.sdk.json`;

  const apiText = toPrettyJson(rawBody);
  const sdkText = toPrettyJson(sdkBody);

  const apiPath = join(artifactsDir, apiFilename);
  const sdkPath = join(artifactsDir, sdkFilename);

  writeFileSync(apiPath, apiText);
  writeFileSync(sdkPath, sdkText);

  return {
    apiPath: `tests/${ARTIFACTS_DIRNAME}/${apiFilename}`,
    sdkPath: `tests/${ARTIFACTS_DIRNAME}/${sdkFilename}`,
    apiPreview: preview(apiText),
    sdkPreview: preview(sdkText),
  };
}

function defaultSDKRequest(operationId: string): any {
  // Ensure SDK input validation passes so we reach the HTTP call and get server errors on failures.
  switch (operationId) {
    case "get-media":
    case "get-media-summary":
    case "retrieveMediaInputInfo":
    case "list-playback-ids":
    case "get-media-clips":
      return { mediaId: PLACEHOLDER_UUID };
    case "get-playback-id":
      return { mediaId: PLACEHOLDER_UUID, playbackId: PLACEHOLDER_UUID };
    case "list-live-clips":
      return { livestreamId: PLACEHOLDER_UUID };
    case "get-playlist-by-id":
      return { playlistId: PLACEHOLDER_UUID };
    case "getDrmConfigurationById":
      return { drmConfigurationId: PLACEHOLDER_UUID };
    case "get-live-stream-by-id":
    case "get-live-stream-viewer-count-by-id":
      return { streamId: PLACEHOLDER_UUID };
    case "get-live-stream-playback-id":
      return { streamId: PLACEHOLDER_UUID, playbackId: PLACEHOLDER_UUID };
    case "get-specific-simulcast-of-stream":
      return { streamId: PLACEHOLDER_UUID, simulcastId: PLACEHOLDER_UUID };
    case "get-signing_key_by_id":
      return { signingKeyId: PLACEHOLDER_UUID };
    case "get_video_view_details":
      return { viewId: PLACEHOLDER_UUID };
    case "list_filter_values_for_dimension":
      return { dimensionsId: "browser_name" };
    case "list_breakdown_values":
      return {
        metricId: "quality_of_experience_score",
        timespan: "24:hours",
        groupBy: "browser_name",
      };
    case "list_overall_values":
      return { metricId: "quality_of_experience_score", timespan: "24:hours" };
    case "get_timeseries_data":
      return {
        metricId: "quality_of_experience_score",
        timespan: "24:hours",
        groupBy: "hour",
      };
    case "list_comparison_values":
      return { timespan: "24:hours", dimension: "browser_name", value: "Chrome" };
    case "list_errors":
      return { timespan: "24:hours", limit: 5 };
    case "list_video_views":
      return { timespan: "24:hours", limit: 5, offset: 1 };
    case "list_by_top_content":
      return { timespan: "24:hours", limit: 5 };
    case "list-media":
      return { limit: 5, offset: 1, orderBy: "desc" };
    case "list-uploads":
      return { limit: 5, offset: 1, orderBy: "desc" };
    case "get-all-streams":
      return { limit: 5, offset: 1, orderBy: "desc" };
    case "getDrmConfiguration":
      return { limit: 10, offset: 1 };
    case "get-all-playlists":
      return { limit: 5, offset: 1 };
    case "list_signing_keys":
      return { limit: 5, offset: 1 };
    case "list_dimensions":
      return undefined;
    default:
      return undefined;
  }
}

function buildSDKRequest(endpoint: EndpointInfo, fixtures: Fixture | null): any {
  const opFixture = fixtures?.operations?.[endpoint.operationId];
  const fromFixture = opFixture
    ? { ...opFixture.pathParams, ...opFixture.query }
    : undefined;

  // If fixtures exist, use them as-is (they match SDK request shapes).
  if (fromFixture) return fromFixture;

  // Prefer operation-specific defaults (handles required query params too).
  const def = defaultSDKRequest(endpoint.operationId);
  if (def !== undefined) return def;

  // Otherwise: auto-generate a placeholder request object for required path params.
  const requiredPathParams = endpoint.parameters
    .filter((p) => p?.in === "path" && p?.required)
    .map((p) => p.name);

  if (requiredPathParams.length === 0) return undefined;

  const req: Record<string, string> = {};
  for (const name of requiredPathParams) req[name] = PLACEHOLDER_UUID;
  return req;
}

type GoSDKResult =
  | { ok: true; value: any }
  | { ok: false; error: any };


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
        "/opt/homebrew/bin/go",        // Homebrew on Apple Silicon / Intel Mac
        "/usr/local/go/bin/go",        // Standard Go install on Linux/macOS
        "/usr/local/bin/go",
        "/usr/bin/go",
        "/home/linuxbrew/.linuxbrew/bin/go", // Homebrew on Linux
      ];
  for (const p of candidates) {
    if (existsSync(p)) return p;
  }
  return "/opt/homebrew/bin/go";
}

const GO_BINARY = resolveAbsoluteGoBinary();

// Path to the Go harness package, relative to the SDK repo root (the parent of
// tests/). `go run` is executed with cwd = repo root so the harness resolves
// the local module's packages.
const GO_HARNESS_PKG = "./tests/sdkharness";

// Invoke the Go SDK by running the compiled harness (tests/sdkharness) as a
// subprocess and passing the operation + request as JSON on stdin. The harness
// dispatches to the matching SDK method and prints back the deserialized
// response body (or a normalized error). `go run` caches builds, so only the
// first call pays the compile cost.
function invokeGoSDK(
  operationId: string,
  request: any,
  baseUrl: string,
  username: string,
  password: string,
): GoSDKResult {
  const repoRoot = join(__dirname, "..");
  mkdirSync(join(repoRoot, ".gotmp"), { recursive: true });

  const child = spawnSync(GO_BINARY, ["run", GO_HARNESS_PKG], {
    input: JSON.stringify({ operationId, request: request ?? {}, baseUrl, username, password }),
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

  // Build/compile failures (and panics) surface on stderr with no JSON on stdout.
  if (!stdout.startsWith("{") && !stdout.startsWith("[")) {
    const msg = (stderr || stdout).slice(0, 1000) || "Go harness produced no output";
    return { ok: false, error: { name: "GoRuntimeError", message: msg } };
  }
  if (stderr) {
    // Non-fatal go toolchain chatter can still accompany valid output; log it.
    console.error(`go stderr: ${stderr.split("\n").slice(0, 3).join(" ")}`);
  }

  try {
    const parsed = JSON.parse(stdout);
    if (parsed?.ok) return { ok: true, value: parsed.value };
    return { ok: false, error: parsed?.error ?? { name: "GoSDKError", message: stdout.slice(0, 500) } };
  } catch (e: any) {
    return { ok: false, error: { name: "GoOutputParseError", message: `Failed to parse JSON: ${e.message}. Output: ${stdout.slice(0, 500)}` } };
  }
}

function readFixtures(): Fixture | null {
  const p = join(__dirname, "get-endpoints-fixtures.json");
  if (!existsSync(p)) return null;
  return JSON.parse(readFileSync(p, "utf-8")) as Fixture;
}

function resolveSpecPath(): string {
  // Deterministic search order. An explicit FASTPIX_OPENAPI_SPEC env var wins,
  // then the known spec filenames at the repo root (parent of tests/).
  const candidates = [
    process.env.FASTPIX_OPENAPI_SPEC,
    join(__dirname, "../fixed 7.yaml"),
    join(__dirname, "../fastpix.yaml"),
    join(__dirname, "../../fastpix.yaml"),
    join(__dirname, "../fixed.yaml"),
    join(__dirname, "../../fixed.yaml"),
    join(__dirname, "../fastpix-openapi.yaml"),
    join(__dirname, "../../fastpix-openapi.yaml"),
  ].filter((p): p is string => Boolean(p));
  for (const p of candidates) {
    if (existsSync(p)) return p;
  }
  throw new Error(
    `OpenAPI spec not found. Tried: ${candidates.map((c) => JSON.stringify(c)).join(", ")}`,
  );
}

function loadOpenAPISpec(): any {
  const specPath = resolveSpecPath();
  return yaml.load(readFileSync(specPath, "utf-8"));
}

function extractGetEndpoints(spec: any): EndpointInfo[] {
  const out: EndpointInfo[] = [];
  for (const [path, methods] of Object.entries(spec.paths || {})) {
    const m = methods as any;
    if (!m.get) continue;
    out.push({
      path,
      method: "GET",
      operationId: m.get.operationId,
      responses: m.get.responses || {},
      parameters: [...(m.get.parameters || []), ...(m.parameters || [])],
    });
  }
  return out;
}

// Convert OpenAPI 3 schema refs (#/components/schemas/X) to the format used by openapi-response-validator (#/definitions/X)
function convertRefsToDefinitions(node: any): any {
  if (node == null || typeof node !== "object") return node;
  if (Array.isArray(node)) return node.map(convertRefsToDefinitions);
  const out: any = {};
  for (const [k, v] of Object.entries(node)) {
    if (k === "$ref" && typeof v === "string") {
      out[k] = v.replace("#/components/schemas/", "#/definitions/");
    } else {
      out[k] = convertRefsToDefinitions(v);
    }
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
    responses[status] = {
      description: d.description || "",
      schema: convertRefsToDefinitions(schema),
    };
  }

  if (Object.keys(responses).length === 0) return null;

  return new OpenAPIResponseValidator({
    responses,
    definitions,
  });
}

function hasOpenapiError(r: EndpointResult, includes: string): boolean {
  return (r.openapiErrors || []).some((e) => (e?.message ?? "").includes(includes));
}

function openapiErrorPaths(r: EndpointResult): string[] {
  return (r.openapiErrors || [])
    .map((e) => e?.path)
    .filter((p): p is string => typeof p === "string" && p.length > 0);
}

function generateFixSuggestions(r: EndpointResult): FixSuggestion[] {
  const out: FixSuggestion[] = [];
  const paths = openapiErrorPaths(r);

  // 1) Generic: oneOf overlap on tracks
  const hasTracksOneOf =
    hasOpenapiError(r, "must match exactly one schema in oneOf") &&
    paths.some((p) => p.includes("tracks"));
  if (hasTracksOneOf) {
    out.push({
      title: "Fix `tracks[].oneOf` overlap by constraining `type` per track schema",
      why:
        "The current track schemas overlap (e.g. `type` is a free string and distinguishing fields are not required), so a single track object can match multiple branches. `oneOf` requires exactly one match.",
      where:
        "In OpenAPI spec: `components/schemas/{VideoTrack,VideoTrackForGetAll,AudioTrack,SubtitleTrack}.properties.type`",
      pasteYaml: [
        "# Apply these changes inside each schema's `properties:` block:",
        "",
        "# VideoTrack (and VideoTrackForGetAll)",
        "type:",
        "  type: string",
        "  enum: [video]",
        "  example: video",
        "",
        "# AudioTrack",
        "type:",
        "  type: string",
        "  enum: [audio]",
        "  example: audio",
        "",
        "# SubtitleTrack",
        "type:",
        "  type: string",
        "  enum: [subtitle]",
        "  example: subtitle",
      ].join("\n"),
    });
  }

  // 2) Enum mismatch: sourceResolution
  const hasSourceResolutionEnum =
    hasOpenapiError(r, "must be equal to one of the allowed values") &&
    paths.some((p) => p.includes("sourceResolution"));
  if (hasSourceResolutionEnum) {
    out.push({
      title: "Fix `sourceResolution` enum mismatch (API may return values without `p`)",
      why:
        "The API can return values like `\"1080\"` but the spec constrains the enum to `\"1080p\"`-style values.",
      where:
        "In OpenAPI spec: under the relevant media response schema(s) `sourceResolution:` field definition",
    });
  }

  // 3) Redundant oneOf for /data/dimensions
  const hasDimensionsOneOf =
    hasOpenapiError(r, "must match exactly one schema in oneOf") &&
    (r.endpoint === "/data/dimensions" || paths.some((p) => p.includes("dimensions")));
  if (hasDimensionsOneOf) {
    out.push({
      title: "Remove redundant `oneOf` on `/data/dimensions` response schema",
      why:
        "`data` is defined as `oneOf: [array<string>, $ref: Dimensions]` and `Dimensions` itself is also `array<string>`, so valid responses can match multiple branches.",
      where:
        "In OpenAPI spec: `paths./data/dimensions.get.responses.200.content.application/json.schema.properties.data.oneOf`",
    });
  }

  // 4) Overlapping numeric oneOf: integer vs number
  const hasIntegerVsNumber =
    hasOpenapiError(r, "must match exactly one schema in oneOf") &&
    paths.some((p) => p.includes("value"));
  if (hasIntegerVsNumber) {
    out.push({
      title: "Avoid `oneOf: [integer, number]` overlaps (integers are also numbers)",
      why:
        "In JSON Schema, `integer` is a subset of `number`. A value like `0` matches both, causing oneOf validation errors.",
      where:
        "In OpenAPI spec: metrics schemas that use `oneOf: [integer, number]`",
    });
  }

  // 5) Nullable mismatch: fpApiVersion
  const hasFpApiVersionNull =
    hasOpenapiError(r, "must be string") &&
    paths.some((p) => p.includes("fpApiVersion"));
  if (hasFpApiVersionNull) {
    out.push({
      title: "Make `fpApiVersion` nullable in the spec",
      why: "The API can return `null` for fpApiVersion but the schema declares `string` only.",
      where: "In OpenAPI spec: `components/schemas/Views.properties.fpApiVersion`",
    });
  }

  // 6) Placeholder fixture guidance (common 404)
  const placeholderUsed = (r.note || "").includes("Placeholder used");
  const likely404 =
    r.sdkParseOk === false &&
    /404|not found/i.test(r.sdkParseError || "") &&
    placeholderUsed;
  if (likely404) {
    out.push({
      title: "Provide real fixture IDs for this operationId",
      why:
        "A placeholder UUID was used for required path params; the API likely returned 404 because the resource doesn't exist. Add a real ID under `tests/get-endpoints-fixtures.json` for this operationId.",
    });
  }

  // 7) Playlist playOrder default / missing
  const playOrderMissing = r.missingInAPI.some((p) => p.includes("playOrder")) ||
    r.missingInSDK.some((p) => p.includes("playOrder"));
  if (playOrderMissing) {
    out.push({
      title: "Ensure `playOrder` is correctly modeled for smart playlists only",
      why:
        "If `playOrder` is present/required only for `type: smart`, the response schemas should reflect that (e.g. discriminator split).",
      where:
        "In OpenAPI spec: playlist response schemas for create/update/get-by-id",
    });
  }

  // 8) simulcastResponses missing
  const hasSimulcastResponses = r.missingInSDK.some((p) => p.includes("simulcastResponses"));
  if (hasSimulcastResponses) {
    out.push({
      title: "Add `simulcastResponses` to the live stream response schema",
      why:
        "The API response includes simulcastResponses but the OpenAPI schema (and generated SDK inbound schema) does not, causing the SDK to drop the field.",
      where:
        "In OpenAPI spec: live stream response schema(s) for get/list streams",
    });
  }

  return out;
}


function formatPaths(paths: string[]): string {
  return paths.length ? paths.map((p) => `\`${p}\``).join(", ") : "None";
}

function pushPathList(lines: string[], paths: string[]): void {
  if (paths.length === 0) {
    lines.push("- None");
  } else {
    for (const p of paths) lines.push(`- \`${p}\``);
  }
}

function pushOpenAPIErrors(lines: string[], errors: Array<{ path?: string; message?: string; errorCode?: string }>): void {
  lines.push("### Observed OpenAPI errors", "");
  for (const e of errors) {
    const loc = e.path ? `\`${e.path}\`` : "";
    const msg = e.message ?? "";
    lines.push(`- ${loc} ${msg}`.trim());
  }
  lines.push("");
}

function pushSuggestion(lines: string[], s: FixSuggestion): void {
  lines.push(`- **${s.title}**`, `  - **why**: ${s.why}`);
  if (s.where) lines.push(`  - **where**: ${s.where}`);
  if (s.pasteYaml) {
    lines.push("  - **paste**:", "", "```yaml", s.pasteYaml, "```");
  }
  lines.push("");
}

function pushFailingEndpointSection(lines: string[], r: EndpointResult): void {
  const suggestions = r.fixSuggestions ?? [];
  lines.push(
    `## ${r.operationId} (\`${r.endpoint}\`)`,
    "",
    `- **Status**: ${r.status}`,
    `- **OpenAPI valid**: ${r.openapiValid ? "yes" : "no"}`,
    `- **SDK parse**: ${r.sdkParseOk ? "ok" : "failed"}`,
  );
  if (r.apiResponseFile) lines.push(`- **API artifact**: \`${r.apiResponseFile}\``);
  if (r.sdkResponseFile) lines.push(`- **SDK artifact**: \`${r.sdkResponseFile}\``);
  lines.push("");
  if (!r.openapiValid && (r.openapiErrors?.length ?? 0) > 0) {
    pushOpenAPIErrors(lines, r.openapiErrors);
  }
  if (suggestions.length === 0) {
    lines.push("### Suggested fixes", "", "- No heuristic suggestions available for this failure yet.", "");
    return;
  }
  lines.push("### Suggested fixes", "");
  for (const s of suggestions) {
    pushSuggestion(lines, s);
  }
}

function writeFixSuggestions(results: EndpointResult[]) {
  const failing = results.filter((r) => r.status === "FAIL");
  const outPath = join(__dirname, FIX_SUGGESTIONS_MD);
  const lines: string[] = [];

  lines.push(
    "# GET Endpoints — OpenAPI Response Fix Suggestions",
    "",
    `Generated: ${new Date().toISOString()}`,
    "",
    `Total failing endpoints: ${failing.length}`,
    "",
  );

  for (const r of failing) {
    pushFailingEndpointSection(lines, r);
  }

  writeFileSync(outPath, lines.join("\n"));
}

function collectEmptyArraysFromArray(arr: any[], prefix: string, out: Set<string>): void {
  const arrayPrefix = prefix ? `${prefix}[]` : "[]";
  for (const item of arr) {
    for (const p of collectEmptyArrayFieldPaths(item, arrayPrefix)) out.add(p);
  }
}

function collectEmptyArraysFromObject(obj: Record<string, any>, prefix: string, out: Set<string>): void {
  for (const [k, v] of Object.entries(obj)) {
    const p = prefix ? `${prefix}.${k}` : k;
    if (Array.isArray(v) && v.length === 0) out.add(p);
    for (const child of collectEmptyArrayFieldPaths(v, p)) out.add(child);
  }
}

function collectEmptyArrayFieldPaths(value: any, prefix = ""): Set<string> {
  const out = new Set<string>();
  if (value === null || value === undefined || typeof value !== "object") return out;
  if (Array.isArray(value)) {
    collectEmptyArraysFromArray(value, prefix, out);
  } else {
    collectEmptyArraysFromObject(value as Record<string, any>, prefix, out);
  }
  return out;
}

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
  for (const item of arr) {
    for (const p of collectJsonPaths(item, arrayPrefix, opts)) out.add(p);
  }
}

function collectJsonPaths(
  value: any,
  prefix = "",
  opts: { includeEmptyArrays?: boolean } = {},
): Set<string> {
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

function sortUnique(arr: string[]) {
  return Array.from(new Set(arr)).sort((a, b) => a.localeCompare(b));
}

function canonicalizeKey(key: string): string {
  // 1) snake_case -> camelCase
  const camel = key.includes("_")
    ? key
        .toLowerCase()
        .replace(/_([a-z0-9])/g, (_, c) => String(c).toUpperCase())
    : key;

  // 2) normalize acronyms casing
  return camel.replaceAll("SDK", "Sdk").replaceAll("API", "Api");
}

// Mirrors internal/hooks/events_field_remap_hook.go: the get_video_view_details API
// returns abbreviated wire keys on data.events[] that the SDK intentionally
// remaps to the spec-shaped long names. Apply the same remap to the raw API
// body so the comparison reflects what the SDK is contracted to emit rather
// than flagging the deliberate rename as a discrepancy.
const EVENT_OUTER_REMAP: Record<string, string> = {
  pt: "player_playhead_time",
  e: "event_name",
  d: "event_details",
  vt: "viewer_time",
  et: "event_time",
};
const EVENT_INNER_REMAP: Record<string, string> = {
  br: "bitrate",
  h: "height",
  w: "width",
  cd: "codec",
  host: "hostName",
  txt: "text",
  c: "code",
  err: "error",
  t: "type",
  u: "url",
};

function remapApiForComparison(operationId: string, body: any): any {
  if (operationId !== "get_video_view_details") return body;
  const events = body?.data?.events;
  if (!Array.isArray(events)) return body;

  const rebuiltEvents = events.map((event) => {
    if (event === null || typeof event !== "object" || Array.isArray(event)) {
      return event;
    }
    const rebuilt: Record<string, any> = {};
    for (const [key, value] of Object.entries(event)) {
      const newKey = EVENT_OUTER_REMAP[key] ?? key;
      if (
        newKey === "event_details"
        && value !== null
        && typeof value === "object"
        && !Array.isArray(value)
      ) {
        const inner: Record<string, any> = {};
        for (const [ik, iv] of Object.entries(value)) {
          inner[EVENT_INNER_REMAP[ik] ?? ik] = iv;
        }
        rebuilt[newKey] = inner;
      } else {
        rebuilt[newKey] = value;
      }
    }
    return rebuilt;
  });

  return { ...body, data: { ...body.data, events: rebuiltEvents } };
}

function normalizeJsonForComparison(value: any): any {
  if (value === null || value === undefined) return value;
  if (Array.isArray(value)) return value.map(normalizeJsonForComparison);
  if (typeof value !== "object") return value;
  const out: any = {};
  for (const [k, v] of Object.entries(value)) {
    out[canonicalizeKey(k)] = normalizeJsonForComparison(v);
  }
  return out;
}

function jsonRoundTrip(value: any): any {
  return structuredClone(value);
}

function substitutePathParams(
  path: string,
  requiredPathParams: string[],
  effectiveReq: Record<string, any>,
): { path: string; note: string | undefined } {
  let note: string | undefined;
  for (const name of requiredPathParams) {
    const val = effectiveReq[name] ?? PLACEHOLDER_UUID;
    if (effectiveReq[name] == null) {
      note = note ? `${note}; placeholder used for ${name}` : `Placeholder used for ${name}`;
    }
    path = path.replaceAll(`{${name}}`, encodeURIComponent(val));
  }
  return { path, note };
}

function appendQueryParams(url: URL, queryParams: any[], effectiveReq: Record<string, any>): void {
  for (const p of queryParams) {
    const name: string = p.name;
    const baseName = name.endsWith("[]") ? name.slice(0, -2) : name;
    const val = effectiveReq[name] ?? effectiveReq[baseName];
    if (val == null) continue;
    if (Array.isArray(val)) {
      for (const item of val) url.searchParams.append(name, String(item));
    } else if (name.endsWith("[]")) {
      url.searchParams.append(name, String(val));
    } else {
      url.searchParams.set(name, String(val));
    }
  }
}

function buildUrl(
  baseUrl: string,
  endpoint: EndpointInfo,
  fixture: Fixture | null,
): { url: string; note?: string } {
  const opFixture = fixture?.operations?.[endpoint.operationId];
  const defaults = defaultSDKRequest(endpoint.operationId) ?? {};
  const fromFixture = opFixture ? { ...opFixture.pathParams, ...opFixture.query } : {};
  const effectiveReq: Record<string, any> = { ...defaults, ...fromFixture };

  const requiredPathParams = endpoint.parameters
    .filter((p) => p?.in === "path" && p?.required)
    .map((p) => p.name);

  const { path, note } = substitutePathParams(endpoint.path, requiredPathParams, effectiveReq);

  const base = baseUrl.endsWith("/") ? baseUrl : baseUrl + "/";
  const url = new URL(path.replace(/^\//, ""), base);

  appendQueryParams(url, endpoint.parameters.filter((p) => p?.in === "query"), effectiveReq);

  return { url: url.toString(), note };
}

function basicAuthHeader(username: string, password: string): string {
  const token = Buffer.from(`${username}:${password}`).toString("base64");
  return `Basic ${token}`;
}


function buildTableRow(r: EndpointResult): string {
  const openapiCol = r.openapiValid ? "✅" : "❌";
  const sdkCol = r.sdkParseOk ? "✅" : "❌";
  const missSdk = formatPaths(r.missingInSDK);
  const missApi = formatPaths(r.missingInAPI);
  const emptyOmitted = formatPaths(r.emptyArraysOmittedInSDK);
  const status = r.status === "PASS" ? "✅ PASS" : "❌ FAIL";
  return `| \`${r.endpoint}\` | \`${r.operationId}\` | ${openapiCol} | ${sdkCol} | ${missSdk} | ${missApi} | ${emptyOmitted} | ${status} |`;
}

function pushEndpointMeta(lines: string[], r: EndpointResult): void {
  lines.push(
    `### ${r.operationId} (\`${r.endpoint}\`)`,
    "",
    `- **Status**: ${r.status}`,
  );
  if (r.note) lines.push(`- **Note**: ${r.note}`);
  lines.push(`- **OpenAPI valid**: ${r.openapiValid ? "yes" : "no"}`);
  if (!r.openapiValid && r.openapiErrors.length) {
    lines.push("- **OpenAPI errors**:");
    for (const e of r.openapiErrors) {
      const loc = e.path ? `\`${e.path}\`` : "";
      const msg = e.message ?? "";
      lines.push(`  - ${loc} ${msg}`.trim());
    }
  }
  lines.push(`- **SDK parse**: ${r.sdkParseOk ? "ok" : "failed"}`);
  if (!r.sdkParseOk && r.sdkParseError) lines.push(`- **SDK parse error**: ${r.sdkParseError}`);
  if (r.apiResponseFile) lines.push(`- **API response file**: \`${r.apiResponseFile}\``);
  if (r.sdkResponseFile) lines.push(`- **SDK response file**: \`${r.sdkResponseFile}\``);
  lines.push("");
}

function pushResponsePreview(lines: string[], label: string, preview: string): void {
  lines.push(`**${label}**`, "", "```json", preview, "```", "");
}

function pushPathSection(lines: string[], label: string, paths: string[]): void {
  lines.push(`**${label}**`, "");
  pushPathList(lines, paths);
}

function pushEndpointDetail(lines: string[], r: EndpointResult): void {
  pushEndpointMeta(lines, r);
  if (r.apiResponsePreview) pushResponsePreview(lines, "API response (preview)", r.apiResponsePreview);
  if (r.sdkResponsePreview) pushResponsePreview(lines, "SDK response (preview)", r.sdkResponsePreview);
  pushPathSection(lines, `Missing in SDK (present in API) — ${r.missingInSDK.length}`, r.missingInSDK);
  lines.push("");
  pushPathSection(lines, `Missing in API (present in SDK) — ${r.missingInAPI.length}`, r.missingInAPI);
  lines.push("");
  pushPathSection(lines, `Empty arrays omitted by SDK — ${r.emptyArraysOmittedInSDK.length}`, r.emptyArraysOmittedInSDK);
  lines.push("");
  pushPathSection(lines, `Empty arrays omitted by API — ${r.emptyArraysOmittedInAPI.length}`, r.emptyArraysOmittedInAPI);
  lines.push("");
}

function writeReport(results: EndpointResult[]) {
  const total = results.length;
  const passed = results.filter((r) => r.status === "PASS").length;
  const failed = results.filter((r) => r.status === "FAIL").length;
  const skipped = 0;

  const reportPath = join(__dirname, "GET_ENDPOINTS_OPENAPI_RESPONSE_VALIDATION_REPORT.md");
  const readmePath = join(__dirname, "README.md");
  const generatedAt = new Date().toISOString();

  const lines: string[] = [];
  lines.push(
    "# GET Endpoints — OpenAPI Response Validation Report",
    "",
    `Generated: ${generatedAt}`,
    "",
    "## Summary",
    "",
    `- **Total GET endpoints**: ${total}`,
    `- **PASS**: ${passed}`,
    `- **FAIL**: ${failed}`,
    `- **SKIP**: ${skipped}`,
    "",
    "## Consolidated report",
    "",
    "| Endpoint | OperationId | OpenAPI valid | SDK parse | Missing in SDK (present in API) | Missing in API (present in SDK) | Empty arrays omitted by SDK | Status |",
    "|---|---|---:|---:|---|---|---|---|",
  );

  for (const r of results) {
    lines.push(buildTableRow(r));
  }

  lines.push("", "## Per-endpoint details (full missing parameter lists)", "");

  for (const r of results) {
    pushEndpointDetail(lines, r);
  }

  writeFileSync(reportPath, lines.join("\n"));
  writeFixSuggestions(results);

  // Also update tests/README.md with the consolidated report section so it always stays in sync.
  try {
    if (existsSync(readmePath)) {
      const begin = "<!-- BEGIN GET_ENDPOINTS_CONSOLIDATED -->";
      const end = "<!-- END GET_ENDPOINTS_CONSOLIDATED -->";

      const consolidated: string[] = [];
      consolidated.push(
        `Last generated: ${generatedAt}`,
        "",
        `- **Total GET endpoints**: ${total}`,
        `- **PASS**: ${passed}`,
        `- **FAIL**: ${failed}`,
        `- **SKIP**: ${skipped}`,
        "",
        "| Endpoint | OperationId | OpenAPI valid | SDK parse | Missing in SDK (present in API) | Missing in API (present in SDK) | Empty arrays omitted by SDK | Status |",
        "|---|---|---:|---:|---|---|---|---|",
      );
      for (const r of results) {
        consolidated.push(buildTableRow(r));
      }
      consolidated.push("", "#### Missing fields (full lists)", "");
      for (const r of results) {
        consolidated.push(
          `- **${r.operationId}** (\`${r.endpoint}\`)`,
          `  - **Missing in SDK (present in API)**: ${formatPaths(r.missingInSDK)}`,
          `  - **Missing in API (present in SDK)**: ${formatPaths(r.missingInAPI)}`,
          `  - **Empty arrays omitted by SDK**: ${formatPaths(r.emptyArraysOmittedInSDK)}`,
          `  - **Empty arrays omitted by API**: ${formatPaths(r.emptyArraysOmittedInAPI)}`,
        );
      }
      consolidated.push("", `Full details: \`tests/GET_ENDPOINTS_OPENAPI_RESPONSE_VALIDATION_REPORT.md\``);

      const readme = readFileSync(readmePath, "utf-8");
      if (readme.includes(begin) && readme.includes(end)) {
        const block = `${begin}\n${consolidated.join("\n")}\n${end}`;
        const updated = readme.replace(new RegExp(String.raw`${begin}[\s\S]*?${end}`), block);
        writeFileSync(readmePath, updated);
      }
    }
  } catch {
    // ignore README update failures
  }

  // eslint-disable-next-line no-console
  console.log(`Report generated: ${reportPath}`);
  // eslint-disable-next-line no-console
  console.log(`Fix suggestions generated: ${join(__dirname, FIX_SUGGESTIONS_MD)}`);
  // eslint-disable-next-line no-console
  console.log(`Summary: total=${total} pass=${passed} fail=${failed} skip=${skipped}`);
}

type FetchResult = { httpStatus: number; rawBody: any; requestError: string | undefined };

async function fetchApiResponse(url: string, username: string, password: string): Promise<FetchResult> {
  let httpStatus = 0;
  let rawBody: any = null;
  let requestError: string | undefined;
  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 30000);
    const res = await fetch(url, {
      method: "GET",
      headers: { Accept: "application/json", Authorization: basicAuthHeader(username, password) },
      signal: controller.signal,
    });
    clearTimeout(timeoutId);
    httpStatus = res.status;
    const bodyText = await res.text();
    try { rawBody = bodyText ? JSON.parse(bodyText) : null; } catch { rawBody = bodyText; }
  } catch (e: any) {
    requestError = e.name === "AbortError" ? "Request timeout (30s)" : (e?.message ?? String(e));
    // eslint-disable-next-line no-console
    console.error(`  ⚠️  API request failed: ${requestError}`);
  }
  return { httpStatus, rawBody, requestError };
}

function validateApiResponse(
  spec: any, ep: EndpointInfo, httpStatus: number, rawBody: any, requestError: string | undefined,
): { openapiValid: boolean; openapiErrors: any[] } {
  if (requestError) return { openapiValid: false, openapiErrors: [{ message: `Request failed: ${requestError}` }] };
  const validator = makeOpenAPIResponseValidator(spec, ep);
  if (!validator) return { openapiValid: true, openapiErrors: [] };
  const err = validator.validateResponse(String(httpStatus), rawBody);
  return err ? { openapiValid: false, openapiErrors: err.errors || [] } : { openapiValid: true, openapiErrors: [] };
}

type SDKResult = { sdkParseOk: boolean; sdkParseError: string | undefined; sdkPrinted: any; sdkValueForDiff: any };

function invokeSDK(ep: EndpointInfo, fixtures: Fixture | null, baseUrl: string, username: string, password: string): SDKResult {
  const sdkReq = buildSDKRequest(ep, fixtures);
  const sdk = invokeGoSDK(ep.operationId, sdkReq, baseUrl, username, password);
  if (sdk.ok) return { sdkParseOk: true, sdkParseError: undefined, sdkPrinted: sdk.value, sdkValueForDiff: sdk.value };
  // eslint-disable-next-line no-console
  console.error(`  ⚠️  Go SDK call failed: ${sdk.error?.message ?? "Go SDK call failed"}`);
  // Persist just the raw API error body (when present) as the SDK artifact so it
  // matches the .api.json byte-for-byte, instead of the harness error envelope
  // ({name, message, statusCode, bodyJson}).
  return { sdkParseOk: false, sdkParseError: sdk.error?.message ?? "Go SDK call failed", sdkPrinted: sdk.error?.bodyJson ?? sdk.error, sdkValueForDiff: null };
}

type DiffResult = { missingInSDK: string[]; missingInAPI: string[]; emptyArraysOmittedInSDK: string[]; emptyArraysOmittedInAPI: string[] };

function computeDiff(operationId: string, rawBody: any, sdkValueForDiff: any): DiffResult {
  const apiNormalized = normalizeJsonForComparison(remapApiForComparison(operationId, rawBody));
  const sdkJsonLike = (sdkValueForDiff && typeof sdkValueForDiff === "object") ? jsonRoundTrip(sdkValueForDiff) : null;
  const sdkNormalized = sdkJsonLike ? normalizeJsonForComparison(sdkJsonLike) : null;
  const apiPaths = collectJsonPaths(apiNormalized, "", { includeEmptyArrays: false });
  const sdkPaths = sdkNormalized ? collectJsonPaths(sdkNormalized, "", { includeEmptyArrays: false }) : new Set<string>();
  const missingInSDK = sdkPaths.size ? sortUnique([...apiPaths].filter((p) => !sdkPaths.has(p))) : [];
  const missingInAPI = sdkPaths.size ? sortUnique([...sdkPaths].filter((p) => !apiPaths.has(p))) : [];
  const apiStrictPaths = collectJsonPaths(apiNormalized, "", { includeEmptyArrays: true });
  const sdkStrictPaths = sdkNormalized ? collectJsonPaths(sdkNormalized, "", { includeEmptyArrays: true }) : new Set<string>();
  const apiEmptyArrayFields = collectEmptyArrayFieldPaths(apiNormalized);
  const sdkEmptyArrayFields = sdkNormalized ? collectEmptyArrayFieldPaths(sdkNormalized) : new Set<string>();
  const emptyArraysOmittedInSDK = sortUnique([...apiEmptyArrayFields].filter((p) => !sdkStrictPaths.has(p)));
  const emptyArraysOmittedInAPI = sortUnique([...sdkEmptyArrayFields].filter((p) => !apiStrictPaths.has(p)));
  return { missingInSDK, missingInAPI, emptyArraysOmittedInSDK, emptyArraysOmittedInAPI };
}

async function processEndpoint(
  ep: EndpointInfo,
  spec: any,
  fixtures: Fixture | null,
  baseUrl: string,
  username: string,
  password: string,
): Promise<EndpointResult> {
  let rawBody: any = null;
  let sdkPrinted: any = null;
  let result: EndpointResult;

  try {
    const { url, note } = buildUrl(baseUrl, ep, fixtures);
    const fetchRes = await fetchApiResponse(url, username, password);
    rawBody = fetchRes.rawBody;
    const { httpStatus, requestError } = fetchRes;
    const { openapiValid, openapiErrors } = validateApiResponse(spec, ep, httpStatus, rawBody, requestError);
    const sdkRes = invokeSDK(ep, fixtures, baseUrl, username, password);
    sdkPrinted = sdkRes.sdkPrinted;
    const { sdkParseOk, sdkParseError, sdkValueForDiff } = sdkRes;
    const { missingInSDK, missingInAPI, emptyArraysOmittedInSDK, emptyArraysOmittedInAPI } = computeDiff(ep.operationId, rawBody, sdkValueForDiff);
    const pass = openapiValid && sdkParseOk && missingInSDK.length === 0 && missingInAPI.length === 0;
    result = {
      endpoint: ep.path, operationId: ep.operationId, method: "GET",
      openapiValid, openapiErrors, sdkParseOk, sdkParseError,
      missingInSDK, missingInAPI, emptyArraysOmittedInSDK, emptyArraysOmittedInAPI,
      apiResponseFile: undefined, sdkResponseFile: undefined,
      apiResponsePreview: undefined, sdkResponsePreview: undefined,
      status: pass ? "PASS" : "FAIL", note, fixSuggestions: undefined,
    };
  } catch (error: any) {
    // eslint-disable-next-line no-console
    console.error(`  ✗ Unexpected error processing ${ep.operationId}:`, error?.message ?? String(error));
    result = {
      endpoint: ep.path, operationId: ep.operationId, method: "GET",
      openapiValid: false, openapiErrors: [{ message: `Unexpected error: ${error?.message ?? String(error)}` }],
      sdkParseOk: false, sdkParseError: error?.message ?? String(error),
      missingInSDK: [], missingInAPI: [], emptyArraysOmittedInSDK: [], emptyArraysOmittedInAPI: [],
      status: "FAIL", note: "Unexpected error during processing", fixSuggestions: undefined,
    };
  }

  // Write artifacts OUTSIDE try/catch so they always get written
  const artifacts = writeArtifactFiles(ep.operationId, rawBody, sdkPrinted);
  result.apiResponseFile = artifacts.apiPath;
  result.sdkResponseFile = artifacts.sdkPath;
  result.apiResponsePreview = artifacts.apiPreview;
  result.sdkResponsePreview = artifacts.sdkPreview;

  return result;
}
const ENV_FALLBACK_USER = "your-access-token";
const ENV_FALLBACK_PASS = "your-secret-key";

async function main(): Promise<void> {
  const spec = loadOpenAPISpec();
  const endpoints = extractGetEndpoints(spec);
  const fixtures = readFixtures();

  const baseUrl: string =
    process.env.FASTPIX_BASE_URL
    ?? ((spec.servers?.[0]?.url as string | undefined) ?? "https://api.fastpix.com/v1/");

  const username = process.env.FASTPIX_USERNAME ?? ENV_FALLBACK_USER;
  const password = process.env.FASTPIX_PASSWORD ?? ENV_FALLBACK_PASS;

  if (!username || !password || username === ENV_FALLBACK_USER || password === ENV_FALLBACK_PASS) {
    throw new Error("Set FASTPIX_USERNAME and FASTPIX_PASSWORD env vars for BasicAuth (use real credentials for live API validation)");
  }

  const results: EndpointResult[] = [];
  const totalEndpoints = endpoints.length;

  for (let i = 0; i < endpoints.length; i++) {
    const ep = endpoints[i];
    // eslint-disable-next-line no-console
    console.log(`[${i + 1}/${totalEndpoints}] Processing: ${ep.operationId} (${ep.path})`);
    const result = await processEndpoint(ep, spec, fixtures, baseUrl, username, password);
    results.push(result);
    // eslint-disable-next-line no-console
    console.log(`  ✓ Completed: ${ep.operationId} - ${result.status}`);
  }

  for (const r of results) {
    if (r.status !== "FAIL") continue;
    r.fixSuggestions = generateFixSuggestions(r);
  }

  writeReport(results);
}

await main();