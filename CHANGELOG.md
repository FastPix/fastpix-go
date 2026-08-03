# Changelog

---
## [1.1.5]

Regenerated against the updated OpenAPI spec. Validators resolve `old.yaml` —
that spec minus the `/ai/` endpoints, plus the mismatches below.

### Breaking

- `mp4Support` is now `OptionalNullable[[]XMp4Support]` instead of a string
  enum, on the six media schemas. Each entry has `Type`, `Status`, `Height`,
  `Width`, `Ext`. Request-side enums unchanged.
- `components.GetMediaResponse` → `GetMediaDetailResponse`, with its nested
  types. `operations.GetMediaResponse` is not renamed.
- `UpdateTrackRequest.URL` removed — a track's file can no longer be swapped.
- `UpdateMediaMaxResolution` drops `360p`.

### Added

- `Title` on the ten track schemas.
- `OptimizeAudio` (nullable `bool`) on the six media schemas.
- `sourceResolution` gains `360p`/`360`.

### Changed

- SDK version `1.1.4` → `1.1.5`.
- Doc links rewritten to current `fastpix.com/docs` paths; `docs/` regenerated;
  hardcoded IDs in examples replaced with placeholders; `UpdateTrack` example
  arg order corrected to `(ctx, trackID, mediaID)`.

### Spec mismatches

Not followed, because they contradict the live API:

- `VideoTrackForGetAll.frameRate` kept — still returned by `list-media`.
- `maxDuration` keeps `minimum: 0` — streams with no limit return `0`.

---
## [1.1.4]

### Changed

- **SDK version bump: `1.1.3` → `1.1.4`.**
  A maintenance release that updates the SDK's internal version identifiers.
  It contains no functional, API, or behavioral changes and is fully
  backward compatible with `1.1.3`.

  Updated identifiers:
  - `SDKVersion` constant — now reports `1.1.4`.
  - `User-Agent` header — outbound requests now identify as
    `fastpix-sdk/go 1.1.4`.

### Compatibility

- No changes to public types, method signatures, request/response models,
  default server URLs, hooks, or retry logic.
- No action required for existing integrations — update the dependency and
  rebuild.

---
## [1.1.3]

### Changed — FastPix `.io` → `.com` migration

FastPix is moving its hosts and documentation to the `.com` TLD. This release updates every reference in the SDK and its docs:

| Old (`.io`) | New (`.com`) |
|---|---|
| `api.fastpix.io` | `api.fastpix.com` |
| `stream.fastpix.io` | `stream.fastpix.com` |
| `images.fastpix.io` | `images.fastpix.com` |
| `static.fastpix.io` | `static.fastpix.com` |
| `docs.fastpix.io/...` | `fastpix.com/docs/...` |

- The default server URL is now `https://api.fastpix.com/v1/` (was `https://api.fastpix.io/v1/`). If you rely on SDK defaults, no code change is required — upgrading and rebuilding is enough.
- If you pass an explicit override (e.g. `fastpixgo.New(fastpixgo.WithServerURL("https://api.fastpix.io/v1/"))`), change it to `https://api.fastpix.com/v1/`.
- All README and per-service documentation links now point at `https://fastpix.com/docs/...`.

The `.io` hosts still serve traffic during the transition but are slated for deprecation — update any hard-coded `.io` references (server URL, playback/asset URLs) in your application.

### Fixed
- Added missing `frameRate` field to video tracks in the list media response (`VideoTrackForGetAll`).
- Added missing `data.message` field to the delete signing key response (`DeleteSigningKeyResponse`).

## [1.1.2]

### Fixed
- Fixed data event field remapping in hooks.

## [1.1.1]

- Replaced documentation placeholders from `your-video-id` to `your-media-id`.

## [1.1.0]

### Fixed
- Fixed missing parameters in multiple API methods.

### Improved
- Improved overall developer experience through more accurate typings.

## [1.0.0]

### Added
- Complete API coverage for Media, Live Streaming, Video Data, and Signing Keys
- Go 1.22+ support with comprehensive type safety
- Media upload, management, and processing capabilities
- Live streaming with simulcasting support
- Video analytics and performance tracking
- Cryptographic signing keys for secure authentication
- In-video AI processing features
- DRM configuration and management
- Playlist creation and management
- Comprehensive error handling with specific error types
- Built-in retry mechanisms and timeout handling
- Comprehensive test suite with 14 test files
- Full API specification compliance

### Changed
- Reorganized package structure for better maintainability
- Updated dependencies to modern Go packages
- Improved API design with better error handling
- Enhanced documentation and examples
- Updated SDK version to 1.0.0
- Updated minimum Go version requirement to 1.22+ for better compatibility and performance

### Fixed
- Direct upload metadata handling
- Response object access patterns
- Type mismatches in method parameters
- Error handling for validation responses
- Test data structure alignment with API expectations
- Improved error handling with specific error types
- Fixed type annotation issues for better IDE support
- Ensured consistent API patterns across modules

---

## [0.0.1]

### Added
- Initial release of FastPix Go SDK
- HTTP client support with Go standard library
- Media API integration with upload, management, and processing
- Playback ID management for media files
- Media operations (list, get, update, delete)
- Direct upload support for video files
- Live stream API integration
- Live stream management (create, update, delete)
- Playback ID management for live streams
- Simulcast configuration for live streams
- HTTP Basic authentication support
- Server URL override support
- Comprehensive error handling and custom error classes
- Example usage and quick start documentation
- Go Modules integration
- Type checking integration
- Test framework integration
