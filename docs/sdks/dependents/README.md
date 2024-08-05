# Dependents
(*Hris.Dependents*)

### Available Operations

* [List](#list) - List Dependents
* [Retrieve](#retrieve) - Retrieve Dependent

## List

List Dependents

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
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Hris.Dependents.List(ctx, xConnectionToken, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `xConnectionToken`                                       | *string*                                                 | :heavy_check_mark:                                       | The connection token                                     |
| `remoteData`                                             | **bool*                                                  | :heavy_minus_sign:                                       | Set to true to include data from the original software.  |
| `limit`                                                  | **float64*                                               | :heavy_minus_sign:                                       | Set to get the number of records.                        |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | Set to get the number of records after this cursor.      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.ListHrisDependentsResponse](../../models/operations/listhrisdependentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve a Dependent from any connected Hris software

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
    var xConnectionToken string = "<value>"

    var id string = "801f9ede-c698-4e66-a7fc-48d19eebaa4f"

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Hris.Dependents.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisDependentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |                                                              |
| `xConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |                                                              |
| `id`                                                         | *string*                                                     | :heavy_check_mark:                                           | id of the dependent you want to retrieve.                    | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                         |
| `remoteData`                                                 | **bool*                                                      | :heavy_minus_sign:                                           | Set to true to include data from the original Hris software. | false                                                        |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |                                                              |


### Response

**[*operations.RetrieveHrisDependentResponse](../../models/operations/retrievehrisdependentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
