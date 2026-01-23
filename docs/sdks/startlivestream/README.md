# StartLiveStream

## Overview

### Available Operations

* [Create](#create) - Create a new stream

## Create

Creates a new <a href="https://docs.fastpix.io/docs/get-started-with-live-streaming">RTMPS</a> or <a href="https://docs.fastpix.io/docs/using-srt-to-live-stream">SRT</a> live stream in FastPix. When you create a stream, FastPix generates a unique `streamKey` and `srtSecret` that you can use with broadcasting software such as OBS to connect to FastPix RTMPS or SRT servers. Use SRT for live streaming in unstable network conditions, as it provides error correction and encryption for a more reliable and secure broadcast.

Leverage SRT for live streaming in environments with unstable networks, taking advantage of its error correction and encryption features for a resilient and secure broadcast. 

<h4>How it works</h4> 

1. Send a `POST` request to this endpoint. You can configure the stream settings, including `metadata` (such as stream name and description), `reconnectWindow` (in case of disconnection), and privacy options (`public` or `private`). 

2. FastPix returns the stream details for both RTMPS and SRT configurations. These keys and IDs from the stream details are essential for connecting the broadcasting software to FastPix’s servers and transmitting the live stream to viewers.

3. After the live stream is created, FastPix sends a `POST` request to your specified webhook endpoint with the event <a href="https://docs.fastpix.io/docs/live-events#videolive_streamcreated">video.live_stream.created</a>.


**Example:**


  Imagine a gaming platform that allows users to live stream gameplay directly from their dashboard. The API creates a new stream, provides the necessary stream key, and sets it to "private" so that only specific viewers can access it. 


Related guide: <a href="https://docs.fastpix.io/docs/how-to-livestream">How to live stream</a>

### Example Usage

<!-- UsageSnippet language="go" operationID="create-new-stream" method="post" path="/live/streams" -->
```go
package main

import(
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
            Username: fastpixgo.Pointer("your access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.StartLiveStream.Create(ctx, components.CreateLiveStreamRequest{
        PlaybackSettings: components.PlaybackSettings{},
        InputMediaSettings: components.InputMediaSettings{
            Metadata: map[string]string{
                "livestream_name": "your-livestream-name",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LiveStreamResponseDTO != nil {
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
                responseJSON, err := json.MarshalIndent(res.LiveStreamResponseDTO, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.LiveStreamResponseDTO)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.LiveStreamResponseDTO, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.LiveStreamResponseDTO)
            } else {
                fmt.Println(string(responseJSON))
            }
        }
    } else if res.DefaultError != nil {
        fmt.Printf("Error: %+v\n", res.DefaultError)
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateLiveStreamRequest](../../models/components/createlivestreamrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateNewStreamResponse](../../models/operations/createnewstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |