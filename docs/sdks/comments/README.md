# Comments
(*Ticketing.Comments*)

## Overview

### Available Operations

* [List](#list) - List Comments
* [Create](#create) - Create Comments
* [Retrieve](#retrieve) - Retrieve Comment

## List

List Comments

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
    res, err := s.Ticketing.Comments.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListTicketingCommentsResponse](../../models/operations/listticketingcommentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create Comments in any supported Ticketing software

### Example Usage

```go
package main

import(
	gosdk "github.com/panoratech/go-sdk"
	"context"
	"github.com/panoratech/go-sdk/models/components"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Ticketing.Comments.Create(ctx, "<value>", components.UnifiedTicketingCommentInput{
        Body: gosdk.String("Assigned to Eric !"),
        HTMLBody: gosdk.String("<p>Assigned to Eric !</p>"),
        IsPrivate: gosdk.Bool(false),
        CreatorType: gosdk.String("USER"),
        TicketID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        ContactID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Attachments: []components.UnifiedTicketingCommentInputAttachments{
            components.CreateUnifiedTicketingCommentInputAttachmentsStr(
                "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            ),
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingCommentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `xConnectionToken`                                                                                 | *string*                                                                                           | :heavy_check_mark:                                                                                 | The connection token                                                                               |
| `unifiedTicketingCommentInput`                                                                     | [components.UnifiedTicketingCommentInput](../../models/components/unifiedticketingcommentinput.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `remoteData`                                                                                       | **bool*                                                                                            | :heavy_minus_sign:                                                                                 | Set to true to include data from the original Ticketing software.                                  |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateTicketingCommentResponse](../../models/operations/createticketingcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieve a Comment from any connected Ticketing software

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
    res, err := s.Ticketing.Comments.Retrieve(ctx, "<value>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `xConnectionToken`                                                | *string*                                                          | :heavy_check_mark:                                                | The connection token                                              |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | id of the `comment` you want to retrive.                          |
| `remoteData`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | Set to true to include data from the original Ticketing software. |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.RetrieveTicketingCommentResponse](../../models/operations/retrieveticketingcommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |