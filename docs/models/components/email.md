# Email


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `EmailAddress`                                                         | *string*                                                               | :heavy_check_mark:                                                     | The email address                                                      |
| `EmailAddressType`                                                     | *string*                                                               | :heavy_check_mark:                                                     | The email address type. Authorized values are either PERSONAL or WORK. |
| `OwnerType`                                                            | [*components.OwnerType](../../models/components/ownertype.md)          | :heavy_minus_sign:                                                     | The owner type of an email                                             |