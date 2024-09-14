# ResyncStatusDto


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Timestamp`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | N/A                                                        |                                                            |
| `Vertical`                                                 | [components.Vertical](../../models/components/vertical.md) | :heavy_check_mark:                                         | N/A                                                        | ticketing                                                  |
| `Provider`                                                 | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | gitlab                                                     |
| `Status`                                                   | [components.Status](../../models/components/status.md)     | :heavy_check_mark:                                         | N/A                                                        | success                                                    |