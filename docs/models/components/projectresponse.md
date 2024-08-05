# ProjectResponse


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `IDProject`                                  | *string*                                     | :heavy_check_mark:                           | Unique identifier for the project            | 123e4567-e89b-12d3-a456-426614174000         |
| `Name`                                       | *string*                                     | :heavy_check_mark:                           | Name of the project                          | My Project                                   |
| `SyncMode`                                   | *string*                                     | :heavy_check_mark:                           | Synchronization mode of the project          | automatic                                    |
| `PullFrequency`                              | *float64*                                    | :heavy_check_mark:                           | Frequency of pulling data in seconds         | 3600                                         |
| `RedirectURL`                                | *string*                                     | :heavy_check_mark:                           | Redirect URL for the project                 | https://example.com/redirect                 |
| `IDUser`                                     | *string*                                     | :heavy_check_mark:                           | User ID associated with the project          | 123e4567-e89b-12d3-a456-426614174001         |
| `IDConnectorSet`                             | *string*                                     | :heavy_check_mark:                           | Connector set ID associated with the project | 123e4567-e89b-12d3-a456-426614174002         |