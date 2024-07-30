# CreateCrmNoteRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `XConnectionToken`                                                               | *string*                                                                         | :heavy_check_mark:                                                               | The connection token                                                             |
| `RemoteData`                                                                     | **bool*                                                                          | :heavy_minus_sign:                                                               | Set to true to include data from the original Crm software.                      |
| `UnifiedCrmNoteInput`                                                            | [components.UnifiedCrmNoteInput](../../models/components/unifiedcrmnoteinput.md) | :heavy_check_mark:                                                               | N/A                                                                              |