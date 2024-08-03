# SignatureVerificationDto


## Fields

| Field                             | Type                              | Required                          | Description                       |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `Payload`                         | map[string]*any*                  | :heavy_check_mark:                | The payload event of the webhook. |
| `Signature`                       | *string*                          | :heavy_check_mark:                | The signature of the webhook.     |
| `Secret`                          | *string*                          | :heavy_check_mark:                | The secret of the webhook.        |