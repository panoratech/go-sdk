# UnifiedFilestorageFolderOutput


## Fields

| Field                                                                                                                            | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                           | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The name of the folder                                                                                                           |
| `Size`                                                                                                                           | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The size of the folder                                                                                                           |
| `FolderURL`                                                                                                                      | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The url of the folder                                                                                                            |
| `Description`                                                                                                                    | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The description of the folder                                                                                                    |
| `DriveID`                                                                                                                        | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The UUID of the drive tied to the folder                                                                                         |
| `ParentFolderID`                                                                                                                 | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The UUID of the parent folder                                                                                                    |
| `SharedLink`                                                                                                                     | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The UUID of the shared link tied to the folder                                                                                   |
| `Permission`                                                                                                                     | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The UUID of the permission tied to the folder                                                                                    |
| `FieldMappings`                                                                                                                  | [components.UnifiedFilestorageFolderOutputFieldMappings](../../models/components/unifiedfilestoragefolderoutputfieldmappings.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `ID`                                                                                                                             | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | The UUID of the folder                                                                                                           |
| `RemoteID`                                                                                                                       | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | The id of the folder in the context of the 3rd Party                                                                             |
| `RemoteData`                                                                                                                     | [components.UnifiedFilestorageFolderOutputRemoteData](../../models/components/unifiedfilestoragefolderoutputremotedata.md)       | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `CreatedAt`                                                                                                                      | [components.UnifiedFilestorageFolderOutputCreatedAt](../../models/components/unifiedfilestoragefolderoutputcreatedat.md)         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `ModifiedAt`                                                                                                                     | [components.UnifiedFilestorageFolderOutputModifiedAt](../../models/components/unifiedfilestoragefolderoutputmodifiedat.md)       | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |