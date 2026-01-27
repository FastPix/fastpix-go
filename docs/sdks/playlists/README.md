# Playlists

## Overview

### Available Operations

* [Create](#create) - Create a new playlist
* [List](#list) - Get all playlists
* [Update](#update) - Update a playlist by ID
* [Delete](#delete) - Delete a playlist by ID
* [DeleteMedia](#deletemedia) - Delete media in a playlist by ID

## Create

This endpoint creates a new playlist within a specified workspace. A playlist acts as a container for organizing media items either manually or based on filters and metadata. <br> <br>
### Playlists can be created in two modes
- **Manual:** Creates an empty playlist without any initial media items. Use this mode for manual curation, where you add items later in a user-defined sequence.
- **Smart:** Auto-populates the playlist at creation time based on the filter criteria (for example, a video creation date range) that you provide in the request.

For more details, see <a href="https://docs.fastpix.io/docs/create-and-manage-playlist">Create and manage playlist</a>.

#### How it works 

 - When you send a `POST` request to this endpoint, FastPix creates a playlist and returns a playlist ID, using which items can be added later in a user-defined sequence.
 - You can create a smart playlist that is auto-populated based on the metadata in the request body.


#### Example
An e-learning platform creates a new playlist titled Beginner Python Series through the API. The response returns a unique playlist ID. The platform uses this ID to add a series of video tutorials to the playlist in a defined order. The playlist appears on the frontend as a structured learning path for learners.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-a-playlist" method="post" path="/on-demand/playlists" -->
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

    res, err := s.Playlists.Create(ctx, components.CreateCreatePlaylistRequestSmart(
        components.CreatePlaylistRequestSmart{
            Name: "your-playlist-name",
            ReferenceID: "your-reference-id",
            Type: components.CreatePlaylistRequestSmartTypeSmart,
            Description: fastpixgo.Pointer("your-playlist-description"),
            PlayOrder: components.PlaylistOrderCreatedDateAsc,
            Limit: fastpixgo.Pointer[int64](20),
            Metadata: components.Metadata{
                CreatedDate: &components.DateRange{
                    StartDate: fastpixgo.Pointer("your-start-date"), // Start date for filtering (YYYY-MM-DD format, e.g., "2024-01-01")
                    EndDate: fastpixgo.Pointer("your-end-date"), // End date for filtering (YYYY-MM-DD format, e.g., "2024-12-31")
                },
                UpdatedDate: &components.DateRange{
                    StartDate: fastpixgo.Pointer("your-start-date"), // Start date for filtering (YYYY-MM-DD format, e.g., "2024-01-01")
                    EndDate: fastpixgo.Pointer("your-end-date"), // End date for filtering (YYYY-MM-DD format, e.g., "2024-12-31")
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistCreatedResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaylistCreatedResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaylistCreatedResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaylistCreatedResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaylistCreatedResponse)
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreatePlaylistRequest](../../models/components/createplaylistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateAPlaylistResponse](../../models/operations/createaplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

This endpoint retrieves all playlists in a specified workspace. It allows you to view the collection of manual and smart playlists along with their associated metadata.
#### How it works

 - When a user sends a GET request to this endpoint, FastPix returns a list of all playlists in the workspace, including details such as playlist IDs, titles, creation mode (manual or smart), and other relevant metadata.

#### Example

  An e-learning platform requests all playlists within a workspace to display an overview of available learning paths. The response includes multiple playlists like "Beginner Python Series" and "Advanced Java Tutorials," enabling the platform to show users a catalog of curated content collections.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-all-playlists" method="get" path="/on-demand/playlists" -->
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

    res, err := s.Playlists.List(ctx, fastpixgo.Pointer[int64](1), fastpixgo.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAllPlaylistsResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.GetAllPlaylistsResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.GetAllPlaylistsResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.GetAllPlaylistsResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.GetAllPlaylistsResponse)
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `limit`                                                                                            | **int64*                                                                                           | :heavy_minus_sign:                                                                                 | The number of playlists to return (default is 10, max is 50).                                      | 1                                                                                                  |
| `offset`                                                                                           | **int64*                                                                                           | :heavy_minus_sign:                                                                                 | The page number to retrieve, starting from 1. Use this parameter to paginate the playlist results. | 1                                                                                                  |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.GetAllPlaylistsResponse](../../models/operations/getallplaylistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

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
    res, err := s.Playlists.Update(ctx, "your-playlist-id", components.UpdatePlaylistRequest{
        Name: "your-updated-playlist-name",
        Description: "your-updated-playlist-description",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistCreatedResponse != nil {
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
                responseJSON, err := json.MarshalIndent(res.PlaylistCreatedResponse, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.PlaylistCreatedResponse)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.PlaylistCreatedResponse, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.PlaylistCreatedResponse)
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
| `playlistID`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | The unique id of the playlist you want to retrieve.                                  |                                                                                      |
| `body`                                                                               | [components.UpdatePlaylistRequest](../../models/components/updateplaylistrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  | {<br/>"name": "your-updated-playlist-name",<br/>"description": "your-updated-playlist-description"<br/>}         |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpdateAPlaylistResponse](../../models/operations/updateaplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

This endpoint allows you to delete an existing playlist from the workspace. After deleted, the playlist and its metadata are permanently removed and cannot be recovered.
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

    // your-playlist-id: Playlist ID returned from create playlist API
    res, err := s.Playlists.Delete(ctx, "your-playlist-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlaylistDeleteResponse != nil {
        // Print response in JSON format for better readability
        responseJSON, err := json.MarshalIndent(res.PlaylistDeleteResponse, "", "  ")
        if err != nil {
            log.Printf("Error marshaling response: %v", err)
            fmt.Printf("Response: %+v\n", res.PlaylistDeleteResponse)
        } else {
            fmt.Println(string(responseJSON))
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
| `playlistID`                                             | *string*                                                 | :heavy_check_mark:                                       | The unique id of the playlist you want to delete.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAPlaylistResponse](../../models/operations/deleteaplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteMedia

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
    res, err := s.Playlists.DeleteMedia(ctx, "your-playlist-id", &components.MediaIdsRequest{
        MediaIds: []string{
            "your-media-id",
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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `playlistID`                                                              | *string*                                                                  | :heavy_check_mark:                                                        | The unique id of the playlist you want to perform the operation on.       |
| `body`                                                                    | [*components.MediaIdsRequest](../../models/components/mediaidsrequest.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.DeleteMediaFromPlaylistResponse](../../models/operations/deletemediafromplaylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |