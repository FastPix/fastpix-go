# Playlist
(*Playlist*)

## Overview

### Available Operations

* [CreateAPlaylist](#createaplaylist) - Create a new playlist
* [GetAllPlaylists](#getallplaylists) - Get all playlists
* [GetPlaylistByID](#getplaylistbyid) - Get a playlist by ID
* [UpdateAPlaylist](#updateaplaylist) - Update a playlist by ID
* [DeleteAPlaylist](#deleteaplaylist) - Delete a playlist by ID
* [AddMediaToPlaylist](#addmediatoplaylist) - Add media to a playlist by ID
* [ChangeMediaOrderInPlaylist](#changemediaorderinplaylist) - Change media order in a playlist by ID
* [DeleteMediaFromPlaylist](#deletemediafromplaylist) - Delete media in a playlist by ID

## CreateAPlaylist

This endpoint creates a new playlist within a specified workspace. A playlist acts as a container for organizing media items either manually or based on filters and metadata. <br> <br>
### Playlists can be created in two modes
- **Manual:** An empty playlist is created without any initial media items. It's intended for manual curation, where items can be added later in a user-defined sequence.
- **Smart:** The playlist is auto-populated at creation time based on filters (video creation date range) criteria provided in the request.

#### How it works 

 - When a user sends a POST request to this endpoint, FastPix creates a playlist and returns a playlist ID, using which items can be added later in a user-defined sequence.
 - For a smart playlist, the playlist will be auto-populated based on metadata in the request body.


#### Example
An e-learning platform creates a new playlist titled "Beginner Python Series" via the API. The response includes a unique playlist ID. The platform then uses this ID to add a series of video tutorials to the playlist in a defined order. The playlist is presented to learners on the frontend as a structured learning path.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-a-playlist" method="post" path="/on-demand/playlists" -->
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

    res, err := s.Playlist.CreateAPlaylist(ctx, components.CreatePlaylistRequest{
        Name: "playlist name",
        ReferenceID: "a1",
        Type: components.CreatePlaylistRequestTypeSmart,
        Description: fastpixgo.Pointer("This is a playlist"),
        PlayOrder: components.PlaylistOrderCreatedDateAsc.ToPointer(),
        Limit: fastpixgo.Pointer[int64](20),
        Metadata: &components.CreatePlaylistRequestMetadata{
            CreatedDate: &components.DateRange{
                StartDate: fastpixgo.Pointer("2024-11-11"),
                EndDate: fastpixgo.Pointer("2024-12-12"),
            },
            UpdatedDate: &components.DateRange{
                StartDate: fastpixgo.Pointer("2024-11-11"),
                EndDate: fastpixgo.Pointer("2024-12-12"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistCreatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreatePlaylistRequest](../../models/components/createplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateAPlaylistResponse](../../models/operations/createaplaylistresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.UnauthorizedError                 | 401                                         | application/json                            |
| apierrors.InvalidPermissionError            | 403                                         | application/json                            |
| apierrors.DuplicateReferenceIDErrorResponse | 409                                         | application/json                            |
| apierrors.ValidationErrorResponse           | 422                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## GetAllPlaylists

This endpoint retrieves all playlists present within a specified workspace. It allows users to view the collection of playlists that have been created, whether manual or smart, along with their associated metadata.
#### How it works

 - When a user sends a GET request to this endpoint, FastPix returns a list of all playlists in the workspace, including details such as playlist IDs, titles, creation mode (manual or smart), and other relevant metadata.
 
#### Example

  An e-learning platform requests all playlists within a workspace to display an overview of available learning paths. The response includes multiple playlists like "Beginner Python Series" and "Advanced Java Tutorials," enabling the platform to show users a catalog of curated content collections.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-all-playlists" method="get" path="/on-demand/playlists" -->
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

    res, err := s.Playlist.GetAllPlaylists(ctx, fastpixgo.Pointer[int64](1), fastpixgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAllPlaylistsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `limit`                                                                                 | **int64*                                                                                | :heavy_minus_sign:                                                                      | The number of playlists to return (default is 10, max is 50).                           | 1                                                                                       |
| `offset`                                                                                | **int64*                                                                                | :heavy_minus_sign:                                                                      | The page number to retrieve, starting from 1. Used for paginating the playlist results. | 1                                                                                       |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.GetAllPlaylistsResponse](../../models/operations/getallplaylistsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.UnauthorizedError | 401                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |

## GetPlaylistByID

This endpoint retrieves detailed information about a specific playlist using its unique `playlistId`. It provides comprehensive metadata about the playlist, including its title, creation mode (manual or smart), media items along with the metadata of each media in the playlist.

 
#### Example
An e-learning platform requests details for the playlist "Beginner Python Series" by providing its unique `playlistId`. The response includes the playlist's title, creation mode, and the ordered list of video tutorials contained within, enabling the platform to present the full learning path to users.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-playlist-by-id" method="get" path="/on-demand/playlists/{playlistId}" -->
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

    res, err := s.Playlist.GetPlaylistByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
        // handle response
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

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.NotFoundError                  | 404                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## UpdateAPlaylist

This endpoint allows you to update the name and description of an existing playlist. It enables modifications to the playlist's metadata without altering the media items or playlist structure.
#### How it works

 - When a user sends a PUT request to this endpoint with the `playlistId` and updated name and description in the request body, FastPix updates the playlist metadata accordingly and returns the updated playlist details.

#### Example
An e-learning platform updates the playlist titled "Beginner Python Series" to rename it as "Python Basics" and add a more detailed description. The updated metadata is reflected when retrieving the playlist, helping users better understand the playlist content.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-a-playlist" method="put" path="/on-demand/playlists/{playlistId}" -->
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

    res, err := s.Playlist.UpdateAPlaylist(ctx, "<id>", components.UpdatePlaylistRequest{
        Name: "updated name",
        Description: "updated description",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistCreatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `playlistID`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | The unique id of the playlist you want to retrieve.                                  |                                                                                      |
| `updatePlaylistRequest`                                                              | [components.UpdatePlaylistRequest](../../models/components/updateplaylistrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  | {<br/>"name": "updated name",<br/>"description": "updated description"<br/>}         |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpdateAPlaylistResponse](../../models/operations/updateaplaylistresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.InvalidPermissionError         | 403                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## DeleteAPlaylist

This endpoint allows you to delete an existing playlist from the workspace. Once deleted, the playlist and its metadata are permanently removed and cannot be recovered.
#### How it works
 - When a user sends a DELETE request to this endpoint with the `playlistId`, FastPix removes the specified playlist from the workspace and returns a confirmation of successful deletion.
 
#### Example
An e-learning platform deletes an outdated playlist titled "Old Python Tutorials" by providing its unique playlist ID. The platform receives confirmation that the playlist has been removed, ensuring learners no longer see the obsolete content.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-a-playlist" method="delete" path="/on-demand/playlists/{playlistId}" -->
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

    res, err := s.Playlist.DeleteAPlaylist(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `playlistID`                                             | *string*                                                 | :heavy_check_mark:                                       | The unique id of the playlist you want to delete.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAPlaylistResponse](../../models/operations/deleteaplaylistresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.InvalidPermissionError         | 403                                      | application/json                         |
| apierrors.NotFoundError                  | 404                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## AddMediaToPlaylist

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

    res, err := s.Playlist.AddMediaToPlaylist(ctx, "<id>", components.MediaIdsRequest{
        MediaIds: []string{
            "a1cd180e-f9b5-4e99-9d44-b9c9baabad89",
            "245800c3-7b73-47d9-a201-e961260dcb30",
            "41316aac-5396-4278-8f44-08d5f2495b12",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `playlistID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | The unique id of the playlist you want to perform the operation on.      |
| `mediaIdsRequest`                                                        | [components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.AddMediaToPlaylistResponse](../../models/operations/addmediatoplaylistresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.InvalidPermissionError         | 403                                      | application/json                         |
| apierrors.NotFoundError                  | 404                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## ChangeMediaOrderInPlaylist

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

    res, err := s.Playlist.ChangeMediaOrderInPlaylist(ctx, "<id>", components.MediaIdsRequest{
        MediaIds: []string{
            "a1cd180e-f9b5-4e99-9d44-b9c9baabad89",
            "245800c3-7b73-47d9-a201-e961260dcb30",
            "41316aac-5396-4278-8f44-08d5f2495b12",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `playlistID`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | The unique id of the playlist you want to perform the operation on.      |
| `mediaIdsRequest`                                                        | [components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.ChangeMediaOrderInPlaylistResponse](../../models/operations/changemediaorderinplaylistresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.InvalidPermissionError         | 403                                      | application/json                         |
| apierrors.NotFoundError                  | 404                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## DeleteMediaFromPlaylist

This endpoint allows you to delete one or more media items from an existing playlist. By passing the media ID(s) in the request, the specified media items are removed from the playlist.
#### How it works

 - When a user sends a DELETE request to this endpoint with the playlist ID as the path parameter and the media ID(s) to be removed in the request body, FastPix deletes the specified media items from the playlist and returns the updated playlist details.
 
#### Example
An e-learning platform removes outdated video tutorials from the "Beginner Python Series" playlist by specifying their media IDs in the request. The playlist is updated to exclude these items, ensuring learners only access relevant content.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-media-from-playlist" method="delete" path="/on-demand/playlists/{playlistId}/media" -->
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

    res, err := s.Playlist.DeleteMediaFromPlaylist(ctx, "<id>", &components.MediaIdsRequest{
        MediaIds: []string{
            "a1cd180e-f9b5-4e99-9d44-b9c9baabad89",
            "245800c3-7b73-47d9-a201-e961260dcb30",
            "41316aac-5396-4278-8f44-08d5f2495b12",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `playlistID`                                                              | *string*                                                                  | :heavy_check_mark:                                                        | The unique id of the playlist you want to perform the operation on.       |
| `mediaIdsRequest`                                                         | [*components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.DeleteMediaFromPlaylistResponse](../../models/operations/deletemediafromplaylistresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.UnauthorizedError              | 401                                      | application/json                         |
| apierrors.InvalidPermissionError         | 403                                      | application/json                         |
| apierrors.NotFoundError                  | 404                                      | application/json                         |
| apierrors.InvalidPlaylistIDResponseError | 422                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |