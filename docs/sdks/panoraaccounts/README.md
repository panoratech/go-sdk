# PanoraAccounts
(*Accounting.Accounts*)

### Available Operations

* [List](#list) - List  Accounts
* [Create](#create) - Create Accounts
* [Retrieve](#retrieve) - Retrieve Accounts

## List

List  Accounts

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
    res, err := s.Accounting.Accounts.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAccountingAccountsResponse](../../models/operations/listaccountingaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create accounts in any supported Accounting software

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

    unifiedAccountingAccountInput := components.UnifiedAccountingAccountInput{}
    ctx := context.Background()
    res, err := s.Accounting.Accounts.Create(ctx, xConnectionToken, unifiedAccountingAccountInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingAccountOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `xConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |
| `unifiedAccountingAccountInput`                                                                      | [components.UnifiedAccountingAccountInput](../../models/components/unifiedaccountingaccountinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `remoteData`                                                                                         | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Set to true to include data from the original Accounting software.                                   |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateAccountingAccountResponse](../../models/operations/createaccountingaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Accounts from any connected Accounting software

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
    res, err := s.Accounting.Accounts.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingAccountOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the account you want to retrieve.                            |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |


### Response

**[*operations.RetrieveAccountingAccountResponse](../../models/operations/retrieveaccountingaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
