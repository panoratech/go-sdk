# CreateAccountingExpenseRequest


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `XConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |
| `RemoteData`                                                                                         | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Set to true to include data from the original Accounting software.                                   |
| `UnifiedAccountingExpenseInput`                                                                      | [components.UnifiedAccountingExpenseInput](../../models/components/unifiedaccountingexpenseinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |