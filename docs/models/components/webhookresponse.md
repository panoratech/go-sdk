# WebhookResponse


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `IDWebhookEndpoint`                       | *string*                                  | :heavy_check_mark:                        | The unique UUID of the webhook.           |
| `EndpointDescription`                     | *string*                                  | :heavy_check_mark:                        | The description of the webhook.           |
| `URL`                                     | *string*                                  | :heavy_check_mark:                        | The endpoint url of the webhook.          |
| `Secret`                                  | *string*                                  | :heavy_check_mark:                        | The secret of the webhook.                |
| `Active`                                  | *bool*                                    | :heavy_check_mark:                        | The status of the webhook.                |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The created date  of the webhook.         |
| `Scope`                                   | []*string*                                | :heavy_check_mark:                        | The events that the webhook listen to.    |
| `IDProject`                               | *string*                                  | :heavy_check_mark:                        | The project id tied to the webhook.       |
| `LastUpdate`                              | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The last update date of the webhook.      |