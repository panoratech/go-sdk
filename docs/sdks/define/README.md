# Define
(*FieldMappings.Define*)

### Available Operations

* [Definitions](#definitions) - Define target Field

## Definitions

Define target Field

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
    request := components.DefineTargetFieldDto{
        ObjectTypeOwner: gosdk.String("<value>"),
        Name: gosdk.String("<value>"),
        Description: gosdk.String("Universal heuristic matrices"),
        DataType: gosdk.String("decimal"),
    }
    ctx := context.Background()
    res, err := s.FieldMappings.Define.Definitions(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFieldResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.DefineTargetFieldDto](../../models/components/definetargetfielddto.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |


### Response

**[*operations.DefinitionsResponse](../../models/operations/definitionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
