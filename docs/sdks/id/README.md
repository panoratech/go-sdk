# ID
(*LinkedUsers.{id}*)

### Available Operations

* [Retrieve](#retrieve) - Retrieve Linked Users

## Retrieve

Retrieve Linked Users

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
    var id string = "801f9ede-c698-4e66-a7fc-48d19eebaa4f"
    ctx := context.Background()
    res, err := s.LinkedUsers.{id}.Retrieve(ctx, id)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |


### Response

**[*operations.RetrieveLinkedUserResponse](../../models/operations/retrievelinkeduserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
