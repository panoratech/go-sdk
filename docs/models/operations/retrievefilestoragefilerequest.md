# RetrieveFilestorageFileRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `XConnectionToken`                                                   | *string*                                                             | :heavy_check_mark:                                                   | The connection token                                                 |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | id of the file you want to retrieve.                                 |
| `RemoteData`                                                         | **bool*                                                              | :heavy_minus_sign:                                                   | Set to true to include data from the original File Storage software. |