# FastPix Go SDK Examples

This directory contains comprehensive, production-ready examples demonstrating how to use the FastPix Go SDK for various video and live streaming operations.

## Quick Start

### Prerequisites
1. **Go 1.22 or later** installed
2. **FastPix API credentials** (access token and secret key)
3. **Environment variables** set up:
   ```bash
   export FASTPIX_USERNAME="your-access-token"
   export FASTPIX_PASSWORD="your-secret-key"
   ```

### Running Examples

#### Individual Examples
```bash
# Basic usage and connectivity
go run ./basic-usage

# Media management operations
go run ./media-management

# Live streaming operations
go run ./live-streaming

# Playback and playlist management
go run ./playback-playlist

# AI-powered video features
go run ./ai-features

# Analytics and metrics
go run ./analytics-metrics

# Security and signing keys
go run ./security-signing-keys

# Error handling patterns
go run ./error-handling

# Create a signed upload URL (then PUT your file to it)
go run ./create-upload

# Verify a FastPix webhook signature (offline; no credentials needed)
go run ./verify-webhook
```

> Each example is its own package under `examples/`, so you run it by folder
> path — e.g. `go run ./basic-usage`. `verify-webhook` runs offline and needs no
> credentials.

#### Using Makefile
```bash
# Show all available targets
make help

# Run specific example
make basic
make media
make live
make playback
make ai
make analytics
make security
make errors
make create-upload
make verify-webhook

# Run all examples
make all

# Clean up generated files
make clean
```

## Example Overview

| Example | Description | Key Features |
|--------|-------------|--------------|
| `basic-usage/` | SDK initialization and basic connectivity | Authentication, basic API calls |
| `media-management/` | Complete media lifecycle management | Upload, list, update, delete, tracks |
| `live-streaming/` | Live streaming operations | Create streams, simulcast, playback |
| `playback-playlist/` | Playback and playlist management | Playlists, DRM, playback IDs |
| `ai-features/` | AI-powered video features | Summaries, chapters, moderation |
| `analytics-metrics/` | Analytics and performance metrics | Views, metrics, dimensions |
| `security-signing-keys/` | Security and cryptographic keys | Key management, rotation |
| `error-handling/` | Comprehensive error handling | Retry, fallback, graceful degradation |
| `create-upload/` | Create a signed direct-upload URL | Direct upload, upload lifecycle |
| `verify-webhook/` | Verify a FastPix webhook signature | HMAC-SHA256, constant-time compare |

## Key Features Demonstrated

### 🔐 Authentication & Security
- Basic authentication with username/password
- Environment variable configuration
- Cryptographic signing key management
- Key rotation strategies

### 📹 Media Operations
- Upload from URLs and direct file upload
- Comprehensive metadata management
- Audio and subtitle track management
- Source access control and MP4 support

### 📡 Live Streaming
- Complete stream lifecycle management
- Real-time stream control (enable/disable)
- Playback ID management
- Simulcasting to external platforms

### 🤖 AI Features
- Automatic video summarization
- Chapter generation
- Named entity extraction
- Content moderation

### 📊 Analytics & Monitoring
- Comprehensive view tracking
- Performance metrics and breakdowns
- Error monitoring and analytics
- Custom dimensions and filtering

### 🛡️ Error Handling & Resilience
- Specific error type handling
- Retry mechanisms with exponential backoff
- Fallback strategies
- Graceful degradation

## Production-Ready Features

All examples include:
- ✅ **Robust Error Handling**: Comprehensive error checking and logging
- ✅ **Nil Pointer Safety**: All pointer accesses are protected
- ✅ **Response Validation**: Proper validation of API responses
- ✅ **Consistent Patterns**: Uniform approach across all examples
- ✅ **Clear Documentation**: Well-commented and documented code
- ✅ **Real-world Scenarios**: Practical examples developers can use

## Best Practices

1. **Always handle errors** - Check for errors after every API call
2. **Use context** - Pass context for cancellation and timeouts
3. **Implement retries** - Use retry configuration for resilience
4. **Secure credentials** - Never hardcode credentials in code
5. **Monitor usage** - Track API usage and implement rate limiting
6. **Clean up resources** - Delete test data after operations
7. **Use pagination** - Handle large datasets with pagination
8. **Log appropriately** - Implement proper logging for debugging

## Common Patterns

### SDK Initialization
```go
client := fastpixgo.New(
    fastpixgo.WithSecurity(components.Security{
        Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")), // your-access-token
        Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")), // your-secret-key
    }),
    fastpixgo.WithTimeout(30*time.Second),
)
```

### Error Handling
```go
response, err := client.ManageVideos.GetMedia(ctx, mediaID)
if err != nil {
    var notFoundErr *apierrors.NotFoundError
    if errors.As(err, &notFoundErr) {
        // Handle not found error
    } else {
        // Handle other errors
    }
}
```

### Retry Configuration
```go
client := fastpixgo.New(
    fastpixgo.WithRetryConfig(fastpixgo.RetryConfig{
        Strategy: "backoff",
        Backoff: &fastpixgo.BackoffStrategy{
            InitialInterval: 1,
            MaxInterval:     50,
            Exponent:        2,
            MaxElapsedTime:  100,
        },
        RetryConnectionErrors: true,
    }),
)
```

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   - Verify credentials are correct
   - Check environment variables are set
   - Ensure account has proper permissions

2. **Timeout Errors**
   - Increase timeout duration
   - Check network connectivity
   - Implement retry logic

3. **Validation Errors**
   - Check request parameters
   - Verify required fields are provided
   - Review API documentation

4. **Rate Limiting**
   - Implement exponential backoff
   - Monitor API usage
   - Consider request batching

### Getting Help

- Check the [main README](../README.md) for general SDK information
- Review the [API documentation](https://fastpix.com/docs) for endpoint details
- Contact FastPix support for account-specific issues

## License

These examples are provided under the same license as the FastPix Go SDK.