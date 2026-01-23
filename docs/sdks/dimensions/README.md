# Dimensions

## Overview

Operations involving dimensions

### Available Operations

* [List](#list) - List the dimensions
* [ListFilterValues](#listfiltervalues) - List the filter values for a dimension

## List

Retrieves a list of dimensions that can be used as query parameters across various data endpoints. Each dimension has a unique id that can be used to filter data effectively. 

The dimensions retrieved from this endpoint can be used in conjunction with the <a href="https://docs.fastpix.io/reference/list_video_views">list video views</a> and <a href="https://docs.fastpix.io/reference/list_by_top_content">list by top content</a> endpoints to filter results based on specific criteria. For example, you can filter views by `browser_name`, `os_name`, `device_type`, and more.

Related guides: <a href="https://docs.fastpix.io/page/what-video-data-do-we-capture#/">What Video Data do we capture?</a> ,   <a href="https://docs.fastpix.io/docs/user-passable-metadata-1">Use passable dimensions</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list_dimensions" method="get" path="/data/dimensions" -->
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

    res, err := s.Dimensions.List(ctx)
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListDimensionsResponse](../../models/operations/listdimensionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListFilterValues

This endpoint returns the filter values associated with a specific dimension, along with the total number of video views for each value. For example, it can list all `browser_name` (dimension) and show how many views occurred for all available browsers like Chrome, Safari (filter values). 


In order to use the <a href="https://docs.fastpix.io/docs/custom-business-metadata">Custom Dimensions</a>, you must enable them in the dashboard under settings option based on the plan you have opted for.

#### Example

A developer wants to know how their video content performs across different browsers. By calling this endpoint for the `device_type` dimension, they can retrieve a breakdown of video views by each device (for example, Desktop, Mobile, Tablet). This data helps the developer understand where optimizations or troubleshooting is necessary.


Related guide: <a href="https://docs.fastpix.io/docs/understand-dashboard-ui#filters-and-timeframes">Filters and timespan</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list_filter_values_for_dimension" method="get" path="/data/dimensions/{dimensionsId}" -->
```go
package main

import(
    "context"
	"encoding/json"
	"fmt"
	"log"
	"io"
	"bytes"

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

    res, err := s.Dimensions.ListFilterValues(ctx, operations.DimensionsIDBrowserName, operations.ListFilterValuesForDimensionTimespanTwentyFourhours.ToPointer(), fastpixgo.Pointer("browser_name:your-browser-name")) // Filter by dimension value (e.g., "browser_name:Chrome")
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

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `dimensionsID`                                                                                                                                                                                                                                                                                                                                                                                                                   | [operations.DimensionsID](../../models/operations/dimensionsid.md)                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                               | Pass Dimensions Id<br/>                                                                                                                                                                                                                                                                                                                                                                                                          | browser_name                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `timespan`                                                                                                                                                                                                                                                                                                                                                                                                                       | [*operations.ListFilterValuesForDimensionTimespan](../../models/operations/listfiltervaluesfordimensiontimespan.md)                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                               | This parameter specifies the time span between which the video views list must be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.<br/><br/>**Accepted formats are:**<br/><br/>array of epoch timestamps for example <br/>`timespan[]=1498867200&timespan[]=1498953600`<br/><br/>duration string for example  <br/>`timespan[]=24:hours` or `timespan[]=7:days`<br/> | 24:hours                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `filterby`                                                                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                               | Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass "!" before the filter value. The list of filters can be obtained from list of dimensions endpoint.<br/>Example Values : [ browser_name:Chrome , os_name:macOS , !device_name:Galaxy ]<br/>                                                                                                | browser_name:Chrome                                                                                                                                                                                                                                                                                                                                                                                                              |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                  |

### Response

**[*operations.ListFilterValuesForDimensionResponse](../../models/operations/listfiltervaluesfordimensionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |