# UnifiedAtsActivityOutput


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ActivityType`                                                                                                       | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The type of activity                                                                                                 |
| `Subject`                                                                                                            | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The subject of the activity                                                                                          |
| `Body`                                                                                                               | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The body of the activity                                                                                             |
| `Visibility`                                                                                                         | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The visibility of the activity                                                                                       |
| `CandidateID`                                                                                                        | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The UUID of the candidate                                                                                            |
| `RemoteCreatedAt`                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | The remote creation date of the activity                                                                             |
| `FieldMappings`                                                                                                      | [components.UnifiedAtsActivityOutputFieldMappings](../../models/components/unifiedatsactivityoutputfieldmappings.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `ID`                                                                                                                 | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The UUID of the activity                                                                                             |
| `RemoteID`                                                                                                           | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The remote ID of the activity in the context of the 3rd Party                                                        |
| `RemoteData`                                                                                                         | [components.UnifiedAtsActivityOutputRemoteData](../../models/components/unifiedatsactivityoutputremotedata.md)       | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `CreatedAt`                                                                                                          | [components.UnifiedAtsActivityOutputCreatedAt](../../models/components/unifiedatsactivityoutputcreatedat.md)         | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `ModifiedAt`                                                                                                         | [components.UnifiedAtsActivityOutputModifiedAt](../../models/components/unifiedatsactivityoutputmodifiedat.md)       | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |