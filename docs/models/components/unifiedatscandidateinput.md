# UnifiedAtsCandidateInput


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `FirstName`                                                                                                          | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The first name of the candidate                                                                                      |
| `LastName`                                                                                                           | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The last name of the candidate                                                                                       |
| `Company`                                                                                                            | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The company of the candidate                                                                                         |
| `Title`                                                                                                              | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The title of the candidate                                                                                           |
| `Locations`                                                                                                          | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The locations of the candidate                                                                                       |
| `IsPrivate`                                                                                                          | **bool*                                                                                                              | :heavy_minus_sign:                                                                                                   | Whether the candidate is private                                                                                     |
| `EmailReachable`                                                                                                     | **bool*                                                                                                              | :heavy_minus_sign:                                                                                                   | Whether the candidate is reachable by email                                                                          |
| `RemoteCreatedAt`                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | The remote creation date of the candidate                                                                            |
| `RemoteModifiedAt`                                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | The remote modification date of the candidate                                                                        |
| `LastInteractionAt`                                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | The last interaction date with the candidate                                                                         |
| `Attachments`                                                                                                        | []*string*                                                                                                           | :heavy_minus_sign:                                                                                                   | The attachments UUIDs of the candidate                                                                               |
| `Applications`                                                                                                       | []*string*                                                                                                           | :heavy_minus_sign:                                                                                                   | The applications UUIDs of the candidate                                                                              |
| `Tags`                                                                                                               | []*string*                                                                                                           | :heavy_minus_sign:                                                                                                   | The tags of the candidate                                                                                            |
| `Urls`                                                                                                               | [][components.URL](../../models/components/url.md)                                                                   | :heavy_minus_sign:                                                                                                   | The urls of the candidate, possible values for Url type are WEBSITE, BLOG, LINKEDIN, GITHUB, or OTHER                |
| `PhoneNumbers`                                                                                                       | [][components.Phone](../../models/components/phone.md)                                                               | :heavy_minus_sign:                                                                                                   | The phone numbers of the candidate                                                                                   |
| `EmailAddresses`                                                                                                     | [][components.Email](../../models/components/email.md)                                                               | :heavy_minus_sign:                                                                                                   | The email addresses of the candidate                                                                                 |
| `FieldMappings`                                                                                                      | [components.UnifiedAtsCandidateInputFieldMappings](../../models/components/unifiedatscandidateinputfieldmappings.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |