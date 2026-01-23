# LivePlayback

## Overview

### Available Operations

* [Create](#create) - Create a playbackId
* [GetPlaybackIDDetails](#getplaybackiddetails) - Get playbackId details
* [DeletePlaybackID](#deleteplaybackid) - Delete a playbackId

## Create

Generates a new playback ID for the live stream, allowing viewers to access the stream through this ID. The playback ID can be shared with viewers for direct access to the live broadcast. 

  By calling this endpoint with the `streamId`, FastPix returns a unique `playbackId`, which can be used to stream the live content. 

  #### Example

  A media platform needs to distribute a unique playback ID to users for an exclusive live concert. The platform can also embed the stream on various partner websites.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-playbackId-of-stream" method="post" path="/live/streams/{streamId}/playback-ids" -->
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

    // your-stream-id: Live stream ID returned from create stream API
    res, err := s.LivePlayback.Create(ctx, "your-stream-id", components.PlaybackIDRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaybackIDSuccessResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaybackIDSuccessResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaybackIDSuccessResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaybackIDSuccessResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaybackIDSuccessResponse)
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `streamID`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | After creating a new live stream, FastPix assigns a unique identifier to the stream. | your-stream-id                                                     |
| `body`                                                                               | [components.PlaybackIDRequest](../../models/components/playbackidrequest.md)         | :heavy_check_mark:                                                                   | N/A                                                                                  | {<br/>"accessPolicy": "public"<br/>}                                                 |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.CreatePlaybackIDOfStreamResponse](../../models/operations/createplaybackidofstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPlaybackIDDetails

Retrieves details for an existing playback ID. When you provide the playbackId returned from a previous stream or playback creation request, FastPix returns the associated playback information, including the access policy.

#### Example
A developer needs to confirm the access policy of the playback ID to ensure whether the stream is public or private for viewers.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-live-stream-playback-id" method="get" path="/live/streams/{streamId}/playback-ids/{playbackId}" -->
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

    // your-stream-id: Live stream ID returned from create stream API
    // your-playback-id: Playback ID returned from create playback ID API
    res, err := s.LivePlayback.GetPlaybackIDDetails(ctx, "your-stream-id", "your-playback-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaybackIDSuccessResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaybackIDSuccessResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaybackIDSuccessResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaybackIDSuccessResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaybackIDSuccessResponse)
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

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `streamID`                                                                            | *string*                                                                              | :heavy_check_mark:                                                                    | After creating a new live stream, FastPix assigns a unique identifier to the stream.  | your-stream-id                                                      |
| `playbackID`                                                                          | *string*                                                                              | :heavy_check_mark:                                                                    | After creating a new playbackId, FastPix assigns a unique identifier to the playback. | your-playback-id                                                      |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.GetLiveStreamPlaybackIDResponse](../../models/operations/getlivestreamplaybackidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeletePlaybackID

Deletes a previously created playback ID for a live stream.This prevents new viewers from accessing the stream using the playback ID, while current viewers can continue watching for a short period before the connection ends. FastPix deletes the ID and ensures the new playback request fails.

#### Example
A streaming service wants to prevent new users from joining a live stream that is nearing its end. The host can delete the playback ID to ensure no one can join the stream or replay it once it ends.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-playbackId-of-stream" method="delete" path="/live/streams/{streamId}/playback-ids" -->
```go
package main

import(
	"context"
	"encoding/json"
	"fmt"
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

    // your-stream-id: Live stream ID returned from create stream API
    // your-playback-id: Playback ID returned from create playback ID API
    res, err := s.LivePlayback.DeletePlaybackID(ctx, "your-stream-id", "your-playback-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.LiveStreamDeleteResponse != nil {
        // Print response in JSON format for better readability
        responseJSON, err := json.MarshalIndent(res.LiveStreamDeleteResponse, "", "  ")
        if err != nil {
            log.Printf("Error marshaling response: %v", err)
            fmt.Printf("Response: %+v\n", res.LiveStreamDeleteResponse)
        } else {
            fmt.Println(string(responseJSON))
        }
    } else if res.DefaultError != nil {
        fmt.Printf("Error: %+v\n", res.DefaultError)
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `streamID`                                                                          | *string*                                                                            | :heavy_check_mark:                                                                  | Upon creating a new live stream, FastPix assigns a unique identifier to the stream. | your-stream-id                                                    |
| `playbackID`                                                                        | *string*                                                                            | :heavy_check_mark:                                                                  | Unique identifier for the playbackId                                                | your-playback-id                                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.DeletePlaybackIDOfStreamResponse](../../models/operations/deleteplaybackidofstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |