# Tickets
(*Ticketing.Tickets*)

## Overview

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

    ctx := context.Background()
    res, err := s.Ticketing.Tickets.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListTicketingTicketResponse](../../models/operations/listticketingticketresponse.md), error**

### Errors

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
    res, err := s.Ticketing.Tickets.Create(ctx, "<value>", components.UnifiedTicketingTicketInput{
        Name: gosdk.String("Customer Service Inquiry"),
        Status: gosdk.String("OPEN"),
        Description: gosdk.String("Help customer"),
        DueDate: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Type: gosdk.String("BUG"),
        ParentTicket: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Collections: []components.UnifiedTicketingTicketInputCollections{
            components.CreateUnifiedTicketingTicketInputCollectionsStr(
                "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            ),
        },
        Tags: []components.UnifiedTicketingTicketInputTags{
            components.CreateUnifiedTicketingTicketInputTagsStr(
                "my_tag",
            ),
            components.CreateUnifiedTicketingTicketInputTagsStr(
                "urgent_tag",
            ),
        },
        CompletedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Priority: gosdk.String("HIGH"),
        AssignedTo: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Comment: &components.UnifiedTicketingTicketInputComment{
            Body: gosdk.String("Assigned to Eric !"),
            HTMLBody: gosdk.String("<p>Assigned to Eric !</p>"),
            IsPrivate: gosdk.Bool(false),
            CreatorType: gosdk.String("USER"),
            TicketID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
            Attachments: []components.UnifiedTicketingTicketInputCommentAttachments{
                components.CreateUnifiedTicketingTicketInputCommentAttachmentsStr(
                    "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
                ),
            },
        },
        AccountID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Attachments: []components.UnifiedTicketingTicketInputAttachments{
            components.CreateUnifiedTicketingTicketInputAttachmentsStr(
                "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            ),
        },
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }, gosdk.Bool(false))
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

### Errors

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

    ctx := context.Background()
    res, err := s.Ticketing.Tickets.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
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

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
