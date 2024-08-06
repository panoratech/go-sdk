# ListFilestorageFileRequest


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `XConnectionToken`                                      | *string*                                                | :heavy_check_mark:                                      | The connection token                                    |                                                         |
| `RemoteData`                                            | **bool*                                                 | :heavy_minus_sign:                                      | Set to true to include data from the original software. | true                                                    |
| `Limit`                                                 | **float64*                                              | :heavy_minus_sign:                                      | Set to get the number of records.                       | 10                                                      |
| `Cursor`                                                | **string*                                               | :heavy_minus_sign:                                      | Set to get the number of records after this cursor.     | 1b8b05bb-5273-4012-b520-8657b0b90874                    |