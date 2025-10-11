# GetDataViewlistCurrentViewsGetTimeseriesViewsData


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `IntervalTime`                                      | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | The timestamp for the interval (ISO 8601 format).   |
| `MetricValue`                                       | **int64*                                            | :heavy_minus_sign:                                  | Reserved for future metric values (currently null). |
| `NumberOfViews`                                     | **int64*                                            | :heavy_minus_sign:                                  | Number of concurrent viewers at the given interval. |