# Tickets
(*Ticketing.Tickets*)

### Available Operations

* [List](#list) - List  Tickets
* [Create](#create) - Create Tickets
* [Retrieve](#retrieve) - Retrieve Tickets

## List

List  Tickets

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

    var remoteData *bool = gosdk.Bool(false)

    var limit *float64 = gosdk.Float64(50)

    var cursor *string = gosdk.String("<value>")
    ctx := context.Background()
    res, err := s.Ticketing.Tickets.List(ctx, xConnectionToken, remoteData, limit, cursor)
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

**[*operations.ListTicketingTicketResponse](../../models/operations/listticketingticketresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Tickets in any supported Ticketing software

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

    unifiedTicketingTicketInput := components.UnifiedTicketingTicketInput{
        Name: "<value>",
        Description: "Multi-tiered human-resource model",
        FieldMappings: components.UnifiedTicketingTicketInputFieldMappings{},
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ticketing.Tickets.Create(ctx, xConnectionToken, unifiedTicketingTicketInput, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingTicketOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `xConnectionToken`                                                                               | *string*                                                                                         | :heavy_check_mark:                                                                               | The connection token                                                                             |
| `unifiedTicketingTicketInput`                                                                    | [components.UnifiedTicketingTicketInput](../../models/components/unifiedticketingticketinput.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `remoteData`                                                                                     | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Set to true to include data from the original Ticketing software.                                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreateTicketingTicketResponse](../../models/operations/createticketingticketresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Tickets from any connected Ticketing software

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

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ticketing.Tickets.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingTicketOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `xConnectionToken`                                                | *string*                                                          | :heavy_check_mark:                                                | The connection token                                              |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | id of the `ticket` you want to retrive.                           |
| `remoteData`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | Set to true to include data from the original Ticketing software. |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |


### Response

**[*operations.RetrieveTicketingTicketResponse](../../models/operations/retrieveticketingticketresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |