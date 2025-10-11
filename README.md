# FastPix Go SDK

A robust, type-safe Go SDK designed for seamless integration with the FastPix API platform.



<!-- Start Summary [summary] -->
## Introduction

The FastPix Go SDK simplifies integration with the FastPix platform. It provides a clean, Go interface for secure and efficient communication with the FastPix API, enabling easy management of media uploads, live streaming, on‑demand content, playlists, video analytics, and signing keys for secure access and token management. It is intended for use with Go 1.22 and above.

## Prerequisites

### Environment and Version Support

<table>
<tr>
<th>Requirement</th>
<th>Version</th>
<th>Description</th>
</tr>
<tr>
<td><strong>Go</strong></td>
<td><code>1.22+</code></td>
<td>Core runtime environment</td>
</tr>
<tr>
<td><strong>Go Modules</strong></td>
<td><code>Enabled</code></td>
<td>Dependency management for Go packages</td>
</tr>
<tr>
<td><strong>Internet</strong></td>
<td><code>Required</code></td>
<td>API communication and authentication</td>
</tr>
</table>

> **Pro Tip:** We recommend using Go 1.23+ for optimal performance and the latest language features.

### Getting Started with FastPix

To get started with the **FastPix Go SDK**, ensure you have the following:

- The FastPix APIs are authenticated using a **Username** and a **Password**. You must generate these credentials to use the SDK.

- Follow the steps in the [Authentication with Basic Auth](https://docs.fastpix.io/docs/basic-authentication) guide to obtain your credentials.

### Environment Variables (Optional)

Configure your FastPix credentials using environment variables for enhanced security and convenience:

```bash
# Set your FastPix credentials
export FASTPIX_USERNAME="your-access-token"
export FASTPIX_PASSWORD="your-secret-key"
```

> **Security Note:** Never commit your credentials to version control. Use environment variables or secure credential management systems.

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Fastpixgo](#fastpixgo)
  * [Setup](#setup)
  * [Example Usage](#example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Development](#development)

<!-- End Table of Contents [toc] -->

<!-- Start Setup [setup] -->
## Setup

### Installation

Install the FastPix Go SDK using Go modules:

```bash
go get github.com/FastPix/fastpix-go
```

### Imports

Import the necessary modules for your FastPix integration:

```go
// Basic imports
import (
    "context"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)
```

### Initialization

Initialize the FastPix SDK with your credentials:

```go
import (
    "context"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)

client := fastpix.New(
    fastpix.WithSecurity(components.Security{
        Username: fastpix.Pointer("your-access-token"),
        Password: fastpix.Pointer("your-secret-key"),
    }),
)
```

Or using environment variables:

```go
import (
    "context"
    "os"
    "github.com/FastPix/fastpix-go"
    "github.com/FastPix/fastpix-go/models/components"
)

client := fastpix.New(
    fastpix.WithSecurity(components.Security{
        Username: fastpix.Pointer(os.Getenv("FASTPIX_USERNAME")),
        Password: fastpix.Pointer(os.Getenv("FASTPIX_PASSWORD")),
    }),
)
```

<!-- End Setup [setup] -->

<!-- Start Example Usage [example-usage] -->
## Example Usage

### Example

```go
package main

import (
	"context"
	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpix.New(
		fastpix.WithSecurity(components.Security{
			Username: fastpix.Pointer("your-access-token"),
			Password: fastpix.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateMediaSuccessResponse != nil {
		// handle response
	}
}

```
<!-- End Example Usage [example-usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

Comprehensive Go SDK for FastPix platform integration with full API coverage.

### Media API

Upload, manage, and transform video content with comprehensive media management capabilities.

For detailed documentation, see [FastPix Video on Demand Overview](https://docs.fastpix.io/docs/video-on-demand-overview).

#### Input Video
- [Create from URL](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/inputvideo/README.md#createmedia) - Upload video content from external URL
- [Upload from Device](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/inputvideo/README.md#directuploadvideomedia) - Upload video files directly from device

#### Manage Videos
- [List All Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#listmedia) - Retrieve complete list of all media files
- [Get Media by ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#getmedia) - Get detailed information for specific media
- [Update Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#updatedmedia) - Modify media metadata and settings
- [Delete Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#deletemedia) - Remove media files from library
- [Add Track](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#addmediatrack) - Add audio or subtitle tracks to media
- [Cancel Upload](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#cancelupload) - Stop ongoing media upload process
- [Update Track](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#updatemediatrack) - Modify existing audio or subtitle tracks
- [Delete Track](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#deletemediatrack) - Remove audio or subtitle tracks
- [Generate Subtitles](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#generatesubtitletrack) - Create automatic subtitles for media
- [Update Source Access](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#updatedsourceaccess) - Control access permissions for media source
- [Update MP4 Support](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#updatedmp4support) - Configure MP4 download capabilities
- [Get Input Info](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#retrievemediainputinfo) - Retrieve detailed input information
- [List Uploads](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#listuploads) - Get all available upload URLs
- [Get Media Clips](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managevideos/README.md#getmediaclips) - Retrieve all video clips for media

#### Playback
- [Create Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playback/README.md#createmediaplaybackid) - Generate secure playback identifier
- [Delete Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playback/README.md#deletemediaplaybackid) - Remove playback access
- [Get Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playback/README.md#getplaybackid) - Retrieve playback configuration details

#### Playlist
- [Create Playlist](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#createaplaylist) - Create new video playlist
- [List Playlists](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#getallplaylists) - Get all available playlists
- [Get Playlist](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#getplaylistbyid) - Retrieve specific playlist details
- [Update Playlist](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#updateaplaylist) - Modify playlist settings and metadata
- [Delete Playlist](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#deleteaplaylist) - Remove playlist from library
- [Add Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#addmediatoplaylist) - Add media items to playlist
- [Reorder Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#changemediaorderinplaylist) - Change order of media in playlist
- [Remove Media](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/playlist/README.md#deletemediafromplaylist) - Remove media from playlist

#### Signing Keys
- [Create Key](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/signingkeys/README.md#createsigningkey) - Generate new signing key pair
- [List Keys](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/signingkeys/README.md#listsigningkeys) - Get all available signing keys
- [Delete Key](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/signingkeys/README.md#deletesigningkey) - Remove signing key from system
- [Get Key](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/signingkeys/README.md#getsigningkeybyid) - Retrieve specific signing key details

#### DRM Configurations
- [List DRM Configs](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/drmconfigurations/README.md#getdrmconfiguration) - Get all DRM configuration options
- [Get DRM Config](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/drmconfigurations/README.md#getdrmconfigurationbyid) - Retrieve specific DRM configuration

### Live API 

Stream, manage, and transform live video content with real-time broadcasting capabilities.

For detailed documentation, see [FastPix Live Stream Overview](https://docs.fastpix.io/docs/live-stream-overview).

#### Start Live Stream
- [Create Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/startlivestream/README.md#createnewstream) - Initialize new live streaming session

#### Manage Live Stream
- [List Streams](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#getallstreams) - Retrieve all active live streams
- [Get Viewer Count](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#getlivestreamviewercountbyid) - Get real-time viewer statistics
- [Get Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#getlivestreambyid) - Retrieve detailed stream information
- [Delete Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#deletelivestream) - Terminate and remove live stream
- [Update Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#updatelivestream) - Modify stream settings and configuration
- [Enable Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#enablelivestream) - Activate live streaming
- [Disable Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#disablelivestream) - Pause live streaming
- [Complete Stream](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/managelivestream/README.md#completelivestream) - Finalize and archive stream

#### Live Playback
- [Create Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/liveplayback/README.md#createplaybackidofstream) - Generate secure live playback access
- [Delete Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/liveplayback/README.md#deleteplaybackidofstream) - Revoke live playback access
- [Get Playback ID](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/liveplayback/README.md#getlivestreamplaybackid) - Retrieve live playback configuration

#### Simulcast Stream
- [Create Simulcast](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/simulcaststream/README.md#createsimulcastofstream) - Set up multi-platform streaming
- [Delete Simulcast](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/simulcaststream/README.md#deletesimulcastofstream) - Remove simulcast configuration
- [Get Simulcast](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/simulcaststream/README.md#getspecificsimulcastofstream) - Retrieve simulcast settings
- [Update Simulcast](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/simulcaststream/README.md#updatespecificsimulcastofstream) - Modify simulcast parameters

### Video Data API 

Monitor video performance and quality with comprehensive analytics and real-time metrics.

For detailed documentation, see [FastPix Video Data Overview](https://docs.fastpix.io/docs/video-data-overview).

#### Metrics
- [List Breakdown Values](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/metrics/README.md#listbreakdownvalues) - Get detailed breakdown of metrics by dimension
- [List Overall Values](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/metrics/README.md#listoverallvalues) - Get aggregated metric values across all content
- [Get Timeseries Data](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/metrics/README.md#gettimeseriesdata) - Retrieve time-based metric trends and patterns
- [List Comparison Values](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/metrics/README.md#listcomparisonvalues) - Compare metrics across different time periods

#### Views
- [List Video Views](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/views/README.md#listvideoviews) - Get comprehensive list of video viewing sessions
- [Get View Details](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/views/README.md#getvideoviewdetails) - Retrieve detailed information about specific video views
- [List Top Content](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/views/README.md#listbytopcontent) - Find your most popular and engaging content
- [Get Concurrent Viewers](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/views/README.md#getdataviewlistcurrentviewsgettimeseriesviews) - Monitor real-time viewer counts over time
- [Get Viewer Breakdown](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/views/README.md#getdataviewlistcurrentviewsfilter) - Analyze viewers by device, location, and other dimensions

#### Dimensions
- [List Dimensions](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/dimensions/README.md#listdimensions) - Get available data dimensions for filtering and analysis
- [List Filter Values](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/dimensions/README.md#listfiltervaluesfordimension) - Get specific values for a particular dimension

### In-Video AI Features

Enhance video content with AI-powered features including moderation, summarization, and intelligent categorization.

For detailed documentation, see [Video Moderation Guide](https://docs.fastpix.io/docs/using-nsfw-and-profanity-filter-for-video-moderation).

- [Generate Summary](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/invideoaifeatures/README.md#updatemediasummary) - Create AI-generated video summaries
- [Create Chapters](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/invideoaifeatures/README.md#updatemediachapters) - Automatically generate video chapter markers
- [Extract Entities](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/invideoaifeatures/README.md#updatemedianamedentities) - Identify and extract named entities from content
- [Enable Moderation](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/invideoaifeatures/README.md#updatemediamoderation) - Activate content moderation and safety checks

### Error Handling

Handle and manage errors with comprehensive error handling capabilities and detailed error information for all API operations.

- [List Errors](https://github.com/FastPix/fastpix-go/blob/main/docs/sdks/errors/README.md#listerrors) - Retrieve comprehensive error logs and diagnostics

<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := fastpix.New(
		fastpix.WithSecurity(components.Security{
			Username: fastpix.Pointer("your-access-token"),
			Password: fastpix.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
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
	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpix.New(
		fastpix.WithRetryConfig(
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
		fastpix.WithSecurity(components.Security{
			Username: fastpix.Pointer("your-access-token"),
			Password: fastpix.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
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

For example, the `CreateMedia` function may return the following errors:

| Error Type                        | Status Code | Content Type     |
| --------------------------------- | ----------- | ---------------- |
| apierrors.BadRequestError         | 400         | application/json |
| apierrors.InvalidPermissionError  | 401         | application/json |
| apierrors.ForbiddenError          | 403         | application/json |
| apierrors.ValidationErrorResponse | 422         | application/json |
| apierrors.APIError                | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpix.New(
		fastpix.WithSecurity(components.Security{
			Username: fastpix.Pointer("your-access-token"),
			Password: fastpix.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
	})
	if err != nil {

		var e *apierrors.BadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.InvalidPermissionError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ForbiddenError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ValidationErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	fastpix "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpix.New(
		fastpix.WithServerURL("https://api.fastpix.io/v1/"),
		fastpix.WithSecurity(components.Security{
			Username: fastpix.Pointer("your-access-token"),
			Password: fastpix.Pointer("your-secret-key"),
		}),
	)

	res, err := s.InputVideo.CreateMedia(ctx, components.CreateMediaRequest{
		Inputs: []components.Input{
			components.CreateInputVideoInput(
				components.VideoInput{
					Type: "video",
					URL:  "https://static.fastpix.io/sample.mp4",
				},
			),
		},
		Metadata: map[string]string{
			"key1": "value1",
		},
		AccessPolicy: components.CreateMediaRequestAccessPolicyPublic,
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

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/FastPix/fastpix-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = fastpix.New(fastpix.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future fastpix SDK Sections -->

# Development

This Go SDK is programmatically generated from our API specifications. Any manual modifications to internal files will be overwritten during subsequent generation cycles. 

We value community contributions and feedback. Feel free to submit pull requests or open issues with your suggestions, and we'll do our best to include them in future releases.

## Detailed Usage

For comprehensive understanding of each API's functionality, including detailed request and response specifications, parameter descriptions, and additional examples, please refer to the [FastPix API Reference](https://docs.fastpix.io/reference/signingkeys-overview).

The API reference offers complete documentation for all available endpoints and features, enabling developers to integrate and leverage FastPix APIs effectively.



<style>
  :root {
    --badge-gray-bg: #f3f4f6;
    --badge-gray-border: #d1d5db;
    --badge-gray-text: #374151;
    --badge-blue-bg: #eff6ff;
    --badge-blue-border: #3b82f6;
    --badge-blue-text: #3b82f6;
  }

  @media (prefers-color-scheme: dark) {
    :root {
      --badge-gray-bg: #374151;
      --badge-gray-border: #4b5563;
      --badge-gray-text: #f3f4f6;
      --badge-blue-bg: #1e3a8a;
      --badge-blue-border: #3b82f6;
      --badge-blue-text: #93c5fd;
    }
  }
  
  h1 {
    border-bottom: none !important;
    margin-bottom: 4px;
    margin-top: 0;
    letter-spacing: 0.5px;
    font-weight: 600;
  }
  
  .badge-text {
    letter-spacing: 1px;
    font-weight: 300;
  }
  
  .badge-container {
    display: inline-flex;
    align-items: center;
    background: var(--badge-gray-bg);
    border: 1px solid var(--badge-gray-border);
    border-radius: 6px;
    overflow: hidden;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif;
    font-size: 11px;
    text-decoration: none;
    vertical-align: middle;
  }

  .badge-container.blue {
    background: var(--badge-blue-bg);
    border-color: var(--badge-blue-border);
  }

  .badge-icon-section {
    padding: 4px 8px;
    border-right: 1px solid var(--badge-gray-border);
    display: flex;
    align-items: center;
  }

  .badge-text-section {
    padding: 4px 10px;
    color: var(--badge-gray-text);
    font-weight: 400;
  }

  .badge-container.blue .badge-text-section {
    color: var(--badge-blue-text);
  }
  
  .badge-link {
    text-decoration: none;
    margin-left: 8px;
    display: inline-flex;
    vertical-align: middle;
  }

  .badge-link:hover {
    text-decoration: none;
  }
  
  .badge-link:first-child {
    margin-left: 0;
  }
  
  .badge-icon-section svg {
    color: var(--badge-gray-text);
  }

  .badge-container.blue .badge-icon-section svg {
    color: var(--badge-blue-text);
  }
</style> 