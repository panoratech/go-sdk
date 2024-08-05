# Map
(*FieldMappings.Map*)

### Available Operations

* [Map](#map) - Map Custom Field

## Map

Map Custom Field

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
    request := components.MapFieldToProviderDto{
        AttributeID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        SourceCustomFieldID: gosdk.String("id_1"),
        SourceProvider: gosdk.String("hubspot"),
        LinkedUserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
    }
    ctx := context.Background()
    res, err := s.FieldMappings.Map.Map(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomFieldResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.MapFieldToProviderDto](../../models/components/mapfieldtoproviderdto.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.MapResponse](../../models/operations/mapresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
