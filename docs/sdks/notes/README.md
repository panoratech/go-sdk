# Notes
(*Crm.Notes*)

### Available Operations

* [List](#list) - List Notes
* [Create](#create) - Create Notes
* [Retrieve](#retrieve) - Retrieve Notes

## List

List Notes

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
    res, err := s.Crm.Notes.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListCrmNoteResponse](../../models/operations/listcrmnoteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Notes in any supported Crm software

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/models/components"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var xConnectionToken string = "<value>"

    unifiedCrmNoteInput := components.UnifiedCrmNoteInput{
        Content: gosdk.String("My notes taken during the meeting"),
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CompanyID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        DealID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Crm.Notes.Create(ctx, xConnectionToken, unifiedCrmNoteInput, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmNoteOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `xConnectionToken`                                                               | *string*                                                                         | :heavy_check_mark:                                                               | The connection token                                                             |                                                                                  |
| `unifiedCrmNoteInput`                                                            | [components.UnifiedCrmNoteInput](../../models/components/unifiedcrmnoteinput.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `remoteData`                                                                     | **bool*                                                                          | :heavy_minus_sign:                                                               | Set to true to include data from the original Crm software.                      | false                                                                            |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |


### Response

**[*operations.CreateCrmNoteResponse](../../models/operations/createcrmnoteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Notes from any connected Crm software

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
    res, err := s.Crm.Notes.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmNoteOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the note you want to retrieve.                        | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Crm software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |


### Response

**[*operations.RetrieveCrmNoteResponse](../../models/operations/retrievecrmnoteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
