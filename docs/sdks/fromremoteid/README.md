# Fromremoteid
(*LinkedUsers.Fromremoteid*)

### Available Operations

* [RemoteID](#remoteid) - Retrieve a Linked User From A Remote Id

## RemoteID

Retrieve a Linked User From A Remote Id

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var remoteID string = "id_1"
    ctx := context.Background()
    res, err := s.LinkedUsers.Fromremoteid.RemoteID(ctx, remoteID)
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkedUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `remoteID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | id_1                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.RemoteIDResponse](../../models/operations/remoteidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
