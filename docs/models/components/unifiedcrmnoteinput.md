# UnifiedCrmNoteInput


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Content`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The content of the note                                                                                    |
| `UserID`                                                                                                   | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | The UUID of the user tied the note                                                                         |
| `CompanyID`                                                                                                | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | The UUID of the company tied to the note                                                                   |
| `ContactID`                                                                                                | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | The UUID fo the contact tied to the note                                                                   |
| `DealID`                                                                                                   | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | The UUID of the deal tied to the note                                                                      |
| `FieldMappings`                                                                                            | [components.UnifiedCrmNoteInputFieldMappings](../../models/components/unifiedcrmnoteinputfieldmappings.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |