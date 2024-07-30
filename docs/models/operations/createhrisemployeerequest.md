# CreateHrisEmployeeRequest


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `XConnectionToken`                                                                         | *string*                                                                                   | :heavy_check_mark:                                                                         | The connection token                                                                       |
| `RemoteData`                                                                               | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Set to true to include data from the original Hris software.                               |
| `UnifiedHrisEmployeeInput`                                                                 | [components.UnifiedHrisEmployeeInput](../../models/components/unifiedhrisemployeeinput.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |