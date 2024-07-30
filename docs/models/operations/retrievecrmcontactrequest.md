# RetrieveCrmContactRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `XConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the `contact` you want to retrive.                    |
| `RemoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original CRM software. |