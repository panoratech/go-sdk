# Expenses
(*Accounting.Expenses*)

## Overview

### Available Operations

* [List](#list) - List  Expenses
* [Create](#create) - Create Expenses
* [Retrieve](#retrieve) - Retrieve Expenses

## List

List  Expenses

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

    ctx := context.Background()
    res, err := s.Accounting.Expenses.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
                for {
            // handle items
        
            res, err = res.Next()
        
            if err != nil {
                // handle error
            }
        
            if res == nil {
                break
            }
        }
        
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `xConnectionToken`                                       | *string*                                                 | :heavy_check_mark:                                       | The connection token                                     |                                                          |
| `remoteData`                                             | **bool*                                                  | :heavy_minus_sign:                                       | Set to true to include data from the original software.  | true                                                     |
| `limit`                                                  | **float64*                                               | :heavy_minus_sign:                                       | Set to get the number of records.                        | 10                                                       |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | Set to get the number of records after this cursor.      | 1b8b05bb-5273-4012-b520-8657b0b90874                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListAccountingExpenseResponse](../../models/operations/listaccountingexpenseresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Expenses in any supported Accounting software

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"github.com/panoratech/go-sdk/types"
	"github.com/panoratech/go-sdk/models/components"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Accounting.Expenses.Create(ctx, "<value>", components.UnifiedAccountingExpenseInput{
        TransactionDate: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
        TotalAmount: gosdk.Float64(10000),
        SubTotal: gosdk.Float64(9000),
        TotalTaxAmount: gosdk.Float64(1000),
        Currency: gosdk.String("USD"),
        ExchangeRate: gosdk.String("1.2"),
        Memo: gosdk.String("Business lunch with client"),
        AccountID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CompanyInfoID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        TrackingCategories: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        LineItems: []components.LineItem{
            components.LineItem{
                Name: gosdk.String("Net Income"),
                Value: gosdk.Float64(100000),
                Type: gosdk.String("Operating Activities"),
                ParentItem: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
                RemoteID: gosdk.String("report_item_1234"),
                RemoteGeneratedAt: types.MustNewTimeFromString("2024-07-01T12:00:00Z"),
                CompanyInfoID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
                CreatedAt: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
                ModifiedAt: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
            },
        },
        FieldMappings: &components.UnifiedAccountingExpenseInputFieldMappings{},
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingExpenseOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `xConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |                                                                                                      |
| `unifiedAccountingExpenseInput`                                                                      | [components.UnifiedAccountingExpenseInput](../../models/components/unifiedaccountingexpenseinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `remoteData`                                                                                         | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Set to true to include data from the original Accounting software.                                   | false                                                                                                |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreateAccountingExpenseResponse](../../models/operations/createaccountingexpenseresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve Expenses from any connected Accounting software

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

    ctx := context.Background()
    res, err := s.Accounting.Expenses.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingExpenseOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |                                                                    |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the expense you want to retrieve.                            | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                               |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. | false                                                              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.RetrieveAccountingExpenseResponse](../../models/operations/retrieveaccountingexpenseresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
