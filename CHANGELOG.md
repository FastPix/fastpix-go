# Changelog

---

Regenerated against the updated OpenAPI spec (`fastpix-openai.yaml`).
`old.yaml` is the vendored spec describing what this SDK actually implements:
`fastpix-openai.yaml` minus the `/ai/` endpoints the SDK does not expose, plus
`VideoTrackForGetAll.frameRate` (see *Kept against the spec* below). The
endpoint validators resolve `old.yaml` by default.

### Kept against the spec

- **`VideoTrackForGetAll.frameRate` retained.** The updated spec drops this
  property, but the live API still returns it on `list-media` video tracks
  (`"frameRate": "60.000"`), and the spec's own `example` block for that schema
  still shows `frameRate: 30/1`. Removing it would have reverted the 1.1.3 fix
  *"Added missing `frameRate` field to video tracks in the list media response"*.
  The upstream spec needs correcting.

- **`maxDuration` keeps `minimum: 0`** on `getCreateLiveStreamResponseDTO`,
  `CreateLiveStreamResponseDTO`, and `patchResponseData`. The updated spec
  raises the minimum to 60 on these three *response* schemas, but the live API
  returns `maxDuration: 0` for streams with no enforced limit — the meaning the
  old spec spelled out (*"`0` means no enforced maximum (unbounded)"*). With
  `minimum: 60`, every live-stream create/update/get response fails schema
  validation despite the SDK parsing it correctly. There is no request-side
  `maxDuration`, so this only ever constrained server output. The upstream spec
  needs correcting.

### Breaking

- **`mp4Support` is now a list of renditions instead of a single string.**
  On `GetAllMediaResponse`, `GetMediaDetailResponse`, `Live-Media-Clips`,
  `Media`, `sourceAccessMedia`, and `Update-Media`, the field changed from
  `*XMp4Support` (enum: `none`, `capped_4k`, `audioOnly`,
  `audioOnly,capped_4k`) to
  `optionalnullable.OptionalNullable[[]XMp4Support]`, where each entry is:

  | Field    | Type                | Values                            |
  |----------|---------------------|-----------------------------------|
  | `Type`   | `*XMp4SupportType`   | `capped_4k`, `audioOnly`          |
  | `Status` | `*XMp4SupportStatus` | `preparing`, `ready`, `failed`    |
  | `Height` | `*int64` (nullable)  | omitted for `audioOnly`           |
  | `Width`  | `*int64` (nullable)  | omitted for `audioOnly`           |
  | `Ext`    | `*XMp4SupportExt`    | `mp4`, `m4a`                      |

  The **request** side is unchanged — `CreateMediaRequestMp4Support`,
  `UpdatedMp4SupportMp4Support`, and `DirectUploadVideoMediaMp4Support` are
  still string enums.

- **`components.GetMediaResponse` renamed to `components.GetMediaDetailResponse`**
  (and its nested types: `...MediaQuality`, `...MaxResolution`,
  `...SourceResolution`, `...Status`, `...Mp4Support`, `...Track`).
  `operations.GetMediaResponse` is derived from the operation ID and is
  **not** renamed.

- **`UpdateTrackRequest.URL` removed.** A track's file can no longer be
  swapped; only `LanguageCode`, `LanguageName`, and the new `Title` may be
  updated. `url` is also no longer in the schema's `required` list.

- **`UpdateMediaMaxResolution` no longer accepts `360p`**, matching the other
  `maxResolution` enums.

### Added

- **Optional `Title` on audio/subtitle tracks** — added to `AddTrackRequest`,
  `AddTrackResponse`, `AudioTrack`, `GenerateTrackResponse`, `SubtitleTrack`,
  `TrackSubtitlesGenerateRequest`, `UpdateTrackRequest`,
  `UpdateTrackResponse`, `VideoTrack`, and `VideoTrackForGetAll`.

- **`OptimizeAudio`** (nullable `bool`, "whether the audio track has been
  volume-normalized") on the six media schemas listed above.

- **`sourceResolution` gained `360p` / `360`** on `GetAllMediaResponse`,
  `GetMediaDetailResponse`, `Live-Media-Clips`, `Media`, and
  `sourceAccessMedia`. `Update-Media` already carried both.

### Changed

- Documentation links updated to the current `fastpix.com/docs` structure:
  `vod-events/*` and `ai-events/*` → `webhooks/*`, `video-intelligence/*` →
  `in-video-ai/*`, `manage-audio-and-subtitle-tracks/*`,
  `playback-and-delivery/*`, `edit-and-transform-videos/*` →
  `video-on-demand/*`, `manage-live-streams/*`, `broadcast/*`,
  `edit-and-transform-live-stream/*` → `live-streaming/*`,
  `working-with-video-data/*` and `concepts/*` → `video-data/*`.

- `docs/` regenerated for all of the above.

### Compatibility

- Callers reading `Mp4Support` must switch from a single enum value to
  iterating the rendition list.
- Callers naming `components.GetMediaResponse` must use
  `components.GetMediaDetailResponse`.
- Callers setting `UpdateTrackRequest.URL` must drop it.

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
