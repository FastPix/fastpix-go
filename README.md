# FastPix Go SDK

A robust, type-safe Go SDK designed for seamless integration with the FastPix API platform.

## Introduction

The FastPix Go SDK simplifies integration with the FastPix platform. It provides a clean, strongly-typed interface for secure and efficient communication with the FastPix API, enabling easy management of media uploads, live streaming, on‑demand content, playlists, video analytics, and signing keys for secure access and token management. It is intended for use with Go 1.21 and above.

## Prerequisites

### Environment and Version Support

| Requirement | Version | Description |
|---|---:|---|
| Go | `1.21+` | Core runtime environment |
| Go Modules | `Enabled` | Dependency management |
| Internet | `Required` | API communication and authentication |

> Pro Tip: We recommend using Go 1.21+ for optimal performance and the latest language features.

### Getting Started with FastPix

To get started with the FastPix Go SDK, ensure you have the following:

- The FastPix APIs are authenticated using a **Username** and a **Password**. You must generate these credentials to use the SDK.
- Follow the steps in the [Authentication with Basic Auth](https://docs.fastpix.io/docs/basic-authentication) guide to obtain your credentials.

### Environment Variables (Optional)

Configure your FastPix credentials using environment variables for enhanced security and convenience:

```bash
# Set your FastPix credentials
export FASTPIX_USERNAME="your-access-token"
export FASTPIX_PASSWORD="your-secret-key"
```

> Security Note: Never commit your credentials to version control. Use environment variables or secure credential management systems.

## Table of Contents

* [FastPix Go SDK](#fastpix-go-sdk)
  * [Setup](#setup)
  * [Example Usage](#example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Development](#development)

## Setup

### Installation

Install the FastPix Go SDK using Go modules:

```bash
go get github.com/FastPix/fastpix-go
```

### Imports

The SDK uses standard Go packages. Import the necessary packages at the top of your files:

```go
import (
    "context"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)
```

### Initialization

Initialize the FastPix SDK with your credentials:

```go
package main

import (
    "context"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )
}
```

Or using environment variables:

```go
package main

import (
    "context"
    "os"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")), // Your Access Token
            Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")), // Your Secret Key
        }),
    )
}
```

## Example Usage

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"

    "github.com/FastPix/fastpix-go/models/components"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateMediaSuccessResponse != nil {
        // Read raw response body to preserve API's JSON field order
        if res.HTTPMeta.Response != nil && res.HTTPMeta.Response.Body != nil {
            rawBody, err := io.ReadAll(res.HTTPMeta.Response.Body)
            if err == nil && len(rawBody) > 0 {
                var buf bytes.Buffer
                if err := json.Indent(&buf, rawBody, "", "  "); err == nil {
                    fmt.Println(buf.String())
                } else {
                    fmt.Println(string(rawBody))
                }
            } else {
                responseJSON, err := json.MarshalIndent(res.CreateMediaSuccessResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.CreateMediaSuccessResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.CreateMediaSuccessResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.CreateMediaSuccessResponse)
            } else {
                fmt.Println(string(responseJSON))
            }
        }
    } else if res.DefaultError != nil {
        fmt.Printf("Error: %+v\n", res.DefaultError)
    }
}
```

## Available Resources and Operations

Comprehensive Go SDK for FastPix platform integration with full API coverage.

### Media API

Upload, manage, and transform video content with comprehensive media management capabilities.

For detailed documentation, see [FastPix Video on Demand Overview](https://docs.fastpix.io/docs/video-on-demand-overview).

#### Input Video
- [Create from URL](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/inputvideo/README.md#create) - Upload video content from external URL
- [Upload from Device](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/inputvideo/README.md#directuploadmedia) - Upload video files directly from device

#### Manage Videos
- [List All Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#list) - Retrieve complete list of all media files
- [Get Media by ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#get) - Get detailed information for specific media
- [Update Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#update) - Modify media metadata and settings
- [Delete Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#delete) - Remove media files from library
- [Cancel Upload](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#cancelupload) - Stop ongoing media upload process
- [Get Input Info](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#getinputinfo) - Retrieve detailed input information
- [List Uploads](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#listunuseduploadurls) - Get all available upload URLs
- [List Clips](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#getmediaclips) - Get all clips of a media

#### Playback
- [Create Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playback/README.md#create) - Generate secure playback identifier
- [List Playback IDs](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playback/README.md#list) - Get all playback IDs for a media
- [Delete Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playback/README.md#delete) - Remove playback access
- [Get Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playback/README.md#getbyid) - Retrieve playback configuration details
- [Update Domain Restrictions](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playback/README.md#updatedomainrestrictions) - Update domain restrictions for a playback ID
- [Update User-Agent Restrictions](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playbackids/README.md#updateuseragentrestrictions) - Update user-agent restrictions for a playback ID

#### Playlist
- [Create Playlist](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlists/README.md#create) - Create new video playlist
- [List Playlists](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlists/README.md#list) - Get all available playlists
- [Get Playlist](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlist/README.md#get) - Retrieve specific playlist details
- [Update Playlist](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlists/README.md#update) - Modify playlist settings and metadata
- [Delete Playlist](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlists/README.md#delete) - Remove playlist from library
- [Add Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlist/README.md#addmedia) - Add media items to playlist
- [Reorder Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlist/README.md#changemediaorder) - Change order of media in playlist
- [Remove Media](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/playlists/README.md#deletemedia) - Remove media from playlist

#### Signing Keys
- [Create Key](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/signingkeys/README.md#create) - Generate new signing key pair
- [List Keys](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/signingkeys/README.md#list) - Get all available signing keys
- [Delete Key](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/signingkeys/README.md#delete) - Remove signing key from system
- [Get Key](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/signingkeys/README.md#getbyid) - Retrieve specific signing key details

#### DRM Configurations
- [List DRM Configs](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/drmconfigurations/README.md#list) - Get all DRM configuration options
- [Get DRM Config](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/drmconfigurations/README.md#getbyid) - Retrieve specific DRM configuration

### Live API

Stream, manage, and transform live video content with real-time broadcasting capabilities.

For detailed documentation, see [FastPix Live Stream Overview](https://docs.fastpix.io/docs/live-stream-overview).

#### Start Live Stream
- [Create Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/startlivestream/README.md#create) - Initialize new live streaming session

#### Manage Live Stream
- [List Streams](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managelivestream/README.md#list) - Retrieve all active live streams
- [Get Viewer Count](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/livestreams/README.md#getviewercount) - Get real-time viewer statistics
- [Get Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/livestreams/README.md#getbyid) - Retrieve detailed stream information
- [Delete Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/livestreams/README.md#delete) - Terminate and remove live stream
- [Update Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managelivestream/README.md#updatelivestream) - Modify stream settings and configuration
- [Enable Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managelivestream/README.md#enable) - Activate live streaming
- [Disable Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/livestreams/README.md#disable) - Pause live streaming
- [Complete Stream](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managelivestream/README.md#complete) - Finalize and archive stream

#### Live Playback
- [Create Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/liveplayback/README.md#create) - Generate secure live playback access
- [Delete Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/liveplayback/README.md#deleteplaybackid) - Revoke live playback access
- [Get Playback ID](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/liveplayback/README.md#getplaybackiddetails) - Retrieve live playback configuration

#### Simulcast Stream
- [Create Simulcast](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/simulcaststreams/README.md#create) - Set up multi-platform streaming
- [Delete Simulcast](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/simulcaststream/README.md#delete) - Remove simulcast configuration
- [Get Simulcast](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/simulcaststreams/README.md#getspecific) - Retrieve simulcast settings
- [Update Simulcast](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/simulcaststreams/README.md#update) - Modify simulcast parameters

### Video Data API

Monitor video performance and quality with comprehensive analytics and real-time metrics.

For detailed documentation, see [FastPix Video Data Overview](https://docs.fastpix.io/docs/video-data-overview).

#### Metrics
- [List Breakdown Values](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/metrics/README.md#listbreakdownvalues) - Get detailed breakdown of metrics by dimension
- [List Overall Values](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/metrics/README.md#listoverallvalues) - Get aggregated metric values across all content
- [Get Timeseries Data](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/metrics/README.md#gettimeseriesdata) - Retrieve time-based metric trends and patterns
- [Compare Values](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/metrics/README.md#listcomparisonvalues) - List comparison values

#### Views
- [List Video Views](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/views/README.md#listvideoviews) - Get comprehensive list of video viewing sessions
- [Get View Details](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/views/README.md#getdetails) - Retrieve detailed information about specific video views
- [List Top Content](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/views/README.md#listbytopcontent) - Find your most popular and engaging content

#### Dimensions
- [List Dimensions](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/dimensions/README.md#list) - Get available data dimensions for filtering and analysis
- [List Filter Values](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/dimensions/README.md#listfiltervalues) - Get specific values for a particular dimension

#### Errors
- [List Errors](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/errors/README.md#list) - Get list of playback errors

### Transformations

Transform and enhance your video content with powerful AI and editing capabilities.

#### In-Video AI Features

Enhance video content with AI-powered features including moderation, summarization, and intelligent categorization.

- [Generate Summary](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/invideoai/README.md#generatesummary) - Create AI-generated video summaries
- [Update Chapters](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/invideoaifeatures/README.md#updatechapters) - Automatically generate video chapter markers
- [Extract Entities](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/invideoaifeatures/README.md#updatemedianamedentities) - Identify and extract named entities from content
- [Enable Moderation](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/invideoaifeatures/README.md#updatemoderation) - Activate content moderation and safety checks

#### Media Clips

- [List Live Clips](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#listliveclips) - Get all clips of a live stream
- [List Media Clips](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#getmediaclips) - Retrieve all clips associated with a source media

#### Subtitles

- [Generate Subtitles](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#generatesubtitletrack) - Create automatic subtitles for media

#### Media Tracks

- [Add Track](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#addmediatrack) - Add audio or subtitle tracks to media
- [Update Track](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/managevideos/README.md#updatetrack) - Modify existing audio or subtitle tracks
- [Delete Track](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#deletetrack) - Remove audio or subtitle tracks

#### Access Control

- [Update Source Access](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#updatesourceaccess) - Control access permissions for media source

#### Format Support

- [Update MP4 Support](https://github.com/FastPix/fastpix-go/blob/feature/fixed-missing-parameters/docs/sdks/videos/README.md#updatemp4support) - Configure MP4 download capabilities

<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:

```go
package main

import (
    "context"
    "log"

    "github.com/FastPix/fastpix-go/models/components"
    "github.com/FastPix/fastpix-go/models/operations"
    "github.com/FastPix/fastpix-go/retry"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    }, operations.WithRetries(
        retry.Config{
            Strategy: "backoff",
            Backoff: &retry.BackoffStrategy{
                InitialInterval: 1,
                MaxInterval:     50,
                Exponent:        1.1,
                MaxElapsedTime:  100,
            },
            RetryConnectionErrors: false,
        }))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateMediaSuccessResponse != nil {
        // handle response
    }
}
```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:

```go
package main

import (
    "context"
    "log"

    "github.com/FastPix/fastpix-go/models/components"
    "github.com/FastPix/fastpix-go/retry"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithRetryConfig(
            retry.Config{
                Strategy: "backoff",
                Backoff: &retry.BackoffStrategy{
                    InitialInterval: 1,
                    MaxInterval:     50,
                    Exponent:        1.1,
                    MaxElapsedTime:  100,
                },
                RetryConnectionErrors: false,
            }),
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateMediaSuccessResponse != nil {
        // handle response
    }
}
```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| apierrors.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
    "context"
    "errors"
    "log"

    "github.com/FastPix/fastpix-go/models/apierrors"
    "github.com/FastPix/fastpix-go/models/components"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    })
    if err != nil {
        var e *apierrors.APIError
        if errors.As(err, &e) {
            // handle error
            log.Fatal(e.Error())
        }
    }
    if res.CreateMediaSuccessResponse != nil {
        // handle response
    }
}
```

### Error Classes

**Primary exception:**
* [`apierrors.APIError`](./models/apierrors/): The base class for HTTP error responses.

<details><summary>Less common exceptions</summary>

* Network connectivity errors: These are typically returned as `apierrors.APIError` with appropriate status codes. For network-level errors, check the underlying error message.

* Response validation errors: When the response data could not be deserialized into the expected type, the SDK will return an `apierrors.APIError` with details about the validation failure.
</details>
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:

```go
package main

import (
    "context"
    "log"

    "github.com/FastPix/fastpix-go/models/components"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithServerURL("your-server-url"),
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateMediaSuccessResponse != nil {
        // handle response
    }
}
```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls using an HTTP client that wraps the standard `net/http` package. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout:

```go
package main

import (
    "context"
    "net/http"
    "time"

    "github.com/FastPix/fastpix-go/models/components"
    fastpixgo "github.com/FastPix/fastpix-go"
)

func main() {
    ctx := context.Background()

    httpClient := &http.Client{Timeout: 30 * time.Second}
    s := fastpixgo.New(
        fastpixgo.WithClient(httpClient),
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.InputVideo.Create(ctx, components.CreateMediaRequest{
        Inputs: []components.Input{
            components.CreateInputPullVideoInput(
                components.PullVideoInput{},
            ),
        },
        Metadata: map[string]string{
            "your-metadata-key": "your-metadata-value",
        },
    })
    if err != nil {
        // handle error
    }
    if res.CreateMediaSuccessResponse != nil {
        // handle response
    }
}
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.

<details>
<summary>For simple debugging, you can enable request/response logging by implementing a custom client:</summary>

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/FastPix/fastpix-go/models/components"
    fastpixgo "github.com/FastPix/fastpix-go"
)

type LoggingHTTPClient struct {
    client *http.Client
}

func (c *LoggingHTTPClient) Do(req *http.Request) (*http.Response, error) {
    // Log request
    fmt.Printf("Sending %s request to %s\n", req.Method, req.URL.String())
    
    resp, err := c.client.Do(req)
    
    // Log response
    if err != nil {
        fmt.Printf("Request failed: %v\n", err)
    } else {
        fmt.Printf("Received %s response\n", resp.Status)
    }
    
    return resp, err
}

func main() {
    ctx := context.Background()

    loggingClient := &LoggingHTTPClient{
        client: &http.Client{Timeout: 30 * time.Second},
    }
    
    s := fastpixgo.New(
        fastpixgo.WithClient(loggingClient),
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    // Use SDK as normal
}
```
</details>
<!-- End Custom HTTP Client [http-client] -->

# Development

This Go SDK is programmatically generated from our API specifications. Any manual modifications to internal files will be overwritten during subsequent generation cycles. 

We value community contributions and feedback. Feel free to submit pull requests or open issues with your suggestions, and we'll do our best to include them in future releases.

## Detailed Usage

For comprehensive understanding of each API's functionality, including detailed request and response specifications, parameter descriptions, and additional examples, please refer to the [FastPix API Reference](https://docs.fastpix.io/reference/signingkeys-overview).

The API reference offers complete documentation for all available endpoints and features, enabling developers to integrate and leverage FastPix APIs effectively.
