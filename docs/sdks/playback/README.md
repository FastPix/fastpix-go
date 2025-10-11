# Playback
(*Playback*)

## Overview

### Available Operations

* [CreateMediaPlaybackID](#createmediaplaybackid) - Create a playback ID
* [DeleteMediaPlaybackID](#deletemediaplaybackid) - Delete a playback ID
* [GetPlaybackID](#getplaybackid) - Get a playback ID

## CreateMediaPlaybackID

You can create a new playback ID for a specific media asset. If you have already retrieved an existing `playbackId` using the <a href="https://docs.fastpix.io/reference/get-media">Get Media by ID</a> endpoint for a media asset, you can use this endpoint to generate a new playback ID with a specified access policy. 



If you want to create a private playback ID for a media asset that already has a public playback ID, this endpoint also allows you to do so by specifying the desired access policy. 

#### How it works

1. Make a `POST` request to this endpoint, replacing `<mediaId>` with the `uploadId` or `id` of the media asset. 

2. Include the `accessPolicy` in the request body with `private` or `public` as the value. 

3. Receive a response containing the newly created playback ID with the requested access level. 


#### Example
A video streaming service generates playback IDs for each media file when users request to view specific content. The playback ID is then used by the video player to stream the video.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-media-playback-id" method="post" path="/on-demand/{mediaId}/playback-ids" -->
```go
package main

import(
	"context"
	"github.com/FastPix/fastpix-go/models/components"
	fastpixgo "github.com/FastPix/fastpix-go"
	"github.com/FastPix/fastpix-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.Playback.CreateMediaPlaybackID(ctx, "dbb8a39a-e4a5-4120-9f22-22f603f1446e", &operations.CreateMediaPlaybackIDRequestBody{
        AccessPolicy: components.AccessPolicyPublic,
        DrmConfigurationID: fastpixgo.Pointer("123e4567-e89b-12d3-a456-426614174000"),
        Resolution: operations.ResolutionOneThousandAndEightyp.ToPointer(),
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
| `requestBody`                                                                                                     | [*operations.CreateMediaPlaybackIDRequestBody](../../models/operations/createmediaplaybackidrequestbody.md)       | :heavy_minus_sign:                                                                                                | Request body for creating playback id for an media                                                                |                                                                                                                   |
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

This endpoint allows you to remove a specific playback ID associated with a media asset. Deleting a `playbackId` will revoke access to the media content linked to that ID. 


#### How it works

1. Make a `DELETE` request to this endpoint, replacing `<mediaId>` with the unique ID of the media asset from which you want to delete the playback ID. 

2. Specify the `playbackId` you wish to delete in the request body. 

#### Example

Your platform offers limited-time access to premium content. When the subscription expires, you can revoke access to the content by deleting the associated playback ID, preventing users from streaming the video further.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-media-playback-id" method="delete" path="/on-demand/{mediaId}/playback-ids" -->
```go
package main

import(
	"context"
	"github.com/FastPix/fastpix-go/models/components"
	fastpixgo "github.com/FastPix/fastpix-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
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

## GetPlaybackID

This endpoint retrieves details about a specific playback ID associated with a media asset. This endpoint is commonly used to check the access policy (e.g., public or private) with the specific playback ID.

**How it works:**
1. Make a GET request to the endpoint, replacing `{mediaId}` with the `id` of the media, and `{playbackId}` with the specific playback ID.
2. Useful for auditing or validation before granting playback access in your application.

**Example:**
A media platform might use this endpoint to verify if a playback ID is public or private before embedding the video in a frontend player or allowing access to a restricted group.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-playback-id" method="get" path="/on-demand/{mediaId}/playback-ids/{playbackId}" -->
```go
package main

import(
	"context"
	"github.com/FastPix/fastpix-go/models/components"
	fastpixgo "github.com/FastPix/fastpix-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := fastpixgo.New(
        fastpixgo.WithSecurity(components.Security{
            Username: fastpixgo.Pointer("your-access-token"),
            Password: fastpixgo.Pointer("your-secret-key"),
        }),
    )

    res, err := s.Playback.GetPlaybackID(ctx, "4fa85f64-5717-4562-b3fc-2c963f66afa6", "4fa85f64-5717-4562-b3fc-2c963f66afa6")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mediaID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 4fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `playbackID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 4fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPlaybackIDResponse](../../models/operations/getplaybackidresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.InvalidPermissionError       | 401                                    | application/json                       |
| apierrors.ForbiddenError               | 403                                    | application/json                       |
| apierrors.MediaOrPlaybackNotFoundError | 404                                    | application/json                       |
| apierrors.ValidationErrorResponse      | 422                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |