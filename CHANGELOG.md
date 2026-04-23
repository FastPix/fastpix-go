# Changelog

All notable changes to this project will be documented in this file.
---
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
