# Payments
(*Accounting.Payments*)

### Available Operations

* [List](#list) - List  Payments
* [Create](#create) - Create Payments
* [Retrieve](#retrieve) - Retrieve Payments

## List

List  Payments

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
    res, err := s.Accounting.Payments.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAccountingPaymentResponse](../../models/operations/listaccountingpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Payments in any supported Accounting software

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

    unifiedAccountingPaymentInput := components.UnifiedAccountingPaymentInput{}
    ctx := context.Background()
    res, err := s.Accounting.Payments.Create(ctx, xConnectionToken, unifiedAccountingPaymentInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingPaymentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `xConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |
| `unifiedAccountingPaymentInput`                                                                      | [components.UnifiedAccountingPaymentInput](../../models/components/unifiedaccountingpaymentinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `remoteData`                                                                                         | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Set to true to include data from the original Accounting software.                                   |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateAccountingPaymentResponse](../../models/operations/createaccountingpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Payments from any connected Accounting software

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
    res, err := s.Accounting.Payments.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingPaymentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the payment you want to retrieve.                            |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |


### Response

**[*operations.RetrieveAccountingPaymentResponse](../../models/operations/retrieveaccountingpaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
