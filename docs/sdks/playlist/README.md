# Playlist

## Overview

Operations for playlist management

### Available Operations

* [Get](#get) - Get a playlist by ID
* [AddMedia](#addmedia) - Add media to a playlist by ID
* [ChangeMediaOrder](#changemediaorder) - Change media order in a playlist by ID

## Get

This endpoint retrieves detailed information about a specific playlist using its unique `playlistId`. It provides comprehensive metadata about the playlist, including its title, creation mode (manual or smart), media items along with the metadata of each media in the playlist.


#### Example
An e-learning platform requests details for the playlist "Beginner Python Series" by providing its unique `playlistId`. The response includes the playlist"s title, creation mode, and the ordered list of video tutorials contained within, enabling the platform to present the full learning path to users.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-playlist-by-id" method="get" path="/on-demand/playlists/{playlistId}" -->
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

    // your-playlist-id: Playlist ID returned from create playlist API
    res, err := s.Playlist.Get(ctx, "your-playlist-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `playlistID`                                             | *string*                                                 | :heavy_check_mark:                                       | The unique id of the playlist you want to retrieve.      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPlaylistByIDResponse](../../models/operations/getplaylistbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMedia

This endpoint allows you to add one or more media items to an existing playlist. By passing the media ID(s) in the request, the specified media items are appended to the playlist in the order provided.
#### How it works

 - When a user sends a PATCH request to this endpoint with the `playlistId` as path parameter and a list of media ID(s) in the request body, FastPix adds the specified media items to the playlist and returns the updated playlist details.

#### Example
An e-learning platform adds new video tutorials to the "Beginner Python Series" playlist by sending their media IDs in the request. The playlist is updated with the new content, ensuring learners have access to the latest tutorials in sequence.

### Example Usage

<!-- UsageSnippet language="go" operationID="add-media-to-playlist" method="patch" path="/on-demand/playlists/{playlistId}/media" -->
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

    // your-playlist-id: Playlist ID returned from create playlist API
    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Playlist.AddMedia(ctx, "your-playlist-id", components.MediaIdsRequest{
        MediaIds: []string{
            "your-media-id",
            "your-media-id-2",
            "your-media-id-3",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `playlistID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | The unique id of the playlist you want to perform the operation on.      |
| `body`                                                                   | [components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.AddMediaToPlaylistResponse](../../models/operations/addmediatoplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ChangeMediaOrder

This endpoint allows you to change the order of media items within a playlist. By passing the complete list of media IDs in the desired sequence, the playlist's play order is updated accordingly.
#### How it works

 - When a user sends a PUT request to this endpoint with the `playlistId` as path parameter and the reordered list of all media IDs in the request body, FastPix updates the playlist to reflect the new media sequence and returns the updated playlist details.

#### Example
An e-learning platform rearranges the "Beginner Python Series" playlist by submitting a reordered list of media IDs. The playlist now follows the new sequence, providing learners with a better structured learning path.

### Example Usage

<!-- UsageSnippet language="go" operationID="change-media-order-in-playlist" method="put" path="/on-demand/playlists/{playlistId}/media" -->
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

    // your-playlist-id: Playlist ID returned from create playlist API
    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Playlist.ChangeMediaOrder(ctx, "your-playlist-id", components.MediaIdsRequest{
        MediaIds: []string{
            "your-media-id",
            "your-media-id-2",
            "your-media-id-3",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaylistByIDResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaylistByIDResponse)
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `playlistID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | The unique id of the playlist you want to perform the operation on.      |
| `body`                                                                   | [components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.ChangeMediaOrderInPlaylistResponse](../../models/operations/changemediaorderinplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |