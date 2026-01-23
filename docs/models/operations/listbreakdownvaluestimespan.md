# ListBreakdownValuesTimespan

This parameter specifies the time span between which the video views list must be retrieved by. You can provide either from and to unix epoch timestamps or time duration. The scope of duration is between 60 minutes to 30 days.

**Accepted formats are:**

array of epoch timestamps for example 
`timespan[]=1498867200&timespan[]=1498953600`

duration string for example  
`timespan[]=24:hours` or `timespan[]=7:days`



## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ListBreakdownValuesTimespanSixtyminutes`    | 60:minutes                                   |
| `ListBreakdownValuesTimespanSixhours`        | 6:hours                                      |
| `ListBreakdownValuesTimespanTwentyFourhours` | 24:hours                                     |
| `ListBreakdownValuesTimespanThreedays`       | 3:days                                       |
| `ListBreakdownValuesTimespanSevendays`       | 7:days                                       |
| `ListBreakdownValuesTimespanThirtydays`      | 30:days                                      |