# Playback
(*Playback*)

## Overview

### Available Operations

* [CreatePlaybackIDOfStream](#createplaybackidofstream) - Create a playbackId
* [DeletePlaybackIDOfStream](#deleteplaybackidofstream) - Delete a playbackId
* [GetLiveStreamPlaybackID](#getlivestreamplaybackid) - Get stream's playbackId
* [CreateMediaPlaybackID](#createmediaplaybackid) - Create a playback ID
* [DeleteMediaPlaybackID](#deletemediaplaybackid) - Delete a playback ID

## CreatePlaybackIDOfStream

Generates a new playback ID for the live stream, allowing viewers to access the stream through this ID. The playback ID can be shared with viewers for direct access to the live broadcast. 

  By calling this endpoint with the streamId, FastPix returns a unique playbackId, which can be used to stream the live content. 

  **Use case:** A media platform needs to distribute a unique playback ID to users for an exclusive live concert. The platform can also embed the stream on various partner websites. 

### Example Usage

```go
package main

import(
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

    res, err := s.Playback.CreatePlaybackIDOfStream(ctx, "8717422d89288ad5958d4a86e9afe2a2", &components.PlaybackIDRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaybackIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `streamID`                                                                          | *string*                                                                            | :heavy_check_mark:                                                                  | Upon creating a new live stream, FastPix assigns a unique identifier to the stream. | 8717422d89288ad5958d4a86e9afe2a2                                                    |
| `playbackIDRequest`                                                                 | [*components.PlaybackIDRequest](../../models/components/playbackidrequest.md)       | :heavy_minus_sign:                                                                  | N/A                                                                                 | {<br/>"accessPolicy": "public"<br/>}                                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.CreatePlaybackIDOfStreamResponse](../../models/operations/createplaybackidofstreamresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.UnauthorizedError       | 401                               | application/json                  |
| apierrors.InvalidPermissionError  | 403                               | application/json                  |
| apierrors.NotFoundError           | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## DeletePlaybackIDOfStream

Deletes a previously created playback ID for a live stream. This will prevent any new viewers from accessing the stream through the playback ID, though current viewers will be able to continue watching for a limited time before being disconnected. By providing the playbackId, FastPix deletes the ID and ensures new playback requests will fail. 

  **Use case:** A streaming service wants to prevent new users from joining a live stream that is nearing its end. The host can delete the playback ID to ensure no one can join the stream or replay it once it ends. 

### Example Usage

```go
package main

import(
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

    res, err := s.Playback.DeletePlaybackIDOfStream(ctx, "8717422d89288ad5958d4a86e9afe2a2", "88b7ac0f-2504-4dd5-b7b4-d84ab4fee1bd")
    if err != nil {
        log.Fatal(err)
    }
    if res.LiveStreamDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |                                                                                     |
| `streamID`                                                                          | *string*                                                                            | :heavy_check_mark:                                                                  | Upon creating a new live stream, FastPix assigns a unique identifier to the stream. | 8717422d89288ad5958d4a86e9afe2a2                                                    |
| `playbackID`                                                                        | *string*                                                                            | :heavy_check_mark:                                                                  | Unique identifier for the playbackId                                                | 88b7ac0f-2504-4dd5-b7b4-d84ab4fee1bd                                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |                                                                                     |

### Response

**[*operations.DeletePlaybackIDOfStreamResponse](../../models/operations/deleteplaybackidofstreamresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.UnauthorizedError       | 401                               | application/json                  |
| apierrors.InvalidPermissionError  | 403                               | application/json                  |
| apierrors.NotFoundErrorPlaybackID | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## GetLiveStreamPlaybackID

Retrieves details about a previously created playback ID. If you provide the distinct playback ID that was given back to you in the previous stream or playbackId create request, FastPix will provide the relevant playback details such as the access policy. 

  **Use case:** A developer needs to confirm the playback ID details to ensure the right stream is being accessed by viewers. 

### Example Usage

```go
package main

import(
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

    res, err := s.Playback.GetLiveStreamPlaybackID(ctx, "61a264dcc447b63da6fb79ef925cd76d", "61a264dcc447b63da6fb79ef925cd76d")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaybackIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `streamID`                                                                           | *string*                                                                             | :heavy_check_mark:                                                                   | Upon creating a new live stream, FastPix assigns a unique identifier to the stream.  | 61a264dcc447b63da6fb79ef925cd76d                                                     |
| `playbackID`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | Upon creating a new playbackId, FastPix assigns a unique identifier to the playback. | 61a264dcc447b63da6fb79ef925cd76d                                                     |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.GetLiveStreamPlaybackIDResponse](../../models/operations/getlivestreamplaybackidresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.UnauthorizedError       | 401                               | application/json                  |
| apierrors.InvalidPermissionError  | 403                               | application/json                  |
| apierrors.NotFoundErrorPlaybackID | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## CreateMediaPlaybackID

You can create a new playback ID for a specific media asset. If you have already retrieved an existing playbackId using the "Get Media by ID" endpoint for a media asset, you can use this endpoint to generate a new playback ID with a specified access policy. 



If you want to create a private playback ID for a media asset that already has a public playback ID, this endpoint also allows you to do so by specifying the desired access policy. 

#### How it works

1. **Make a POST request** to the **/on-demand/`<mediaId>`/playback-ids** endpoint, replacing `<mediaId>` with the uploadId or id of the media asset. 

2. Include the **access policy** in the request body to indicate whether the new playback ID should be private or public. 

3. Receive a response containing the newly created playback ID with the requested access level. 


**Use case:** A video streaming service generates playback IDs for each media file when users request to view specific content. The playback ID is then used by the video player to stream the video.


### Example Usage

```go
package main

import(
	"context"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
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

    res, err := s.Playback.CreateMediaPlaybackID(ctx, "dbb8a39a-e4a5-4120-9f22-22f603f1446e", &operations.CreateMediaPlaybackIDRequestBody{
        AccessPolicy: operations.CreateMediaPlaybackIDAccessPolicyPublic,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |                                                                                                                   |
| `mediaID`                                                                                                         | *string*                                                                                                          | :heavy_check_mark:                                                                                                | When creating the media, FastPix assigns a universally unique identifier with a maximum length of 255 characters. | dbb8a39a-e4a5-4120-9f22-22f603f1446e                                                                              |
| `requestBody`                                                                                                     | [*operations.CreateMediaPlaybackIDRequestBody](../../models/operations/createmediaplaybackidrequestbody.md)       | :heavy_minus_sign:                                                                                                | Request body for creating playback id for an media                                                                | {<br/>"accessPolicy": "public"<br/>}                                                                              |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |                                                                                                                   |

### Response

**[*operations.CreateMediaPlaybackIDResponse](../../models/operations/createmediaplaybackidresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ForbiddenError          | 403                               | application/json                  |
| apierrors.MediaNotFoundError      | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## DeleteMediaPlaybackID

This endpoint allows you to remove a specific playback ID associated with a media asset. Deleting a playbackId will revoke access to the media content linked to that ID. 


#### How it works

1. Make a DELETE request to the **/on-demand/`<mediaId>`/playback-ids** endpoint, replacing `<mediaId>` with the uploadId or id of the media asset from which you want to delete the playback ID. 

2. Specify the playback ID you wish to delete in the request body. 

**Use case:** Your platform offers limited-time access to premium content. When the subscription expires, you can revoke access to the content by deleting the associated playback ID, preventing users from streaming the video further. 


### Example Usage

```go
package main

import(
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

    res, err := s.Playback.DeleteMediaPlaybackID(ctx, "dbb8a39a-e4a5-4120-9f22-22f603f1446e", "dbb8a39a-e4a5-4120-9f22-22f603f1446e")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `mediaID`                                                                                             | *string*                                                                                              | :heavy_check_mark:                                                                                    | Return the universal unique identifier for media which can contain a maximum of 255 characters.       | dbb8a39a-e4a5-4120-9f22-22f603f1446e                                                                  |
| `playbackID`                                                                                          | *string*                                                                                              | :heavy_check_mark:                                                                                    | Return the universal unique identifier for playbacks  which can contain a maximum of 255 characters.  | dbb8a39a-e4a5-4120-9f22-22f603f1446e                                                                  |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.DeleteMediaPlaybackIDResponse](../../models/operations/deletemediaplaybackidresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.InvalidPermissionError       | 401                                    | application/json                       |
| apierrors.ForbiddenError               | 403                                    | application/json                       |
| apierrors.MediaOrPlaybackNotFoundError | 404                                    | application/json                       |
| apierrors.ValidationErrorResponse      | 422                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |