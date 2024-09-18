# Rag
(*Rag*)

## Overview

### Available Operations

* [Query](#query) - Query using RAG Search

## Query

Query across your connected data sources using RAG Search

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"github.com/panoratech/go-sdk/models/components"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Rag.Query(ctx, "<value>", components.QueryBody{
        Query: "When does Panora incorporated?",
        TopK: gosdk.Float64(3),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RagQueryOutputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `xConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |
| `queryBody`                                                  | [components.QueryBody](../../models/components/querybody.md) | :heavy_check_mark:                                           | N/A                                                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.QueryResponse](../../models/operations/queryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
