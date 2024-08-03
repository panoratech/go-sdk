# Batch
(*LinkedUsers.Batch*)

### Available Operations

* [ImportBatch](#importbatch) - Add Batch Linked Users

## ImportBatch

Add Batch Linked Users

### Example Usage

```go
package main

import(
	"os"
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/models/components"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(os.Getenv("API_KEY")),
    )
    request := components.CreateBatchLinkedUserDto{
        LinkedUserOriginIds: []string{
            "<value>",
        },
        Alias: gosdk.String("<value>"),
    }
    ctx := context.Background()
    res, err := s.LinkedUsers.Batch.ImportBatch(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkedUserResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateBatchLinkedUserDto](../../models/components/createbatchlinkeduserdto.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ImportBatchResponse](../../models/operations/importbatchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
