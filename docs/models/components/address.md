# Address


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Street1`                                                        | *string*                                                         | :heavy_check_mark:                                               | The street                                                       |
| `Street2`                                                        | *string*                                                         | :heavy_check_mark:                                               | More information about the street                                |
| `City`                                                           | *string*                                                         | :heavy_check_mark:                                               | The city                                                         |
| `State`                                                          | *string*                                                         | :heavy_check_mark:                                               | The state                                                        |
| `PostalCode`                                                     | *string*                                                         | :heavy_check_mark:                                               | The postal code                                                  |
| `Country`                                                        | *string*                                                         | :heavy_check_mark:                                               | The country                                                      |
| `AddressType`                                                    | [components.AddressType](../../models/components/addresstype.md) | :heavy_check_mark:                                               | The address type. Authorized values are either PERSONAL or WORK. |
| `OwnerType`                                                      | *string*                                                         | :heavy_check_mark:                                               | The owner type of the address                                    |