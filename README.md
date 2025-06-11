# FastPix Go SDK

[**Homepage**](https://fastpix.io)

Developer-friendly & type-safe Go SDK specifically designed to leverage the FastPix platform API.

## Introduction

The FastPix Go SDK simplifies integration with the FastPix platform. This SDK is designed for secure and efficient communication with the FastPix API, enabling easy management of media uploads, live streaming, and simulcasting.

## Key Features

- **Media API**
  - **Upload Media**: Upload media files seamlessly from URLs or devices
  - **Manage Media**: List, fetch, update, and delete media assets
  - **Playback IDs**: Generate and manage playback IDs for media access
- **Live API**
  - **Create & Manage Live Streams**: Create, list, update, and delete live streams
  - **Control Stream Access**: Generate playback IDs for live streams to control and manage access
  - **Simulcast to Multiple Platforms**: Stream content to multiple platforms simultaneously

For detailed usage, refer to the [FastPix API Reference](docs/).

## Prerequisites

- Go 1.20 or later
- FastPix API credentials (Username and Password)

## Getting Started with FastPix

To get started with the **FastPix Go SDK**, ensure you have the following:

- FastPix APIs are authenticated using a **Username** and **Password** (HTTP Basic Auth). You must generate these credentials to use the SDK.
- Follow the steps in the Authentication section to obtain your credentials.

## Table of Contents

- [SDK Installation](#sdk-installation)
- [Initialization](#initialization)
- [SDK Example Usage](#sdk-example-usage)
- [Available Resources and Operations](#available-resources-and-operations)
- [Error Handling](#error-handling)
- [Server Selection](#server-selection)
- [Development](#development)
- [Maturity](#maturity)
- [Detailed Usage](#detailed-usage)

## SDK Installation

```bash
go get github.com/FastPix/fastpix-go@v0.1.0
```

## Initialization

You can set the security parameters using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
)

func main() {
	ctx := context.Background()
	sdk := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.String("your-username"),
			Password: fastpixgo.String("your-password"),
		}),
	)
	// Use sdk...
}
```

## SDK Example Usage

```go
package main

import (
	"context"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := fastpixgo.New(
		fastpixgo.WithSecurity(components.Security{
			Username: fastpixgo.String(""),
			Password: fastpixgo.String(""),
		}),
	)

	res, err := s.StartLiveStream.CreateNewStream(ctx, &components.CreateLiveStreamRequest{
		PlaybackSettings: components.PlaybackSettings{},
		InputMediaSettings: components.InputMediaSettings{
			Metadata: &components.CreateLiveStreamRequestMetadata{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.LiveStreamResponseDTO != nil {
		// handle response
	}
}
```

## Available Resources and Operations

### InputVideo

- [CreateMedia](docs/sdks/inputvideo/README.md#createmedia) - Create media from URL
- [DirectUploadVideoMedia](docs/sdks/inputvideo/README.md#directuploadvideomedia) - Upload media from device

### ManageLiveStream

- [GetAllStreams](docs/sdks/managelivestream/README.md#getallstreams) - Get all live streams
- [GetLiveStreamByID](docs/sdks/managelivestream/README.md#getlivestreambyid) - Get stream by ID
- [DeleteLiveStream](docs/sdks/managelivestream/README.md#deletelivestream) - Delete a stream
- [UpdateLiveStream](docs/sdks/managelivestream/README.md#updatelivestream) - Update a stream

### ManageVideos

- [ListMedia](docs/sdks/managevideos/README.md#listmedia) - Get list of all media
- [GetMedia](docs/sdks/managevideos/README.md#getmedia) - Get a media by ID
- [UpdatedMedia](docs/sdks/managevideos/README.md#updatedmedia) - Update a media by ID
- [DeleteMedia](docs/sdks/managevideos/README.md#deletemedia) - Delete a media by ID
- [RetrieveMediaInputInfo](docs/sdks/managevideos/README.md#retrievemediainputinfo) - Get info of media inputs

### Playback

- [CreatePlaybackIDOfStream](docs/sdks/playback/README.md#createplaybackidofstream) - Create a playbackId
- [DeletePlaybackIDOfStream](docs/sdks/playback/README.md#deleteplaybackidofstream) - Delete a playbackId
- [GetLiveStreamPlaybackID](docs/sdks/playback/README.md#getlivestreamplaybackid) - Get stream's playbackId
- [CreateMediaPlaybackID](docs/sdks/playback/README.md#createmediaplaybackid) - Create a playback ID
- [DeleteMediaPlaybackID](docs/sdks/playback/README.md#deletemediaplaybackid) - Delete a playback ID

### SimulcastStream

- [CreateSimulcastOfStream](docs/sdks/simulcaststream/README.md#createsimulcastofstream) - Create a simulcast
- [DeleteSimulcastOfStream](docs/sdks/simulcaststream/README.md#deletesimulcastofstream) - Delete a simulcast
- [GetSpecificSimulcastOfStream](docs/sdks/simulcaststream/README.md#getspecificsimulcastofstream) - Get a specific simulcast of a stream
- [UpdateSpecificSimulcastOfStream](docs/sdks/simulcaststream/README.md#updatespecificsimulcastofstream) - Update a specific simulcast of a stream

### StartLiveStream

- [CreateNewStream](docs/sdks/startlivestream/README.md#createnewstream) - Create a new stream

## Error Handling

All operations return a response object or an error. By default, an API error will return an `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to the respective *Errors* tables in the SDK docs for more details on possible error types for each operation.

For example, the `CreateNewStream` function may return the following errors:

| Error Type                        | Status Code | Content Type     |
| --------------------------------- | ----------- | ---------------- |
| apierrors.UnauthorizedError       | 401         | application/json |
| apierrors.InvalidPermissionError  | 403         | application/json |
| apierrors.ValidationErrorResponse | 422         | application/json |
| apierrors.APIError                | 4XX, 5XX    | */*              |

## Server Selection

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:

```go
sdk := fastpixgo.New(
	fastpixgo.WithServerURL("https://api.fastpix.io/v1/on-demand"),
	fastpixgo.WithSecurity(components.Security{
		Username: fastpixgo.String("your-username"),
		Password: fastpixgo.String("your-password"),
	}),
)
```

## Development

### Maturity

This SDK is currently in beta, and breaking changes may occur between versions even without a major version update. To avoid unexpected issues, we recommend pinning your dependency to a specific version. This ensures consistent behavior unless you intentionally update to a newer release.

## Detailed Usage

For a complete understanding of each API's functionality, including request and response details, parameter descriptions, and additional examples, please refer to the [FastPix API Reference](docs/).
