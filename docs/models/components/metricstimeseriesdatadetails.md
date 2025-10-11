# MetricsTimeseriesDataDetails

The metric's value at specific time intervals.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `IntervalTime`                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | The timestamp for the data point indicating when the metric value was recorded. | 2023-12-04T14:00:00.000Z                                                        |
| `MetricValue`                                                                   | [*components.MetricValue](../../models/components/metricvalue.md)               | :heavy_minus_sign:                                                              | The value of the specified metric at the given interval.                        | 0.793110142151515                                                               |
| `NumberOfViews`                                                                 | **int64*                                                                        | :heavy_minus_sign:                                                              | The total number of views recorded during that interval.                        | 143244                                                                          |