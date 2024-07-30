# UnifiedFilestorageFileOutput


## Fields

| Field                                                                                                                        | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                       | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The name of the file                                                                                                         |
| `FileURL`                                                                                                                    | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The url of the file                                                                                                          |
| `MimeType`                                                                                                                   | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The mime type of the file                                                                                                    |
| `Size`                                                                                                                       | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The size of the file                                                                                                         |
| `FolderID`                                                                                                                   | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The UUID of the folder tied to the file                                                                                      |
| `Permission`                                                                                                                 | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The UUID of the permission tied to the file                                                                                  |
| `SharedLink`                                                                                                                 | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The UUID of the shared link tied to the file                                                                                 |
| `FieldMappings`                                                                                                              | [components.UnifiedFilestorageFileOutputFieldMappings](../../models/components/unifiedfilestoragefileoutputfieldmappings.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `ID`                                                                                                                         | **string*                                                                                                                    | :heavy_minus_sign:                                                                                                           | The UUID of the file                                                                                                         |
| `RemoteID`                                                                                                                   | **string*                                                                                                                    | :heavy_minus_sign:                                                                                                           | The id of the file in the context of the 3rd Party                                                                           |
| `RemoteData`                                                                                                                 | [components.UnifiedFilestorageFileOutputRemoteData](../../models/components/unifiedfilestoragefileoutputremotedata.md)       | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `CreatedAt`                                                                                                                  | [components.UnifiedFilestorageFileOutputCreatedAt](../../models/components/unifiedfilestoragefileoutputcreatedat.md)         | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `ModifiedAt`                                                                                                                 | [components.UnifiedFilestorageFileOutputModifiedAt](../../models/components/unifiedfilestoragefileoutputmodifiedat.md)       | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |