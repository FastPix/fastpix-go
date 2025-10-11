# MediaClipResponsePagination


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `TotalRecords`                                        | **int64*                                              | :heavy_minus_sign:                                    | Total number of records available.                    | 4                                                     |
| `CurrentOffset`                                       | **int64*                                              | :heavy_minus_sign:                                    | The starting offset of the current result set.        | 1                                                     |
| `OffsetCount`                                         | **int64*                                              | :heavy_minus_sign:                                    | The number of items returned in the current response. | 4                                                     |