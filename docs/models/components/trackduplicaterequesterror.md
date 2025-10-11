# TrackDuplicateRequestError

Displays details about the reasons behind the request's failure.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Code`                                                        | **int64*                                                      | :heavy_minus_sign:                                            | Displays the error code indicating the type of the error.     | 400                                                           |
| `Message`                                                     | **string*                                                     | :heavy_minus_sign:                                            | A descriptive message providing more details for the error.   | duplicate language name                                       |
| `Description`                                                 | **string*                                                     | :heavy_minus_sign:                                            | A detailed explanation of the possible causes for the error.<br/> | Duplicate language name exists for the given track ID.        |