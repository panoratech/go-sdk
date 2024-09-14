# Journalentries
(*Accounting.Journalentries*)

## Overview

### Available Operations

* [List](#list) - List  JournalEntrys
* [Create](#create) - Create Journal Entries
* [Retrieve](#retrieve) - Retrieve Journal Entries

## List

List  JournalEntrys

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
    res, err := s.Accounting.Journalentries.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListAccountingJournalEntryResponse](../../models/operations/listaccountingjournalentryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Journal Entries in any supported Accounting software

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
    res, err := s.Accounting.Journalentries.Create(ctx, "<value>", components.UnifiedAccountingJournalentryInput{
        TransactionDate: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
        Payments: []string{
            "payment1",
            "payment2",
        },
        AppliedPayments: []string{
            "appliedPayment1",
            "appliedPayment2",
        },
        Memo: gosdk.String("Monthly expense journal entry"),
        Currency: gosdk.String("USD"),
        ExchangeRate: gosdk.String("1.2"),
        IDAccCompanyInfo: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        JournalNumber: gosdk.String("JE-001"),
        TrackingCategories: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        IDAccAccountingPeriod: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        PostingStatus: gosdk.String("Posted"),
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
        FieldMappings: &components.UnifiedAccountingJournalentryInputFieldMappings{},
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingJournalentryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `xConnectionToken`                                                                                             | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The connection token                                                                                           |                                                                                                                |
| `unifiedAccountingJournalentryInput`                                                                           | [components.UnifiedAccountingJournalentryInput](../../models/components/unifiedaccountingjournalentryinput.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `remoteData`                                                                                                   | **bool*                                                                                                        | :heavy_minus_sign:                                                                                             | Set to true to include data from the original Accounting software.                                             | false                                                                                                          |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.CreateAccountingJournalEntryResponse](../../models/operations/createaccountingjournalentryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve Journal Entries from any connected Accounting software

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
    res, err := s.Accounting.Journalentries.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingJournalentryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |                                                                    |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the journalentry you want to retrieve.                       | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                               |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. | false                                                              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.RetrieveAccountingJournalEntryResponse](../../models/operations/retrieveaccountingjournalentryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
