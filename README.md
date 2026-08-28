# FastPix Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/FastPix/fastpix-go.svg)](https://pkg.go.dev/github.com/FastPix/fastpix-go)
[![license](https://img.shields.io/github/license/FastPix/fastpix-go)](https://github.com/FastPix/fastpix-go/blob/main/LICENSE)
[![Go 1.21+](https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)

A robust, type-safe Go SDK designed for seamless integration with the FastPix API platform.

The FastPix Go SDK is a strongly typed Go client for the FastPix video API. From any Go application, you can upload and manage videos, run live streams and simulcasts, create and secure playback IDs, manage playlists and signing keys, pull video analytics, and use in-video AI features.

**Supported Go:** 1.21 and later
**Package:** `github.com/FastPix/fastpix-go`
**Authentication:** HTTP Basic Authentication
**Dependency management:** Go Modules

📖 **Docs:** https://fastpix.com/docs/language-sdks/go-sdk &nbsp;·&nbsp; 🚀 **Free account:** https://dashboard.fastpix.com

<br />

## Start here

If you are using the FastPix Go SDK for the first time, follow these steps in order:

1. [Check your Go version](#1-check-your-go-version)
2. [Create a Go project](#2-create-a-go-project)
3. [Install the SDK](#3-install-the-sdk)
4. [Verify the installation](#4-verify-the-installation)
5. [Configure authentication](#5-configure-authentication)
6. [Initialize the FastPix client](#6-initialize-the-fastpix-client)
7. [Make your first API request](#7-make-your-first-api-request)
8. [Verify the API response](#8-verify-the-api-response)
9. [Retrieve the media asset](#9-retrieve-the-media-asset)
10. [Verify the media ID](#10-verify-the-media-id)

Do not skip the verification steps. If installation, authentication, or client initialization fails, troubleshoot that problem before continuing to the next API operation.

---

### Before you begin

To use the SDK, make sure you have:

- Go 1.21 or later.
- Go Modules enabled.
- Internet access.
- A FastPix account.
- A FastPix Access Token.
- A FastPix Secret Key.

FastPix uses Basic Authentication:

| SDK value | FastPix credential |
|---|---|
| `Username` | Access Token |
| `Password` | Secret Key |

You can obtain your credentials from the FastPix Dashboard. Follow the [Authentication with Basic Auth](https://fastpix.com/docs/getting-started/activate-your-account) guide to obtain your credentials.

## 1. Check your Go version

Run:

```bash
go version
```

Output is similar to:

```text
go version go1.21.0 darwin/arm64
```

or a later version.

The FastPix Go SDK supports Go 1.21 and later.

If your Go version is earlier than 1.21, install a supported version before continuing.

---

## 2. Create a Go project

### a. Create a new directory for your FastPix application

```bash
mkdir fastpix-go-demo
cd fastpix-go-demo
```

### b. Initialize a Go module

Run:

```bash
go mod init fastpix-go-demo
```

Output is similar to:

```text
go: creating new go.mod: module fastpix-go-demo
```

This creates a `go.mod` file that tracks your project's module name and dependencies.

Your project should now contain:

```text
fastpix-go-demo/
└── go.mod
```

---

## 3. Install the SDK

Install the FastPix Go SDK using Go Modules:

```bash
go get github.com/FastPix/fastpix-go
```

Go downloads the SDK and adds it to your project's dependencies.

Verify the dependency:

```bash
go list -m github.com/FastPix/fastpix-go
```

Output is similar to:

```text
github.com/FastPix/fastpix-go v<version>
```

You can also inspect the `go.mod` file:

```bash
cat go.mod
```

The FastPix SDK should be listed as a dependency.

## 4. Verify the installation

Before making an API request, verify that your Go project can import the SDK.

Create a file named `main.go`:

```go
package main
import (
	"fmt"
	fastpixgo "github.com/FastPix/fastpix-go"
)
func main() {
	_ = fastpixgo.New
	fmt.Println("FastPix SDK imported successfully")
}
```

Run:

```bash
go run .
```

Output is similar to:

```text
FastPix SDK imported successfully
```

If this command fails, do not continue to API calls.

Check:

- Go 1.21 or later is installed.
- The Go module has been initialized.
- `github.com/FastPix/fastpix-go` is listed in `go.mod`.
- You are running the command from the project directory.
- Your Go environment is configured correctly.

You can download or update dependencies with:

```bash
go mod tidy
```

Then run the verification again:

```bash
go run .
```

## 5. Configure authentication

FastPix uses Basic Authentication.

Set the Access Token and Secret Key as environment variables.

### macOS and Linux

```bash
export FASTPIX_USERNAME="<YOUR_ACCESS_TOKEN>"
export FASTPIX_PASSWORD="<YOUR_SECRET_KEY>"
```

### Windows PowerShell

```powershell
$env:FASTPIX_USERNAME="<YOUR_ACCESS_TOKEN>"
$env:FASTPIX_PASSWORD="<YOUR_SECRET_KEY>"
```

The SDK maps these variables as follows:

```text
FASTPIX_USERNAME → Access Token
FASTPIX_PASSWORD → Secret Key
```

### Verify the credentials are set

Do not print the actual credential values.

Instead, run:

### macOS and Linux

```bash
if [ -n "$FASTPIX_USERNAME" ]; then
  echo "Access Token: set"
else
  echo "Access Token: missing"
fi
if [ -n "$FASTPIX_PASSWORD" ]; then
  echo "Secret Key: set"
else
  echo "Secret Key: missing"
fi
```

Output is similar to:

```text
Access Token: set
Secret Key: set
```

### Security

Never:

- Commit credentials to Git.
- Put credentials directly into source code.
- Include credentials in screenshots, logs, or bug reports.
- Print authentication headers during debugging in production.

Use environment variables or a secure credential-management system.

## 6. Initialize the FastPix client

Create or replace `main.go` with:

```go
package main
import (
	"context"
	"fmt"
	"os"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)
func main() {
	ctx := context.Background()
	fastpix := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
	)
	_ = ctx
	_ = fastpix
	fmt.Println("FastPix client initialized")
}
```

Run:

```bash
go run .
```

Output is similar to:

```text
FastPix client initialized
```

### What this code does

`fastpixgo.New()` creates the top-level FastPix SDK client.

`components.Security` contains the credentials used to authenticate API requests.

The SDK client does not make an API request simply because it is initialized.

An API request occurs when you call an operation such as:

```go
fastpix.InputVideo.Create(...)
```

## 7. Make your first API request

The easiest way to verify the complete integration is to create media from a publicly accessible video URL.

FastPix provides a sample video URL:

```text
https://static.fastpix.com/fp-sample-video.mp4
```

The SDK uses the `InputVideo.Create` operation to create media from an external video URL.

Replace `main.go` with:

```go
package main
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)
func main() {
	ctx := context.Background()
	fastpix := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
	)
	response, err := fastpix.InputVideo.Create(
		ctx,
		components.CreateMediaRequest{
			Inputs: []components.Input{
				components.CreateInputPullVideoInput(
					components.PullVideoInput{
						URL: fastpixgo.Pointer(
							"https://static.fastpix.com/fp-sample-video.mp4",
						),
					},
				),
			},
			Metadata: map[string]string{
				"source": "fastpix-go-demo",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	if response.CreateMediaSuccessResponse == nil {
		log.Fatal("FastPix API did not return a successful media response")
	}
	output, err := json.MarshalIndent(
		response.CreateMediaSuccessResponse,
		"",
		"  ",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
```

Run:

```bash
go run .
```

A successful request returns information about the newly created media asset.

## 8. Verify the API response

A successful request returns a media ID.

The response contains a structure similar to:

```json
{
  "success": true,
  "data": {
    "id": "..."
  }
}
```

The value of:

```text
data.id
```

is the unique ID assigned to the media.

The exact response fields depend on the API response and SDK version.

Save the media ID

You will need the media ID for subsequent media operations.

For example:

```text
MEDIA_ID=<value returned in data.id>
```

Do not confuse a `media_id` with a `playback_id`.

They identify different resources and are used for different operations.

## 9. Retrieve the media asset

Use the media ID returned by the create operation to retrieve the media asset.

The Go SDK exposes the operation through:

```go
fastpix.Videos.Get(ctx, mediaID)
```

The `mediaID` must be the ID returned when the media was created.

Update `main.go` to retrieve the media after creating it:

```go
package main
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)
func main() {
	ctx := context.Background()
	fastpix := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.Pointer(os.Getenv("FASTPIX_USERNAME")),
			Password: fastpixgo.Pointer(os.Getenv("FASTPIX_PASSWORD")),
		}),
	)
	// Create media.
	fmt.Println("Creating media...")
	createResponse, err := fastpix.InputVideo.Create(
		ctx,
		components.CreateMediaRequest{
			Inputs: []components.Input{
				components.CreateInputPullVideoInput(
					components.PullVideoInput{
						URL: fastpixgo.Pointer(
							"https://static.fastpix.com/fp-sample-video.mp4",
						),
					},
				),
			},
			Metadata: map[string]string{
				"source": "fastpix-go-demo",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	if createResponse.CreateMediaSuccessResponse == nil {
		log.Fatal("FastPix API did not return a successful media response")
	}
	createData := createResponse.CreateMediaSuccessResponse
	output, err := json.MarshalIndent(createData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCREATE MEDIA")
	fmt.Println(string(output))
	// Get the media ID.
	//
	// The generated response type exposes the media data.
	// Extract the ID returned by the create operation here.
	mediaID := createData.Data.Id
	fmt.Println("\nMEDIA ID:")
	fmt.Println(mediaID)
	// Retrieve the media.
	fmt.Println("\nRetrieving media...")
	mediaResponse, err := fastpix.Videos.Get(ctx, mediaID)
	if err != nil {
		log.Fatal(err)
	}
	if mediaResponse.Object == nil {
		log.Fatal("FastPix API did not return media data")
	}
	mediaOutput, err := json.MarshalIndent(
		mediaResponse.Object,
		"",
		"  ",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nGET MEDIA")
	fmt.Println(string(mediaOutput))
}
```

Run:

```bash
go run .
```

The output should contain:

```text
Creating media...
CREATE MEDIA
{
  ...
}
MEDIA ID:
<media-id>
Retrieving media...
GET MEDIA
{
  ...
}
```

The Go SDK's `Videos.Get` operation accepts the media ID as a required string and returns a `GetMediaResponse`. The generated SDK documentation describes this operation as retrieving detailed information about a specific media item, including its current status.

## 10. Verify the media ID

The `media_id` returned by the create operation identifies the media asset.

The same ID is used to retrieve the media asset:

<Image alt="FastPix Go media ID hand-off: the ID returned by create (data.id) becomes mediaID, which you pass to Videos.Get() to retrieve the media asset." border={false} src="https://static.fastpix.com/go-media-id-flow.png" />

A successful response from the get-media operation confirms that:

- The SDK was installed successfully.
- The Go application can import the SDK.
- Your FastPix credentials are configured.
- The FastPix client was initialized successfully.
- Your application authenticated with the FastPix API.
- You created a media asset.
- FastPix returned a media ID.
- The media ID can be used in a subsequent API operation.
- The Go SDK can retrieve the media asset.

At this point, the initial SDK integration is complete.

## What you have verified

By completing this guide, you have verified that:

- Go 1.21 or later is installed.
- Go Modules are configured.
- The FastPix Go SDK is installed.
- Your Go application can import the SDK.
- Your FastPix credentials are configured.
- The FastPix client can be initialized.
- Your application can authenticate with the FastPix API.
- You can create a media asset.
- You can retrieve the media asset using its media ID.

Your completed workflow is:

<Image alt="FastPix Go media workflow: a Go application calls the FastPix Go SDK, which calls the FastPix API to create media, returns a media ID, then retrieves the media with Videos.Get." border={false} src="https://static.fastpix.com/go-media-workflow.png" />

You are now ready to use the returned `media_id` with other FastPix API operations.

<br />

## Available Resources and Operations

Comprehensive Go SDK for FastPix platform integration with full API coverage.

### Media API

Upload, manage, and transform video content with comprehensive media management capabilities.

For detailed documentation, see [FastPix Video on Demand Overview](https://fastpix.com/docs/video-on-demand-api/overview).

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

For detailed documentation, see [FastPix Live Stream Overview](https://fastpix.com/docs/live-stream-api/overview).

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

For detailed documentation, see [FastPix Video Data Overview](https://fastpix.com/docs/video-data-api/overview).

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
            Username: fastpixgo.Pointer("your access-token"),
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
            Username: fastpixgo.Pointer("your access-token"),
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
            Username: fastpixgo.Pointer("your access-token"),
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
            Username: fastpixgo.Pointer("your access-token"),
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
            Username: fastpixgo.Pointer("your access-token"),
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
            Username: fastpixgo.Pointer("your access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )
    // Use SDK as normal
}
```
</details>
<!-- End Custom HTTP Client [http-client] -->

## FAQ

**How do I install the FastPix Go SDK?**
Run `go get github.com/FastPix/fastpix-go` and import it in your module. See [Install the SDK](#3-install-the-sdk).

**How do I authenticate the SDK?**
FastPix uses Basic Auth: build the client with `components.Security` passing your access token as `Username` and secret key as `Password`. See [Initialize the FastPix client](#6-initialize-the-fastpix-client).

**How do I upload a video in Go?**
Create media from a URL or a direct upload through `s.InputVideo`, for example `s.InputVideo.Create(ctx, ...)`. See [Make your first API request](#7-make-your-first-api-request) and [Available Resources and Operations](#available-resources-and-operations).

**How do I start a live stream?**
Use the Live API resources to create and manage streams, simulcasts, and live playback IDs. See [Available Resources and Operations](#available-resources-and-operations).

**How do I get video analytics and metrics in Go?**
The Video Data API exposes metrics, views, dimensions, and errors for quality-of-experience monitoring. See [Available Resources and Operations](#available-resources-and-operations).

**How do I handle API errors?**
API errors are returned as `apierrors.APIError`; use `errors.As` to inspect the status code and body. See [Error Handling](#error-handling).

**How do I configure automatic retries?**
Pass `operations.WithRetries(...)` per call or `fastpixgo.WithRetryConfig(...)` at initialization to control the backoff strategy. See [Retries](#retries).

**How do I use a custom HTTP client, proxy, or timeout?**
Provide any client that satisfies the SDK's `HTTPClient` interface (for example an `*http.Client` with a timeout) via `fastpixgo.WithClient(...)`. See [Custom HTTP Client](#custom-http-client).

**How do I change the API base URL?**
Use `fastpixgo.WithServerURL(...)` when constructing the client. See [Server Selection](#server-selection).

**Which Go versions are supported?**
Go 1.21 and above. See [Before you begin](#before-you-begin).

## Which FastPix SDK should I use?

FastPix publishes a server SDK for every major backend language, each generated from the same API specification:

| Language | Repo | Install |
|---|---|---|
| **Go** (this repo) | [fastpix-go](https://github.com/FastPix/fastpix-go) | `go get github.com/FastPix/fastpix-go` |
| Node.js / TypeScript | [node-sdk](https://github.com/FastPix/node-sdk) | `npm install @fastpix/fastpix-node` |
| Python | [fastpix-python](https://github.com/FastPix/fastpix-python) | `pip install fastpix-python` |
| PHP | [fastpix-php](https://github.com/FastPix/fastpix-php) | `composer require fastpix/sdk` |
| Java | [fastpix-java](https://github.com/FastPix/fastpix-java) | `io.fastpix:sdk` (Maven/Gradle) |
| C# / .NET | [fastpix-sdk-csharp](https://github.com/FastPix/fastpix-sdk-csharp) | `dotnet add package Fastpix` |
| Ruby | [fastpix-ruby](https://github.com/FastPix/fastpix-ruby) | `gem install fastpixapi` |

To upload and play the media these SDKs create, use the FastPix browser libraries: [web-uploads-sdk](https://github.com/FastPix/web-uploads-sdk), [react-web-uploader](https://github.com/FastPix/react-web-uploader), and [web-player-component](https://github.com/FastPix/web-player-component). Browse everything in the [FastPix organization](https://github.com/orgs/FastPix/repositories).

## Development

This Go SDK is programmatically generated from our API specifications. Any manual modifications to internal files will be overwritten during subsequent generation cycles. 

We value community contributions and feedback. Feel free to submit pull requests or open issues with your suggestions, and we'll do our best to include them in future releases.

## Detailed Usage

For comprehensive understanding of each API's functionality, including detailed request and response specifications, parameter descriptions, and additional examples, please refer to the [FastPix API Reference](https://fastpix.com/docs/product-os-api/overview).

The API reference offers complete documentation for all available endpoints and features, enabling developers to integrate and leverage FastPix APIs effectively.
