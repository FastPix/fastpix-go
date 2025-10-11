# FastPix Go SDK Samples

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

### Running Samples

#### Individual Samples
```bash
# Basic usage and connectivity
go run basic_usage.go

# Media management operations
go run media_management.go

# Live streaming operations
go run live_streaming.go

# Playback and playlist management
go run playback_playlist.go

# AI-powered video features
go run ai_features.go

# Analytics and metrics
go run analytics_metrics.go

# Security and signing keys
go run security_signing_keys.go

# Error handling patterns
go run error_handling.go
```

#### Using Makefile
```bash
# Show all available targets
make help

# Run specific sample
make basic
make media
make live
make playback
make ai
make analytics
make security
make errors

# Run all samples
make all

# Clean up generated files
make clean
```

## Sample Overview

| Sample | Description | Key Features |
|--------|-------------|--------------|
| `basic_usage.go` | SDK initialization and basic connectivity | Authentication, basic API calls |
| `media_management.go` | Complete media lifecycle management | Upload, list, update, delete, tracks |
| `live_streaming.go` | Live streaming operations | Create streams, simulcast, playback |
| `playback_playlist.go` | Playback and playlist management | Playlists, DRM, playback IDs |
| `ai_features.go` | AI-powered video features | Summaries, chapters, moderation |
| `analytics_metrics.go` | Analytics and performance metrics | Views, metrics, dimensions |
| `security_signing_keys.go` | Security and cryptographic keys | Key management, rotation |
| `error_handling.go` | Comprehensive error handling | Retry, fallback, graceful degradation |

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

All samples include:
- ✅ **Robust Error Handling**: Comprehensive error checking and logging
- ✅ **Nil Pointer Safety**: All pointer accesses are protected
- ✅ **Response Validation**: Proper validation of API responses
- ✅ **Consistent Patterns**: Uniform approach across all samples
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
- Review the [API documentation](https://docs.fastpix.io) for endpoint details
- Contact FastPix support for account-specific issues

## License

These samples are provided under the same license as the FastPix Go SDK.