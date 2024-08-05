# Tasks
(*Crm.Tasks*)

### Available Operations

* [List](#list) - List Tasks
* [Create](#create) - Create Tasks
* [Retrieve](#retrieve) - Retrieve Tasks

## List

List Tasks

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
    res, err := s.Crm.Tasks.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListCrmTaskResponse](../../models/operations/listcrmtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Tasks in any supported Crm software

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

    unifiedCrmTaskInput := components.UnifiedCrmTaskInput{
        Subject: gosdk.String("Answer customers"),
        Content: gosdk.String("Prepare email campaign"),
        Status: components.UnifiedCrmTaskInputStatusPending.ToPointer(),
        DueDate: gosdk.String("2024-10-01T12:00:00Z"),
        FinishedDate: gosdk.String("2024-10-01T12:00:00Z"),
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CompanyID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        DealID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }
    ctx := context.Background()
    res, err := s.Crm.Tasks.Create(ctx, xConnectionToken, unifiedCrmTaskInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmTaskOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `xConnectionToken`                                                               | *string*                                                                         | :heavy_check_mark:                                                               | The connection token                                                             |
| `unifiedCrmTaskInput`                                                            | [components.UnifiedCrmTaskInput](../../models/components/unifiedcrmtaskinput.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `remoteData`                                                                     | **bool*                                                                          | :heavy_minus_sign:                                                               | Set to true to include data from the original Crm software.                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.CreateCrmTaskResponse](../../models/operations/createcrmtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Tasks from any connected Crm software

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
    res, err := s.Crm.Tasks.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmTaskOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the task you want to retrieve.                        | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Crm software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |


### Response

**[*operations.RetrieveCrmTaskResponse](../../models/operations/retrievecrmtaskresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
