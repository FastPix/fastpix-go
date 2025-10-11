# Metrics
(*Metrics*)

## Overview

### Available Operations

* [ListBreakdownValues](#listbreakdownvalues) - List breakdown values
* [ListOverallValues](#listoverallvalues) - List overall values
* [GetTimeseriesData](#gettimeseriesdata) - Get timeseries data
* [ListComparisonValues](#listcomparisonvalues) - List comparison values

## ListBreakdownValues

Retrieves breakdown values for a specified metric and timespan, allowing you to analyze the performance of your content based on various dimensions. It provides insights into how different factors contribute to the overall metrics. 

#### How it works

  1. Before using this endpoint, you can call the <a href="https://docs.fastpix.io/reference/list_dimensions">List Dimensions</a> endpoint to retrieve all available dimensions that can be used in your query. 

  2. Make a `GET` request to this endpoint with the required `metricId` and other query parameters. 

  3. Receive a response containing the breakdown values for the specified metric, grouped and filtered according to your parameters. 

  4. Upon successful retrieval, the response will include the breakdown values based on the specified parameters. Note that the time values ( `totalWatchTime` and `totalPlayingTime` ) are in milliseconds. 


#### Example


A developer wants to analyze how watch time varies across different device types. By calling this endpoint for the `playing_time` metric and filtering by `device_type`, they can understand how engagement differs between mobile, desktop, and tablet users. This data will guide optimization efforts for different platforms. 


Related guide: <a href="https://docs.fastpix.io/docs/metrics-overview">Understand data definitions</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list_breakdown_values" method="get" path="/data/metrics/{metricId}/breakdown" -->
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

    res, err := s.Metrics.ListBreakdownValues(ctx, operations.ListBreakdownValuesRequest{
        MetricID: operations.ListBreakdownValuesMetricIDQualityOfExperienceScore,
        Timespan: operations.ListBreakdownValuesTimespanSevendays,
        Filterby: fastpixgo.Pointer("browser_name:Chrome"),
        GroupBy: fastpixgo.Pointer("browser_name"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListBreakdownValuesRequest](../../models/operations/listbreakdownvaluesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListBreakdownValuesResponse](../../models/operations/listbreakdownvaluesresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ViewNotFoundError       | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## ListOverallValues

Retrieves overall values for a specified metric, providing summary statistics that help you understand the performance of your content. The response includes key metrics such as `totalWatchTime`, `uniqueViews`, `totalPlayTime` and `totalViews`. 

#### How it works

  1. Before using this endpoint, you can call the <a href="https://docs.fastpix.io/reference/list_dimensions">list dimensions</a> endpoint to retrieve all available dimensions that can be used in your query. 

  2. Make a `GET` request to this endpoint with the required `metricId` and other query parameters. 

  3. Receive a response containing the overall values for the specified metric, which may vary based on the applied filters. 






#### Key fields in response


  * **value:** The specific metric value calculated based on the applied filters. 
  * **totalWatchTime:** Total time watched across all views, represented in milliseconds. 
  * **uniqueViews:** The count of unique viewers who interacted with the content. 
  * **totalViews:** The total number of views recorded. 
  * **totalPlayTime:** Total time spent playing the video, represented in milliseconds. 
  * **globalValue:** A global metric value that reflects the overall performance of the specified metric across the entire dataset for the given timespan. This value is not affected by specific filters. 


  Related guide: <a href="https://docs.fastpix.io/docs/metrics-overview">Understand data definitions</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list_overall_values" method="get" path="/data/metrics/{metricId}/overall" -->
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

    res, err := s.Metrics.ListOverallValues(ctx, operations.ListOverallValuesMetricIDQualityOfExperienceScore, operations.ListOverallValuesTimespanSevendays, fastpixgo.Pointer("avg"), fastpixgo.Pointer("browser_name:Chrome"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                          |
| `metricID`                                                                                                                                                                                                                                                                                                               | [operations.ListOverallValuesMetricID](../../models/operations/listoverallvaluesmetricid.md)                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                       | Pass metric Id<br/>                                                                                                                                                                                                                                                                                                      | quality_of_experience_score                                                                                                                                                                                                                                                                                              |
| `timespan`                                                                                                                                                                                                                                                                                                               | [operations.ListOverallValuesTimespan](../../models/operations/listoverallvaluestimespan.md)                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                       | This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.<br/>                                                                                | 7:days                                                                                                                                                                                                                                                                                                                   |
| `measurement`                                                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The measurement for the given metrics.<br/>Possible Values : [95th, median, avg, count or sum]<br/>                                                                                                                                                                                                                      | avg                                                                                                                                                                                                                                                                                                                      |
| `filterby`                                                                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.<br/>Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]<br/> | browser_name:Chrome                                                                                                                                                                                                                                                                                                      |
| `opts`                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.ListOverallValuesResponse](../../models/operations/listoverallvaluesresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ViewNotFoundError       | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## GetTimeseriesData

This endpoint retrieves timeseries data for a specified metric, providing insights into how the metric values change over time. The response includes an array of data points, each representing the metric's value at specific intervals. 

Each data point contains the following fields: 

* **intervalTime:** The timestamp for the data point indicating when the metric value was recorded. 
* **metricValue:** The value of the specified metric at the given interval, reflecting the performance or engagement level during that time. 
* **numberOfViews:** The total number of views recorded during that interval, providing context for the metric value. 


### Example Usage

<!-- UsageSnippet language="go" operationID="get_timeseries_data" method="get" path="/data/metrics/{metricId}/timeseries" -->
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

    res, err := s.Metrics.GetTimeseriesData(ctx, operations.GetTimeseriesDataRequest{
        MetricID: operations.GetTimeseriesDataMetricIDQualityOfExperienceScore,
        Timespan: operations.GetTimeseriesDataTimespanSevendays,
        Filterby: fastpixgo.Pointer("browser_name:Chrome"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTimeseriesDataRequest](../../models/operations/gettimeseriesdatarequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetTimeseriesDataResponse](../../models/operations/gettimeseriesdataresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ViewNotFoundError       | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## ListComparisonValues

This endpoint allows you to compare multiple metrics across specified dimensions. You can specify the metrics you want to compare in the query parameters, and the response will include the relevant metrics for the specified dimensions. 

#### How it works 

  1. Before making a request to this endpoint, call the <a href="https://docs.fastpix.io/reference/list_dimensions">list dimensions</a> endpoint to obtain all available dimensions that can be used for comparison. 

  2. Make a `GET` request to this endpoint with the desired metrics specified in the query parameters. 

  3. Receive a response containing the comparison values for the specified metrics across the selected dimensions. 


  Related guide: <a href="https://docs.fastpix.io/docs/understand-dashboard-ui#compare-metrics">Compare metrics in dashboard</a>


### Example Usage

<!-- UsageSnippet language="go" operationID="list_comparison_values" method="get" path="/data/metrics/comparison" -->
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

    res, err := s.Metrics.ListComparisonValues(ctx, operations.ListComparisonValuesTimespanSevendays, fastpixgo.Pointer("browser_name:Chrome"), operations.ListComparisonValuesDimensionBrowserName.ToPointer(), fastpixgo.Pointer("Chrome"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                          |
| `timespan`                                                                                                                                                                                                                                                                                                               | [operations.ListComparisonValuesTimespan](../../models/operations/listcomparisonvaluestimespan.md)                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                                                                                       | This parameter specifies the time span between which the video views list should be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.<br/>                                                                                | 7:days                                                                                                                                                                                                                                                                                                                   |
| `filterby`                                                                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | Pass the dimensions and their corresponding values you want to filter the views by. For excluding the values in the filter we can pass '!' before the filter value. The list of filters can be obtained from list of dimensions endpoint.<br/>Example Values : [ browser_name:Chrome , os_name:macOS , device_name:Galaxy ]<br/> | browser_name:Chrome                                                                                                                                                                                                                                                                                                      |
| `dimension`                                                                                                                                                                                                                                                                                                              | [*operations.ListComparisonValuesDimension](../../models/operations/listcomparisonvaluesdimension.md)                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The dimension id in which the views are watched.<br/>                                                                                                                                                                                                                                                                    | browser_name                                                                                                                                                                                                                                                                                                             |
| `value`                                                                                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The value for the selected dimension. <br/>For example:<br/> If `dimension` is `browser_name`, the value could be  `Chrome` `,` `Firefox` `etc` .<br/> If `dimension` is `os_name`, the value could be `macOS` `,` `Windows` `etc` .<br/>                                                                                | Chrome                                                                                                                                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.ListComparisonValuesResponse](../../models/operations/listcomparisonvaluesresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.InvalidPermissionError  | 401                               | application/json                  |
| apierrors.ViewNotFoundError       | 404                               | application/json                  |
| apierrors.ValidationErrorResponse | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |