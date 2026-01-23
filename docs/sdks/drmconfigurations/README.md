# DrmConfigurations

## Overview

### Available Operations

* [List](#list) - Get list of DRM configuration IDs
* [GetByID](#getbyid) - Get DRM configuration by ID

## List


This endpoint retrieves the DRM configuration (DRM ID) associated with a workspace. It returns a list of DRM configurations, identified by a unique DRM ID, which is used for creating DRM encrypted asset.

**How it works:**
1. Make a GET request to this endpoint.  
2. Optionally use the `offset` and `limit` query parameters to paginate through the list of DRM configurations.  
3. The response includes a list of DRM IDs and pagination metadata.

**Example:**  
A media service provider may retrieve DRM configuration for a workspace to create DRM content.

Related guide: <a href="https://docs.fastpix.io/docs/secure-playback-with-drm">Manage DRM configuration</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="getDrmConfiguration" method="get" path="/on-demand/drm-configurations" -->
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

    res, err := s.DrmConfigurations.List(ctx, fastpixgo.Pointer[int64](1), fastpixgo.Pointer[int64](10))
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `offset`                                                                         | **int64*                                                                         | :heavy_minus_sign:                                                               | Offset determines the starting point for data retrieval within a paginated list. | 1                                                                                |
| `limit`                                                                          | **int64*                                                                         | :heavy_minus_sign:                                                               | Limit specifies the maximum number of items to display per page.                 | 10                                                                               |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetDrmConfigurationResponse](../../models/operations/getdrmconfigurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByID


This endpoint retrieves a DRM configuration ID. It is used to fetch the DRM-related ID for a workspace, typically required when validating or applying DRM policies to video assets.

**How it works:**
1. Make a GET request to this endpoint, replacing `{drmConfigurationId}` with the UUID of the DRM configuration.  
2. The response contains the associated DRM configuration ID.

Related guide: <a href="https://docs.fastpix.io/docs/secure-playback-with-drm">Manage DRM configuration</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="getDrmConfigurationById" method="get" path="/on-demand/drm-configurations/{drmConfigurationId}" -->
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

    // your-drm-configuration-id: DRM configuration ID returned from create DRM configuration API
    res, err := s.DrmConfigurations.GetByID(ctx, "your-drm-configuration-id")
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `drmConfigurationID`                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the DRM configuration.          | your-drm-configuration-id                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetDrmConfigurationByIDResponse](../../models/operations/getdrmconfigurationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |