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
Last generated: 2026-06-04T12:43:20.820Z

- **Total GET endpoints**: 30
- **PASS**: 26
- **FAIL**: 4
- **SKIP**: 0

| Endpoint | OperationId | OpenAPI valid | SDK parse | Missing in SDK (present in API) | Missing in API (present in SDK) | Empty arrays omitted by SDK | Status |
|---|---|---:|---:|---|---|---|---|
| `/on-demand` | `list-media` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{livestreamId}/live-clips` | `list-live-clips` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{mediaId}` | `get-media` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{mediaId}/summary` | `get-media-summary` | ✅ | ❌ | None | None | None | ❌ FAIL |
| `/on-demand/{mediaId}/input-info` | `retrieveMediaInputInfo` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{mediaId}/playback-ids` | `list-playback-ids` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/uploads` | `list-uploads` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{mediaId}/media-clips` | `get-media-clips` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/playlists` | `get-all-playlists` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/playlists/{playlistId}` | `get-playlist-by-id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/{mediaId}/playback-ids/{playbackId}` | `get-playback-id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/on-demand/drm-configurations` | `getDrmConfiguration` | ✅ | ❌ | None | None | None | ❌ FAIL |
| `/on-demand/drm-configurations/{drmConfigurationId}` | `getDrmConfigurationById` | ✅ | ❌ | None | None | None | ❌ FAIL |
| `/live/streams` | `get-all-streams` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/live/streams/{streamId}/viewer-count` | `get-live-stream-viewer-count-by-id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/live/streams/{streamId}` | `get-live-stream-by-id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/live/streams/{streamId}/playback-ids/{playbackId}` | `get-live-stream-playback-id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/live/streams/{streamId}/simulcast/{simulcastId}` | `get-specific-simulcast-of-stream` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/iam/signing-keys` | `list_signing_keys` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/iam/signing-keys/{signingKeyId}` | `get-signing_key_by_id` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/viewlist` | `list_video_views` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/viewlist/{viewId}` | `get_video_view_details` | ✅ | ❌ | None | None | None | ❌ FAIL |
| `/data/viewlist/top-content` | `list_by_top_content` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/dimensions` | `list_dimensions` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/dimensions/{dimensionsId}` | `list_filter_values_for_dimension` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/metrics/{metricId}/breakdown` | `list_breakdown_values` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/metrics/{metricId}/overall` | `list_overall_values` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/metrics/{metricId}/timeseries` | `get_timeseries_data` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/metrics/comparison` | `list_comparison_values` | ✅ | ✅ | None | None | None | ✅ PASS |
| `/data/errors` | `list_errors` | ✅ | ✅ | None | None | None | ✅ PASS |

#### Missing fields (full lists)

- **list-media** (`/on-demand`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list-live-clips** (`/on-demand/{livestreamId}/live-clips`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-media** (`/on-demand/{mediaId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-media-summary** (`/on-demand/{mediaId}/summary`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **retrieveMediaInputInfo** (`/on-demand/{mediaId}/input-info`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list-playback-ids** (`/on-demand/{mediaId}/playback-ids`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list-uploads** (`/on-demand/uploads`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-media-clips** (`/on-demand/{mediaId}/media-clips`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-all-playlists** (`/on-demand/playlists`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-playlist-by-id** (`/on-demand/playlists/{playlistId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-playback-id** (`/on-demand/{mediaId}/playback-ids/{playbackId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **getDrmConfiguration** (`/on-demand/drm-configurations`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **getDrmConfigurationById** (`/on-demand/drm-configurations/{drmConfigurationId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-all-streams** (`/live/streams`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-live-stream-viewer-count-by-id** (`/live/streams/{streamId}/viewer-count`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-live-stream-by-id** (`/live/streams/{streamId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-live-stream-playback-id** (`/live/streams/{streamId}/playback-ids/{playbackId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-specific-simulcast-of-stream** (`/live/streams/{streamId}/simulcast/{simulcastId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_signing_keys** (`/iam/signing-keys`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get-signing_key_by_id** (`/iam/signing-keys/{signingKeyId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_video_views** (`/data/viewlist`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get_video_view_details** (`/data/viewlist/{viewId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_by_top_content** (`/data/viewlist/top-content`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_dimensions** (`/data/dimensions`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_filter_values_for_dimension** (`/data/dimensions/{dimensionsId}`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_breakdown_values** (`/data/metrics/{metricId}/breakdown`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_overall_values** (`/data/metrics/{metricId}/overall`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **get_timeseries_data** (`/data/metrics/{metricId}/timeseries`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_comparison_values** (`/data/metrics/comparison`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None
- **list_errors** (`/data/errors`)
  - **Missing in SDK (present in API)**: None
  - **Missing in API (present in SDK)**: None
  - **Empty arrays omitted by SDK**: None
  - **Empty arrays omitted by API**: None

Full details: `tests/GET_ENDPOINTS_OPENAPI_RESPONSE_VALIDATION_REPORT.md`
<!-- END GET_ENDPOINTS_CONSOLIDATED -->
