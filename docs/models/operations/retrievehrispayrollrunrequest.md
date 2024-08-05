# RetrieveHrisPayrollRunRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `XConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |                                                              |
| `ID`                                                         | *string*                                                     | :heavy_check_mark:                                           | id of the payroll run you want to retrieve.                  | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                         |
| `RemoteData`                                                 | **bool*                                                      | :heavy_minus_sign:                                           | Set to true to include data from the original Hris software. | false                                                        |