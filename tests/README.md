# Endpoint Validation (Hybrid: OpenAPI via Node, SDK via Go)

These dev tools validate the generated **Go** SDK against the live FastPix API and
the OpenAPI spec. The orchestration runs in TypeScript (for the OpenAPI
response-schema validation and JSON-path diffing); the SDK itself is exercised
through a small Go harness (`tests/sdkharness`) invoked as a subprocess.

## Quick Start

1. Install Node deps (from the SDK repo root):

```bash
npm install --prefix tests
```

2. Ensure the Go toolchain is available and the module builds:

```bash
# From the SDK repo root
go build ./tests/sdkharness
```

3. Set env vars:

```bash
export FASTPIX_USERNAME="your-username"
export FASTPIX_PASSWORD="your-password"
# optional:
# export FASTPIX_BASE_URL="https://api.fastpix.com/v1/"
# export FASTPIX_OPENAPI_SPEC="/abs/path/to/spec.yaml"   # overrides spec auto-discovery
#   Auto-discovery: ../openapi.yaml (untracked snapshot of the upstream spec at the repo root).
```

4. Run:

```bash
cd tests
npm run validate:get-endpoints       # GET endpoints
npm run validate:non-get-endpoints   # POST/PUT/PATCH/DELETE lifecycle
```

Artifacts and reports are written into `tests/`.

## Overview

The GET validator implements a **hybrid testing approach**:

1. **Calls the API directly** via HTTP to get raw JSON responses
2. **Validates API responses** against the OpenAPI schema using `openapi-response-validator`
3. **Calls the Go SDK** for the same operation (via the `tests/sdkharness` subprocess)
4. **Compares API vs SDK responses** to identify:
   - Fields missing in SDK (present in API but dropped by SDK parsing)
   - Fields missing in API (present in SDK but not in API response)
   - Empty arrays omitted in SDK vs API
5. **Generates artifacts** (API vs SDK JSON files) and validation reports

The non-GET validator mutates live data, so it cannot hit the API and the SDK
separately. It runs a **create → update → delete** lifecycle through the SDK,
captures each created resource id for downstream steps, and validates the SDK's
raw wire response (captured by a body-tee'ing HTTP client in the harness)
against the OpenAPI schema.

## How It Works

### TypeScript drivers

- `validate-get-endpoints.ts` — extracts all GET endpoints from the spec, calls
  each via fetch + the SDK, diffs the two, and writes:
  - `GET_ENDPOINTS_OPENAPI_RESPONSE_VALIDATION_REPORT.md`
  - `GET_ENDPOINTS_OPENAPI_RESPONSE_FIX_SUGGESTIONS.md`
  - the consolidated summary block in this `README.md`
- `validate-non-get-endpoints.ts` — runs the mutating lifecycle and writes
  `NON_GET_ENDPOINTS_VALIDATION_REPORT.md`.

### Go SDK harness (`tests/sdkharness/main.go`)

The drivers invoke `go run ./tests/sdkharness` (cwd = repo root so the local
module resolves), passing a JSON payload on stdin:

```json
{ "operationId": "...", "request": { ... }, "baseUrl": "...", "username": "...", "password": "..." }
```

The harness dispatches to the matching SDK method, then prints a JSON result:

- success: `{ "ok": true, "value": <body>, "statusCode": <int|null>, "rawBody": <json|string|null> }`
- failure: `{ "ok": false, "error": { "name", "message", "statusCode?", "bodyJson?" } }`

`go run` reuses the Go build cache, so only the first call pays the compile cost.

## Fixtures

`get-endpoints-fixtures.json` contains real IDs for GET endpoints that require
path parameters. Update it with working IDs from your FastPix account for
accurate testing. If a fixture is missing, the GET driver falls back to a
placeholder UUID, which typically yields a 404.

<!-- BEGIN GET_ENDPOINTS_CONSOLIDATED -->
_The GET validator rewrites this section with its latest results on every run. Snapshots are not committed._
<!-- END GET_ENDPOINTS_CONSOLIDATED -->
