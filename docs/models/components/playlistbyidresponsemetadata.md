# PlaylistByIDResponseMetadata

Required when the playlist type is `smart`. Media created between `startDate` and `endDate` of `createdDate` is added. Optionally, you can include media based on `updatedDate`.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreatedDate`                                                 | [*components.DateRange](../../models/components/daterange.md) | :heavy_minus_sign:                                            | Date range with start and end dates.                          |
| `UpdatedDate`                                                 | [*components.DateRange](../../models/components/daterange.md) | :heavy_minus_sign:                                            | Date range with start and end dates.                          |