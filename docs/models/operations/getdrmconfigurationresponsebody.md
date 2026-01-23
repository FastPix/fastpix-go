# GetDrmConfigurationResponseBody

DRM configuration(s) retrieved successfully


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Success`                                                                      | **bool*                                                                        | :heavy_minus_sign:                                                             | Shows the request status. Returns true for success and false for failure.      | true                                                                           |
| `Data`                                                                         | [][components.DrmIDResponse](../../models/components/drmidresponse.md)         | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `Pagination`                                                                   | [*components.Pagination](../../models/components/pagination.md)                | :heavy_minus_sign:                                                             | Pagination organizes content into pages for better readability and navigation. |                                                                                |