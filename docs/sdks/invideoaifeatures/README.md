# InVideoAIFeatures
(*InVideoAIFeatures*)

## Overview

### Available Operations

* [UpdateMediaSummary](#updatemediasummary) - Generate video summary
* [UpdateMediaChapters](#updatemediachapters) - Generate video chapters
* [UpdateMediaNamedEntities](#updatemedianamedentities) - Generate named entities
* [UpdateMediaModeration](#updatemediamoderation) - Enable video moderation

## UpdateMediaSummary

This endpoint allows you to generate the summary for an existing media.

#### How it works
1. Send a PATCH request to this endpoint, replacing `<mediaId>` with the unique ID of the media for which you wish to generate a summary.
2. Include the `generate` parameter in the request body.
3. Include the `summaryLength` parameter, specify the desired length of the summary in words (e.g., 120 words), this determines how concise or detailed the summary will be. If no specific summary length is provided, the default length will be 100 words. 
4. The response will include the updated media data and confirmation of the changes applied.

You can use the <a href="https://docs.fastpix.io/docs/ai-events#videomediaaisummaryready">video.mediaAI.summary.ready</a> webhook event to track and notify about the summary generation.





**Use case**: This is particularly useful when a user uploads a video and later chooses to generate a summary without needing to re-upload the video.

Related guide: <a href="https://docs.fastpix.io/docs/generate-video-summary">Video summary</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="update-media-summary" method="patch" path="/on-demand/{mediaId}/summary" -->
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

    res, err := s.InVideoAIFeatures.UpdateMediaSummary(ctx, "4fa85f64-5717-4562-b3fc-2c963f66afa6", operations.UpdateMediaSummaryRequestBody{
        Generate: true,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `mediaID`                                                                                            | *string*                                                                                             | :heavy_check_mark:                                                                                   | The unique identifier assigned to the media when created. The value should be a valid UUID.<br/>     | 4fa85f64-5717-4562-b3fc-2c963f66afa6                                                                 |
| `requestBody`                                                                                        | [operations.UpdateMediaSummaryRequestBody](../../models/operations/updatemediasummaryrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  | {<br/>"generate": true,<br/>"summaryLength": 100<br/>}                                               |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.UpdateMediaSummaryResponse](../../models/operations/updatemediasummaryresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ForbiddenError          | 403                               | application/json                  |
| apierrors.MediaNotFoundError      | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## UpdateMediaChapters

This endpoint enables you to generate chapters for an existing media file.

#### How it works
1. Make a `PATCH` request to this endpoint, replacing `<mediaId>` with the ID of the media for which you want to generate chapters.
2. Include the `chapters` parameter in the request body to enable.
3. The response will contain the updated media data, confirming the changes made.

You can use the <a href="https://docs.fastpix.io/docs/ai-events#videomediaaichaptersready">video.mediaAI.chapters.ready</a> webhook event to track and notify about the chapters generation.

**Use case:** This is particularly useful when a user uploads a video and later decides to enable chapters without re-uploading the entire video.

Related guide: <a href="https://docs.fastpix.io/reference/update-media-chapters">Video chapters</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="update-media-chapters" method="patch" path="/on-demand/{mediaId}/chapters" -->
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

    res, err := s.InVideoAIFeatures.UpdateMediaChapters(ctx, "4fa85f64-5717-4562-b3fc-2c963f66afa6", operations.UpdateMediaChaptersRequestBody{
        Chapters: true,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `mediaID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier assigned to the media when created. The value should be a valid UUID.<br/>       | 4fa85f64-5717-4562-b3fc-2c963f66afa6                                                                   |
| `requestBody`                                                                                          | [operations.UpdateMediaChaptersRequestBody](../../models/operations/updatemediachaptersrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    | {<br/>"chapters": true<br/>}                                                                           |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdateMediaChaptersResponse](../../models/operations/updatemediachaptersresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ForbiddenError          | 403                               | application/json                  |
| apierrors.MediaNotFoundError      | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## UpdateMediaNamedEntities

This endpoint allows you to extract named entities from an existing media.
Named Entity Recognition (NER) is a fundamental natural language processing (NLP) technique that identifies and classifies key information (entities) in text into predefined categories. For instance:

  - Organizations (e.g., "Microsoft", "United Nations")
  - Locations (e.g., "Paris", "Mount Everest")
  - Product names (e.g., "iPhone", "Coca-Cola")

#### How it works
1. Make a PATCH request to this endpoint, replacing `<mediaId>` with the ID of the media you want to extract named-entities.
2. Include the `namedEntities` parameter in the request body to enable.
3. Receive a response containing the updated media data, confirming the changes made.

You can use the <a href="https://docs.fastpix.io/docs/ai-events#videomediaainamedentitiesready">video.mediaAI.named-entities.ready</a> webhook event to track and notify about the named entities extraction.

**Use case:** If a user uploads a video and later decides to enable named entity extraction without re-uploading the entire video.

Related guide: <a href="https://docs.fastpix.io/docs/generate-named-entities">Named entities</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="update-media-named-entities" method="patch" path="/on-demand/{mediaId}/named-entities" -->
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

    res, err := s.InVideoAIFeatures.UpdateMediaNamedEntities(ctx, "0cec3c88-c69d-4232-9b96-f0976327fa2d", operations.UpdateMediaNamedEntitiesRequestBody{
        NamedEntities: true,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `mediaID`                                                                                                        | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The unique identifier assigned to the media when created. The value should be a valid UUID.<br/>                 | 0cec3c88-c69d-4232-9b96-f0976327fa2d                                                                             |
| `requestBody`                                                                                                    | [operations.UpdateMediaNamedEntitiesRequestBody](../../models/operations/updatemedianamedentitiesrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              | {<br/>"namedEntities": true<br/>}                                                                                |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.UpdateMediaNamedEntitiesResponse](../../models/operations/updatemedianamedentitiesresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ForbiddenError          | 403                               | application/json                  |
| apierrors.MediaNotFoundError      | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## UpdateMediaModeration

This endpoint enables moderation features, such as NSFW and profanity filtering, to detect inappropriate content in existing media.

#### How it works
1. Make a PATCH request to this endpoint, replacing `<mediaId>` with the ID of the media you want to update.
2. Include the `moderation` object and provide the requried `type` parameter in the request body to specify the media type (e.g., video/audio/av).
4. The response will contain the updated media data, confirming the changes made.

You can use the <a href="https://docs.fastpix.io/docs/ai-events#videomediaaimoderationready">video.mediaAI.moderation.ready</a> webhook event to track and notify about the detected moderation results.

**Use case:** This is particularly useful when a user uploads a video and later decides to enable moderation detection without the need to re-upload it.

Related guide: <a href="https://docs.fastpix.io/docs/using-nsfw-and-profanity-filter-for-video-moderation">Moderate NSFW & Profanity</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="update-media-moderation" method="patch" path="/on-demand/{mediaId}/moderation" -->
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

    res, err := s.InVideoAIFeatures.UpdateMediaModeration(ctx, "0cec3c88-c69d-4232-9b96-f0976327fa2d", operations.UpdateMediaModerationRequestBody{
        Moderation: &operations.UpdateMediaModerationModeration{
            Type: components.MediaTypeVideo.ToPointer(),
        },
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `mediaID`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The unique identifier assigned to the media when created. The value should be a valid UUID.<br/>           | 0cec3c88-c69d-4232-9b96-f0976327fa2d                                                                       |
| `requestBody`                                                                                              | [operations.UpdateMediaModerationRequestBody](../../models/operations/updatemediamoderationrequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        | {<br/>"moderation": {<br/>"type": "video"<br/>}<br/>}                                                      |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.UpdateMediaModerationResponse](../../models/operations/updatemediamoderationresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ForbiddenError          | 403                               | application/json                  |
| apierrors.MediaNotFoundError      | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |