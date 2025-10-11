# DuplicateReferenceIDErrorResponseError

Displays details about the reasons behind the request's failure.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Code`                                                        | **int64*                                                      | :heavy_minus_sign:                                            | Displays the error code indicating the type of the error.     | 409                                                           |
| `Message`                                                     | **string*                                                     | :heavy_minus_sign:                                            | A descriptive message providing more details for the error.   | playlist create failed                                        |
| `Description`                                                 | **string*                                                     | :heavy_minus_sign:                                            | A detailed explanation of the possible causes for the error.<br/> | A playlist with the given reference ID already exists.        |