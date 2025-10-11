# CreatePlaylistRequestMetadata

Required when playlist type is smart - media created between startDate and endDate of createdDate will be add, similarily updatedDate (Optional)


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreatedDate`                                                 | [*components.DateRange](../../models/components/daterange.md) | :heavy_minus_sign:                                            | Date range with start and end dates.                          |
| `UpdatedDate`                                                 | [*components.DateRange](../../models/components/daterange.md) | :heavy_minus_sign:                                            | Date range with start and end dates.                          |