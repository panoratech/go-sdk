# Timeoffs
(*Hris.Timeoffs*)

### Available Operations

* [List](#list) - List  Timeoffs
* [Create](#create) - Create Timeoffs
* [Retrieve](#retrieve) - Retrieve Timeoffs

## List

List  Timeoffs

### Example Usage

```go
package main

import(
	"os"
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Hris.Timeoffs.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListHrisTimeoffsResponse](../../models/operations/listhristimeoffsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Timeoffs in any supported Hris software

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
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"

    unifiedHrisTimeoffInput := components.UnifiedHrisTimeoffInput{}
    ctx := context.Background()
    res, err := s.Hris.Timeoffs.Create(ctx, xConnectionToken, unifiedHrisTimeoffInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisTimeoffOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `xConnectionToken`                                                                       | *string*                                                                                 | :heavy_check_mark:                                                                       | The connection token                                                                     |
| `unifiedHrisTimeoffInput`                                                                | [components.UnifiedHrisTimeoffInput](../../models/components/unifiedhristimeoffinput.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `remoteData`                                                                             | **bool*                                                                                  | :heavy_minus_sign:                                                                       | Set to true to include data from the original Hris software.                             |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.CreateHrisTimeoffResponse](../../models/operations/createhristimeoffresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Timeoffs from any connected Hris software

### Example Usage

```go
package main

import(
	"os"
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"

    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Hris.Timeoffs.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisTimeoffOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `xConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |
| `id`                                                         | *string*                                                     | :heavy_check_mark:                                           | id of the timeoff you want to retrieve.                      |
| `remoteData`                                                 | **bool*                                                      | :heavy_minus_sign:                                           | Set to true to include data from the original Hris software. |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |


### Response

**[*operations.RetrieveHrisTimeoffResponse](../../models/operations/retrievehristimeoffresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
