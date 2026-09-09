# LivePlayback

## Overview

### Available Operations

* [Create](#create) - Create a playbackId
* [GetPlaybackIDDetails](#getplaybackiddetails) - Get playbackId details
* [DeletePlaybackID](#deleteplaybackid) - Delete a playbackId
* [UpdateDomainRestrictions](#updatedomainrestrictions) - Update domain restrictions for a live playback ID
* [UpdateUserAgentRestrictions](#updateuseragentrestrictions) - Update user-agent restrictions for a live playback ID

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

## UpdateDomainRestrictions

This endpoint allows updating domain restrictions for a specific playback ID associated with a live stream. 
It can be used to allow or deny specific domains during playback request evaluation.

**How it works:**
1. Make a `PATCH` request to this endpoint with your desired domain access configuration.
2. Specify a default policy (`allow` or `deny`) and provide specific `allow` or `deny` lists.
3. Use this to restrict which websites may embed the stream.

**Example:**
A developer may configure a live playback ID to allow embedding only on their own domain by setting the default policy to `deny` and listing the allowed domains.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-live-stream-domain-restrictions" method="patch" path="/live/streams/{streamId}/playback-ids/{playbackId}/domains" -->
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

    // your-stream-id: Live stream ID returned from the create stream API
    // your-playback-id: Playback ID returned from create playback ID API
    res, err := s.LivePlayback.UpdateDomainRestrictions(ctx, "your-stream-id", "your-playback-id", operations.UpdateLiveStreamDomainRestrictionsRequestBody{
        Allow: []string{
            "yourdomain.com", // Domain or wildcard pattern to allow (e.g., "*.sampledomain.com")
        },
        Deny: []string{
            "malicioussite.io", // Domain to deny
        },
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `streamID`                                                                                                              | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-stream-id                                                                                  |
| `playbackID`                                                                                                           | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-playback-id                                                                                   |
| `body`                                                                                                                 | [operations.UpdateLiveStreamDomainRestrictionsRequestBody](../../models/operations/updatelivestreamdomainrestrictionsrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.UpdateLiveStreamDomainRestrictionsResponse](../../models/operations/updatelivestreamdomainrestrictionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateUserAgentRestrictions

This endpoint allows updating user-agent restrictions for a specific playback ID associated with a live stream. 
It can be used to allow or deny specific user-agents during playback request evaluation.

**How it works:**
1. Make a `PATCH` request to this endpoint with your desired user-agent access configuration.
2. Specify a default policy (`allow` or `deny`) and provide specific `allow` or `deny` lists.
3. Use this to restrict access to specific browsers, devices, or bots.

**Example:**
A developer may configure a playback ID to deny access from known scraping user-agents while allowing all others by default.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-live-stream-user-agent-restrictions" method="patch" path="/live/streams/{streamId}/playback-ids/{playbackId}/user-agents" -->
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

    // your-stream-id: Live stream ID returned from the create stream API
    // your-playback-id: Playback ID returned from create playback ID API
    res, err := s.LivePlayback.UpdateUserAgentRestrictions(ctx, "your-stream-id", "your-playback-id", operations.UpdateLiveStreamUserAgentRestrictionsRequestBody{
        Allow: []string{
            "your-allowed-user-agent", // User agent string to allow (e.g., "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
        },
        Deny: []string{
            "your-denied-user-agent", // User agent string to deny (e.g., "Mozilla/5.0 (compatible; bot)")
        },
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `streamID`                                                                                                              | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-stream-id                                                                                  |
| `playbackID`                                                                                                           | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-playback-id                                                                                   |
| `body`                                                                                                                 | [operations.UpdateLiveStreamUserAgentRestrictionsRequestBody](../../models/operations/updatelivestreamuseragentrestrictionsrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.UpdateLiveStreamUserAgentRestrictionsResponse](../../models/operations/updatelivestreamuseragentrestrictionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |
