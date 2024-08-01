# Verifyevent
(*Webhooks.Verifyevent*)

### Available Operations

* [VerifyEvent](#verifyevent) - Verify payload signature of the webhook

## VerifyEvent

Verify payload signature of the webhook

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/models/components"
	"context"
	"log"
)

func main() {
    s := gosdk.New()
    request := components.SignatureVerificationDto{
        Payload: components.Payload{},
        Signature: "<value>",
        Secret: "<value>",
    }
    ctx := context.Background()
    res, err := s.Webhooks.Verifyevent.VerifyEvent(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.EventPayload != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.SignatureVerificationDto](../../models/components/signatureverificationdto.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.VerifyEventResponse](../../models/operations/verifyeventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
