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
    res, err := s.Ticketing.Tickets.List(ctx, xConnectionToken, nil, nil, nil)
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
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/models/components"
	"github.com/panoratech/go-sdk/types"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var xConnectionToken string = "<value>"

    unifiedTicketingTicketInput := components.UnifiedTicketingTicketInput{
        Name: gosdk.String("Customer Service Inquiry"),
        Status: components.UnifiedTicketingTicketInputStatusOpen.ToPointer(),
        Description: gosdk.String("Help customer"),
        DueDate: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Type: components.UnifiedTicketingTicketInputTypeBug.ToPointer(),
        ParentTicket: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Collections: gosdk.String("[\"801f9ede-c698-4e66-a7fc-48d19eebaa4f\"]"),
        Tags: []string{
            "my_tag",
            "urgent_tag",
        },
        CompletedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Priority: components.UnifiedTicketingTicketInputPriorityHigh.ToPointer(),
        AssignedTo: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Comment: &components.UnifiedTicketingTicketInputComment{
            Body: gosdk.String("Assigned to Eric !"),
            HTMLBody: gosdk.String("<p>Assigned to Eric !</p>"),
            IsPrivate: gosdk.Bool(false),
            CreatorType: components.UnifiedTicketingTicketInputCreatorTypeUser.ToPointer(),
            TicketID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            Attachments: []string{
                "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            },
        },
        AccountID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Attachments: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `xConnectionToken`                                                                               | *string*                                                                                         | :heavy_check_mark:                                                                               | The connection token                                                                             |                                                                                                  |
| `unifiedTicketingTicketInput`                                                                    | [components.UnifiedTicketingTicketInput](../../models/components/unifiedticketingticketinput.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `remoteData`                                                                                     | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Set to true to include data from the original Ticketing software.                                | false                                                                                            |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |


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

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `xConnectionToken`                                                | *string*                                                          | :heavy_check_mark:                                                | The connection token                                              |                                                                   |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | id of the `ticket` you want to retrive.                           | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                              |
| `remoteData`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | Set to true to include data from the original Ticketing software. | false                                                             |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |


### Response

**[*operations.RetrieveTicketingTicketResponse](../../models/operations/retrieveticketingticketresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
