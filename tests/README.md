# FastPix Go SDK Test Suite

This directory contains comprehensive test cases for the FastPix Go SDK. The tests cover all major functionality including media management, live streaming, analytics, and more.

## Prerequisites

Before running the tests, you need to set up the following environment variables:

### Required Environment Variables

```bash
export FASTPIX_USERNAME="your-username"
export FASTPIX_PASSWORD="your-password"
```

### Optional Environment Variables

```bash
export FASTPIX_API_KEY="your-api-key"  # For API key authentication tests
export FASTPIX_BASE_URL="https://api.fastpix.io/v1"  # Default base URL
```

## Test Structure

The test suite is organized into the following files:

### Core Tests
- **`config.go`** - Test configuration and setup utilities
- **`auth_test.go`** - Authentication and security tests

### Feature Tests
- **`media_management_test.go`** - Media upload, management, and operations
- **`live_streaming_test.go`** - Live stream creation, management, and control
- **`playback_playlist_test.go`** - Playback IDs and playlist management
- **`analytics_metrics_test.go`** - Analytics, metrics, and reporting
- **`signing_keys_test.go`** - Cryptographic signing key management
- **`ai_features_test.go`** - AI-powered features (summaries, chapters, etc.)
- **`simulcast_test.go`** - Simulcast streaming to multiple platforms
- **`drm_configurations_test.go`** - DRM configuration management

## Running Tests

### Run All Tests
```bash
go test ./tests/... -v
```

### Run Specific Test Categories
```bash
# Media management tests
go test ./tests/media_management_test.go -v

# Live streaming tests
go test ./tests/live_streaming_test.go -v

# Analytics tests
go test ./tests/analytics_metrics_test.go -v
```

### Run Tests with Coverage
```bash
go test ./tests/... -v -cover
```

### Run Tests in Parallel
```bash
go test ./tests/... -v -parallel 4
```

## Test Features

### Comprehensive Coverage
- **64 API endpoints** covered across all test files
- **Authentication methods** (Basic Auth, API Key)
- **Error handling** and edge cases
- **Pagination** and filtering
- **CRUD operations** for all resources

### Smart Test Design
- **Conditional execution** - Tests skip when required data is not available
- **Resource cleanup** - Tests clean up created resources when possible
- **Realistic scenarios** - Tests use real API calls with proper error handling
- **Modular structure** - Each test file focuses on specific functionality

### Test Data Management
- Tests automatically find existing resources when possible
- Create test resources when needed
- Skip tests gracefully when prerequisites aren't met
- Log important information for debugging

## Important Test Cases

### 1. Authentication Tests
- Basic authentication with username/password
- API key authentication
- Invalid credentials handling
- Security source configuration

### 2. Media Management Tests
- List all media with pagination
- Create media from URLs
- Get specific media by ID
- Update media metadata
- List uploads and media clips

### 3. Live Streaming Tests
- Create new live streams
- List and manage streams
- Enable/disable streams
- Get viewer counts
- Complete streams

### 4. Playback & Playlist Tests
- Create and manage playback IDs
- Create and manage playlists
- Add media to playlists
- Update playlist metadata

### 5. Analytics Tests
- Video view analytics
- Metrics and breakdowns
- Timeseries data
- Top content analysis
- Error reporting

### 6. AI Features Tests
- Media summarization
- Chapter generation
- Named entity extraction
- Content moderation

## Environment Setup

### Development Environment
```bash
# Set your credentials
export FASTPIX_USERNAME="your-username"
export FASTPIX_PASSWORD="your-password"

# Run tests
go test ./tests/... -v
```

### CI/CD Environment
```bash
# Set credentials in your CI environment
export FASTPIX_USERNAME="$FASTPIX_USERNAME"
export FASTPIX_PASSWORD="$FASTPIX_PASSWORD"

# Run tests with timeout
go test ./tests/... -v -timeout 10m
```

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Verify your username and password are correct
   - Check that your account has proper permissions
   - Ensure the base URL is correct

2. **Test Failures**
   - Some tests may fail if no data exists (this is expected)
   - Tests are designed to skip gracefully when prerequisites aren't met
   - Check the test logs for specific error messages

3. **Timeout Issues**
   - Increase the timeout for slow operations
   - Check your network connection
   - Verify the API is accessible

### Debug Mode
```bash
# Run tests with verbose output
go test ./tests/... -v -args -test.v

# Run specific test with debug info
go test ./tests/media_management_test.go -v -run TestCreateMedia
```

## Contributing

When adding new tests:
1. Follow the existing test structure
2. Use descriptive test names
3. Include proper error handling
4. Add appropriate skip conditions
5. Clean up test resources when possible

## License

These tests are part of the FastPix Go SDK and follow the same license terms.
