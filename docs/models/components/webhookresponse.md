# WebhookResponse


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `IDWebhookEndpoint`                       | *string*                                  | :heavy_check_mark:                        | The unique UUID of the webhook.           | 801f9ede-c698-4e66-a7fc-48d19eebaa4f      |
| `EndpointDescription`                     | *string*                                  | :heavy_check_mark:                        | The description of the webhook.           | Webhook to receive connection events      |
| `URL`                                     | *string*                                  | :heavy_check_mark:                        | The endpoint url of the webhook.          | https://acme.com/webhook_receiver         |
| `Secret`                                  | *string*                                  | :heavy_check_mark:                        | The secret of the webhook.                | 801f9ede-c698-4e66-a7fc-48d19eebaa4f      |
| `Active`                                  | *bool*                                    | :heavy_check_mark:                        | The status of the webhook.                | true                                      |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The created date of the webhook.          | 2024-10-01T12:00:00Z                      |
| `Scope`                                   | []*string*                                | :heavy_check_mark:                        | The events that the webhook listen to.    | [<br/>"connection.created"<br/>]          |
| `IDProject`                               | *string*                                  | :heavy_check_mark:                        | The project id tied to the webhook.       | 801f9ede-c698-4e66-a7fc-48d19eebaa4f      |
| `LastUpdate`                              | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The last update date of the webhook.      | 2024-10-01T12:00:00Z                      |