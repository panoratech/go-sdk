# UnifiedAtsApplicationInput


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `AppliedAt`                                                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_minus_sign:                                                                                                       | The application date                                                                                                     |
| `RejectedAt`                                                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                                                               | :heavy_minus_sign:                                                                                                       | The rejection date                                                                                                       |
| `Offers`                                                                                                                 | []*string*                                                                                                               | :heavy_minus_sign:                                                                                                       | The offers UUIDs for the application                                                                                     |
| `Source`                                                                                                                 | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The source of the application                                                                                            |
| `CreditedTo`                                                                                                             | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The UUID of the person credited for the application                                                                      |
| `CurrentStage`                                                                                                           | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The UUID of the current stage of the application                                                                         |
| `RejectReason`                                                                                                           | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The rejection reason for the application                                                                                 |
| `CandidateID`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The UUID of the candidate                                                                                                |
| `JobID`                                                                                                                  | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | The UUID of the job                                                                                                      |
| `FieldMappings`                                                                                                          | [components.UnifiedAtsApplicationInputFieldMappings](../../models/components/unifiedatsapplicationinputfieldmappings.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |