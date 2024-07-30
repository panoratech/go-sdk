# UnifiedTicketingUserOutput


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Name`                                                                                                                   | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The name of the user                                                                                                     |
| `EmailAddress`                                                                                                           | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The email address of the user                                                                                            |
| `Teams`                                                                                                                  | []*string*                                                                                                               | :heavy_minus_sign:                                                                                                       | The teams whose the user is part of                                                                                      |
| `AccountID`                                                                                                              | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The account or organization the user is part of                                                                          |
| `FieldMappings`                                                                                                          | [components.UnifiedTicketingUserOutputFieldMappings](../../models/components/unifiedticketinguseroutputfieldmappings.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `ID`                                                                                                                     | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The UUID of the user                                                                                                     |
| `RemoteID`                                                                                                               | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The id of the user in the context of the 3rd Party                                                                       |
| `RemoteData`                                                                                                             | [components.UnifiedTicketingUserOutputRemoteData](../../models/components/unifiedticketinguseroutputremotedata.md)       | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `CreatedAt`                                                                                                              | [components.UnifiedTicketingUserOutputCreatedAt](../../models/components/unifiedticketinguseroutputcreatedat.md)         | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `ModifiedAt`                                                                                                             | [components.UnifiedTicketingUserOutputModifiedAt](../../models/components/unifiedticketinguseroutputmodifiedat.md)       | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |