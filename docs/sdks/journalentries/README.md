# Journalentries
(*Accounting.Journalentries*)

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
	"os"
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(os.Getenv("API_KEY")),
    )
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Accounting.Journalentries.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAccountingJournalEntryResponse](../../models/operations/listaccountingjournalentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Journal Entries in any supported Accounting software

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
        gosdk.WithSecurity(os.Getenv("API_KEY")),
    )
    var xConnectionToken string = "<value>"

    unifiedAccountingJournalentryInput := components.UnifiedAccountingJournalentryInput{}
    ctx := context.Background()
    res, err := s.Accounting.Journalentries.Create(ctx, xConnectionToken, unifiedAccountingJournalentryInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingJournalentryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `xConnectionToken`                                                                                             | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The connection token                                                                                           |
| `unifiedAccountingJournalentryInput`                                                                           | [components.UnifiedAccountingJournalentryInput](../../models/components/unifiedaccountingjournalentryinput.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `remoteData`                                                                                                   | **bool*                                                                                                        | :heavy_minus_sign:                                                                                             | Set to true to include data from the original Accounting software.                                             |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.CreateAccountingJournalEntryResponse](../../models/operations/createaccountingjournalentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Journal Entries from any connected Accounting software

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
        gosdk.WithSecurity(os.Getenv("API_KEY")),
    )
    var xConnectionToken string = "<value>"

    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Accounting.Journalentries.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAccountingJournalentryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `xConnectionToken`                                                 | *string*                                                           | :heavy_check_mark:                                                 | The connection token                                               |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | id of the journalentry you want to retrieve.                       |
| `remoteData`                                                       | **bool*                                                            | :heavy_minus_sign:                                                 | Set to true to include data from the original Accounting software. |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |


### Response

**[*operations.RetrieveAccountingJournalEntryResponse](../../models/operations/retrieveaccountingjournalentryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
