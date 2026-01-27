# Videos

## Overview

### Available Operations

* [ListLiveClips](#listliveclips) - Get all clips of a live stream
* [Get](#get) - Get a media by ID
* [Update](#update) - Update a media by ID
* [AddMediaTrack](#addmediatrack) - Add audio / subtitle track
* [DeleteTrack](#deletetrack) - Delete audio / subtitle track
* [UpdateSourceAccess](#updatesourceaccess) - Update the source access of a media by ID
* [UpdateMp4Support](#updatemp4support) - Update the mp4Support of a media by ID

## ListLiveClips

Retrieves a list of all media clips generated from a specific livestream. Each media entry includes metadata such as the clip media IDs, and other relevant details. A media clip is a segmented portion of an original media file (source live stream). Clips are often created for various purposes such as previews, highlights, or customized edits.
#### How it works
1. Provide the livestreamId as a parameter when calling this endpoint.

2. The API returns a paginated list of media clips created from the specified livestream.

3. Pagination helps maintain performance and usability when handling large sets of media files, making it easier to organize and manage content in bulk.

#### Use case
Suppose you’re hosting a live gaming event and want to showcase key moments from the stream — such as top plays or final match highlights. You can use this endpoint to fetch all clips generated from that livestream, display them in your dashboard, or use them for post-event editing and sharing.


Related guide: <a href="https://docs.fastpix.io/docs/instant-live-clipping">Instant live clipping</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list-live-clips" method="get" path="/on-demand/{livestreamId}/live-clips" -->
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

    // your-livestream-id: Live stream ID returned from create stream API
    res, err := s.Videos.ListLiveClips(ctx, "your-livestream-id", fastpixgo.Pointer[int64](20), fastpixgo.Pointer[int64](1), components.SortOrderDesc.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `livestreamID`                                                                            | *string*                                                                                  | :heavy_check_mark:                                                                        | The stream Id is unique identifier assigned to the live stream.                           | your-livestream-id                                                          |
| `limit`                                                                                   | **int64*                                                                                  | :heavy_minus_sign:                                                                        | Limit specifies the maximum number of items to display per page.                          | 20                                                                                        |
| `offset`                                                                                  | **int64*                                                                                  | :heavy_minus_sign:                                                                        | Offset determines the starting point for data retrieval within a paginated list.          | 1                                                                                         |
| `orderBy`                                                                                 | [*components.SortOrder](../../models/components/sortorder.md)                             | :heavy_minus_sign:                                                                        | The values in the list can be arranged in two ways: DESC (Descending) or ASC (Ascending). | desc                                                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.ListLiveClipsResponse](../../models/operations/listliveclipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

By calling this endpoint, you can retrieve detailed information about a specific media item, including its current `status` and a `playbackId`. This is particularly useful for retrieving specific media details when managing large content libraries. 



#### How it works

1. Send a GET request to this endpoint. Use the `<mediaId>` you received after uploading the media file.

2. The response includes details about the media:
   - **status** – Indicates whether the media is still *Processing* or has transitioned to *Ready*.
   - **playbackId** – A unique identifier that allows you to stream the media once it is *Ready*.  
     You can construct the stream URL as follows:  
     `https://stream.fastpix.io/<playbackId>.m3u8`

#### Example

If your platform provides users with a dashboard to manage uploaded content, a user might want to check whether a video has finished processing and is ready for playback. You can use the media ID to retrieve the information from FastPix and display it in the user’s dashboard.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-media" method="get" path="/on-demand/{mediaId}" -->
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

    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Videos.Get(ctx, "your-media-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `mediaID`                                                                                 | *string*                                                                                  | :heavy_check_mark:                                                                        | The unique identifier assigned to the media when created. The value must be a valid UUID. | your-media-id                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.GetMediaResponse](../../models/operations/getmediaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

This endpoint allows you to update specific parameters of an existing media file. You can modify the key-value pairs of the metadata that were provided in the payload during the creation of media from a URL or when uploading the media directly from device. 


#### How it works

1. Make a PATCH request to this endpoint. Replace `<mediaId>` with the unique ID (`uploadId` or `id`) of the media you received after uploading to FastPix

2. Include the updated parameters in the request body.

3. The response returns the updated media data, confirming the changes. 

4. Monitor the <a href="https://docs.fastpix.io/docs/media-events#videomediaupdated">video.media.updated</a> webhook event to track the update status in your system.

#### Example
If a user uploads a video and later needs to change the title, add a new description, or update tags, you can use this endpoint to update the media metadata without re-uploading the entire video.


### Example Usage

<!-- UsageSnippet language="go" operationID="updated-media" method="patch" path="/on-demand/{mediaId}" -->
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
	"github.com/FastPix/fastpix-go/models/operations"
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

    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Videos.Update(ctx, "your-media-id", operations.UpdatedMediaRequestBody{
        Metadata: map[string]string{
            "user": "your-metadata-value",
        },
        Title: fastpixgo.Pointer("your-video-title"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `mediaID`                                                                                 | *string*                                                                                  | :heavy_check_mark:                                                                        | The unique identifier assigned to the media when created. The value must be a valid UUID. | your-media-id                                                      |
| `body`                                                                                    | [operations.UpdatedMediaRequestBody](../../models/operations/updatedmediarequestbody.md)  | :heavy_check_mark:                                                                        | N/A                                                                                       |                                                                                           |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.UpdatedMediaResponse](../../models/operations/updatedmediaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMediaTrack

This endpoint allows you to add an audio or subtitle track to an existing media file using its `mediaId`. You need to provide the track `url` along with its `type` (audio or subtitle), `languageName` and `languageCode` in the request payload.


#### How it works

1. Send a POST request to this endpoint, replacing `{mediaId}` with the media ID (`uploadId` or `id`).

2. Provide the necessary details in the request body.

3. Receive a response containing a unique track ID and the details of the newly added track.


#### Webhook events

1. After successfully adding a track, your system must receive the webhook event <a href="https://docs.fastpix.io/docs/transform-media-events#videomediatrackcreated">video.media.track.created</a>.

2. Once the track is processed and ready, you must receive the webhook event <a href="https://docs.fastpix.io/docs/transform-media-events#videomediatrackready">video.media.track.ready</a>.

3. Finally, an update event <a href="https://docs.fastpix.io/docs/media-events#videomediaupdated">video.media.updated</a> must notify your system about the media's updated status.


#### Example
Suppose you have a video uploaded to the FastPix platform, and you want to add an Italian audio track to it. By calling this API, you can attach an external audio file (https://static.fastpix.io/music-1.mp3) to the media file. Similarly, if you need to add subtitles in different languages, you can specify type: `subtitle` with the corresponding subtitle `url`, `languageCode` and `languageName`.

Related guides: <a href="https://docs.fastpix.io/docs/manage-subtitle-tracks">Add own subtitle tracks</a>, <a href="https://docs.fastpix.io/docs/manage-audio-tracks">Add own audio tracks</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="Add-media-track" method="post" path="/on-demand/{mediaId}/tracks" -->
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
	"github.com/FastPix/fastpix-go/models/operations"
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

    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Videos.AddMediaTrack(ctx, "your-media-id", operations.AddMediaTrackRequestBody{
        Tracks: components.AddTrackRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `mediaID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | The unique identifier assigned to the media when created. The value must be a valid UUID.  | 4fa85f64-5717-4562-b3fc-2c963f66afa6                                                       |
| `body`                                                                                     | [operations.AddMediaTrackRequestBody](../../models/operations/addmediatrackrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.AddMediaTrackResponse](../../models/operations/addmediatrackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteTrack

This endpoint allows you to delete an existing audio or subtitle track from a media file. Once deleted, the track must no longer be available for playback.


#### How it works


1. Send a DELETE request to this endpoint, replacing `{mediaId}` with the media ID, and `{trackId}` with the ID of the track you want to remove.

2. The track gets deleted from the media file, and you must receive a confirmation response.

#### Webhook events

1. After successfully deleting a track, your system must receive the webhook event **video.media.track.deleted**.

2. Once the media file is updated to reflect the track removal, a <a href="https://docs.fastpix.io/docs/media-events#videomediaupdated">video.media.updated</a> event must be triggered.


#### Example
Suppose you uploaded an audio track in Italian for a video but later realize it's incorrect or no longer needed. By calling this API, you can remove the specific track while keeping the rest of the media file unchanged. This is useful when:

  - A track was mistakenly added and needs to be removed.
  - The content owner requests the removal of a specific subtitle or audio track.
  - A new version of the track gets uploaded to replace the existing one.

Related guides: <a href="https://docs.fastpix.io/docs/manage-subtitle-tracks">Add own subtitle tracks</a>, <a href="https://docs.fastpix.io/docs/manage-audio-tracks">Add own audio tracks</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-media-track" method="delete" path="/on-demand/{mediaId}/tracks/{trackId}" -->
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

    // your-media-id: FastPix Media ID returned from upload/create API
    // your-track-id: Track ID returned from add track API
    res, err := s.Videos.DeleteTrack(ctx, "your-media-id", "your-track-id")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `mediaID`                                                                                 | *string*                                                                                  | :heavy_check_mark:                                                                        | The unique identifier assigned to the media when created. The value must be a valid UUID. | your-media-id                                                      |
| `trackID`                                                                                 | *string*                                                                                  | :heavy_check_mark:                                                                        | The unique identifier assigned to the media when created. The value must be a valid UUID. | your-track-id                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.DeleteMediaTrackResponse](../../models/operations/deletemediatrackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateSourceAccess

This endpoint allows you to update the `sourceAccess` setting of an existing media file. The `sourceAccess` parameter determines whether the original media file is accessible or restricted. Setting this to `true` enables access to the media source, while setting it to `false` restricts access. 

#### How it works

1. Make a `PATCH` request to this endpoint, replacing `{mediaId}` with the ID of the media you want to update.

2. Include the updated `sourceAccess` parameter in the request body.

3. You receive a response confirming the update to the media’s source access status.
4. Webhook events: <a href="https://docs.fastpix.io/docs/transform-media-events#videomediasourceready">video.media.source.ready</a>, <a href="https://docs.fastpix.io/docs/transform-media-events#videomediasourcedeleted">video.media.source.deleted</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="updated-source-access" method="patch" path="/on-demand/{mediaId}/source-access" -->
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
	"github.com/FastPix/fastpix-go/models/operations"
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

    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Videos.UpdateSourceAccess(ctx, "your-media-id", operations.UpdatedSourceAccessRequestBody{
        SourceAccess: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `mediaID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier assigned to the media when created. The value must be a valid UUID.<br/>         | 4fa85f64-5717-4562-b3fc-2c963f66afa6                                                                   |
| `body`                                                                                                 | [operations.UpdatedSourceAccessRequestBody](../../models/operations/updatedsourceaccessrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    | {<br/>"sourceAccess": true<br/>}                                                                       |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdatedSourceAccessResponse](../../models/operations/updatedsourceaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMp4Support

This endpoint allows you to update the `mp4Support` setting of an existing media file using its media ID. You can specify the MP4 support level, such as `none`, `capped_4k`, `audioOnly`, or a combination of `audioOnly`, `capped_4k`, in the request payload.

#### How it works

1. Send a PATCH request to this endpoint, replacing `{mediaId}` with the media ID.

2. Provide the desired `mp4Support` value in the request body.

3. You receive a response confirming the update, including the media’s updated MP4 support status.

#### MP4 Support Options

- `none` – MP4 support is disabled for this media.

- `capped_4k` – Generates MP4 renditions up to 4K resolution.

- `audioOnly` – Generates an M4A file that contains only the audio track.

- `audioOnly,capped_4k` – Generates both an audio-only M4A file and MP4 renditions up to 4K resolution.

#### Webhook events

- <a href="https://docs.fastpix.io/docs/transform-media-events#videomediamp4supportready">video.media.mp4Support.ready</a> – Triggered when the MP4 support setting is successfully updated.

#### Example
Suppose you have a video uploaded to the FastPix platform, and you want to allow users to download the video in MP4 format. By setting "mp4Support": "capped_4k", the system generates an MP4 rendition of the video up to 4K resolution, making it available for download through the stream URL(`https://stream.fastpix.io/{playbackId}/{capped-4k.mp4 | audio.m4a}`). If you want users to stream only the audio from the media file, you can set "mp4Support": "audioOnly". This provides an audio-only stream URL that allows users to listen to the media without video. By setting "mp4Support": "audioOnly,capped_4k", both options are enabled. Users can download the MP4 video and also stream just the audio version of the media. 


Related guide: <a href="https://docs.fastpix.io/docs/mp4-support-for-offline-viewing">Use MP4 support for offline viewing</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="updated-mp4Support" method="patch" path="/on-demand/{mediaId}/update-mp4Support" -->
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
	"github.com/FastPix/fastpix-go/models/operations"
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

    // your-media-id: FastPix Media ID returned from upload/create API
    res, err := s.Videos.UpdateMp4Support(ctx, "your-media-id", operations.UpdatedMp4SupportRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
                responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
                if err != nil {
                    log.Printf("Error marshaling response: %v", err)
                    fmt.Printf("Response: %+v\n", res.Object)
                } else {
                    fmt.Println(string(responseJSON))
                }
            }
        } else {
            responseJSON, err := json.MarshalIndent(res.Object, "", "  ")
            if err != nil {
                log.Printf("Error marshaling response: %v", err)
                fmt.Printf("Response: %+v\n", res.Object)
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
| `mediaID`                                                                                          | *string*                                                                                           | :heavy_check_mark:                                                                                 | The unique identifier assigned to the media when created. The value must be a valid UUID.<br/>     | 4fa85f64-5717-4562-b3fc-2c963f66afa6                                                               |
| `body`                                                                                             | [operations.UpdatedMp4SupportRequestBody](../../models/operations/updatedmp4supportrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                | {<br/>"mp4Support": "capped_4k"<br/>}                                                              |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.UpdatedMp4SupportResponse](../../models/operations/updatedmp4supportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |