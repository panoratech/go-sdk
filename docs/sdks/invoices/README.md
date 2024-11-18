# Invoices
(*Accounting.Invoices*)

## Overview

### Available Operations

* [List](#list) - List  Invoices
* [Create](#create) - Create Invoices
* [Retrieve](#retrieve) - Retrieve Invoices

## List

List  Invoices

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
    res, err := s.Accounting.Invoices.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListAccountingInvoiceResponse](../../models/operations/listaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create invoices in any supported Accounting software

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
    res, err := s.Accounting.Invoices.Create(ctx, "<value>", components.UnifiedAccountingInvoiceInput{
        Type: gosdk.String("Sales"),
        Number: gosdk.String("INV-001"),
        IssueDate: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
        DueDate: types.MustNewTimeFromString("2024-07-15T12:00:00Z"),
        PaidOnDate: types.MustNewTimeFromString("2024-07-10T12:00:00Z"),
        Memo: gosdk.String("Payment for services rendered"),
        Currency: gosdk.String("USD"),
        ExchangeRate: gosdk.String("1.2"),
        TotalDiscount: gosdk.Float64(1000),
        SubTotal: gosdk.Float64(10000),
        Status: gosdk.String("Paid"),
        TotalTaxAmount: gosdk.Float64(1000),
        TotalAmount: gosdk.Float64(11000),
        Balance: gosdk.Float64(0),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        AccountingPeriodID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        TrackingCategories: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
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
        FieldMappings: &components.UnifiedAccountingInvoiceInputFieldMappings{},
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingInvoiceOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `xConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |                                                                                                      |
| `unifiedAccountingInvoiceInput`                                                                      | [components.UnifiedAccountingInvoiceInput](../../models/components/unifiedaccountinginvoiceinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `remoteData`                                                                                         | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Set to true to include data from the original Accounting software.                                   | false                                                                                                |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CreateAccountingInvoiceResponse](../../models/operations/createaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieve Invoices from any connected Accounting software

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
    res, err := s.Accounting.Invoices.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingInvoiceOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |                                                                    |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the invoice you want to retrieve.                            | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                               |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. | false                                                              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.RetrieveAccountingInvoiceResponse](../../models/operations/retrieveaccountinginvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |