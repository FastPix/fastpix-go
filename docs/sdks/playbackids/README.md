# PlaybackIds

## Overview

### Available Operations

* [UpdateUserAgentRestrictions](#updateuseragentrestrictions) - Update user-agent restrictions for a playback ID

## UpdateUserAgentRestrictions

This endpoint allows updating user-agent restrictions for a specific playback ID associated with a media asset. 
It can be used to allow or deny specific user-agents during playback request evaluation.

**How it works:**
1. Make a `PATCH` request to this endpoint with your desired user-agent access configuration.
2. Specify a default policy (`allow` or `deny`) and provide specific `allow` or `deny` lists.
3. Use this to restrict access to specific browsers, devices, or bots.

**Example:**
A developer may configure a playback ID to deny access from known scraping user-agents while allowing all others by default.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-user-agent-restrictions" method="patch" path="/on-demand/{mediaId}/playback-ids/{playbackId}/user-agents" -->
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

    // your-video-id: FastPix Video ID returned from upload/create API
    // your-playback-id: Playback ID returned from create playback ID API
    res, err := s.PlaybackIds.UpdateUserAgentRestrictions(ctx, "your-video-id", "your-playback-id", operations.UpdateUserAgentRestrictionsRequestBody{
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
| `mediaID`                                                                                                              | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-video-id                                                                                   |
| `playbackID`                                                                                                           | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    | your-playback-id                                                                                   |
| `body`                                                                                                                 | [operations.UpdateUserAgentRestrictionsRequestBody](../../models/operations/updateuseragentrestrictionsrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.UpdateUserAgentRestrictionsResponse](../../models/operations/updateuseragentrestrictionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |