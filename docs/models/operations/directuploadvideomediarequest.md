# DirectUploadVideoMediaRequest

Request body for direct upload


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `CorsOrigin`                                                                      | *string*                                                                          | :heavy_check_mark:                                                                | Upload media directly from a device using the URL name or enter '*' to allow all. | *                                                                                 |
| `PushMediaSettings`                                                               | [*operations.PushMediaSettings](../../models/operations/pushmediasettings.md)     | :heavy_minus_sign:                                                                | Configuration settings for media upload.                                          |                                                                                   |