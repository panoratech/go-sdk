# ListAtsOfferRequest


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `XConnectionToken`                                      | *string*                                                | :heavy_check_mark:                                      | The connection token                                    |
| `RemoteData`                                            | **bool*                                                 | :heavy_minus_sign:                                      | Set to true to include data from the original software. |
| `Limit`                                                 | **float64*                                              | :heavy_minus_sign:                                      | Set to get the number of records.                       |
| `Cursor`                                                | **string*                                               | :heavy_minus_sign:                                      | Set to get the number of records after this cursor.     |