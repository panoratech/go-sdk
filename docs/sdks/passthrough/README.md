# Passthrough
(*Passthrough*)

### Available Operations

* [Request](#request) - Make a passthrough request

## Request

Make a passthrough request

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
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var xConnectionToken string = "<value>"

    passThroughRequestDto := components.PassThroughRequestDto{
        Method: components.PassThroughRequestDtoMethodGet,
        Path: gosdk.String("/dev"),
    }
    ctx := context.Background()
    res, err := s.Passthrough.Request(ctx, xConnectionToken, passThroughRequestDto)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `xConnectionToken`                                                                   | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `passThroughRequestDto`                                                              | [components.PassThroughRequestDto](../../models/components/passthroughrequestdto.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.RequestResponse](../../models/operations/requestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
