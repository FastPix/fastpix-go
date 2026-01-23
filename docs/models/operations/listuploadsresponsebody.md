# ListUploadsResponseBody

List of video media


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Success`                                                                        | **bool*                                                                          | :heavy_minus_sign:                                                               | Shows the request status. Returns true for success and false for failure.        | true                                                                             |
| `Data`                                                                           | [][components.UnusedDirectUpload](../../models/components/unuseddirectupload.md) | :heavy_minus_sign:                                                               | Displays the result of the request.                                              |                                                                                  |
| `Pagination`                                                                     | [*components.Pagination](../../models/components/pagination.md)                  | :heavy_minus_sign:                                                               | Pagination organizes content into pages for better readability and navigation.   |                                                                                  |