# PanoraAttachments
(*Ticketing.Attachments*)

## Overview

### Available Operations

* [List](#list) - List  Attachments
* [Create](#create) - Create Attachments
* [Retrieve](#retrieve) - Retrieve Attachments

## List

List  Attachments

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
    res, err := s.Ticketing.Attachments.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListTicketingAttachmentsResponse](../../models/operations/listticketingattachmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create Attachments in any supported Ticketing software

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
    res, err := s.Ticketing.Attachments.Create(ctx, "<value>", components.UnifiedTicketingAttachmentInput{
        FileName: gosdk.String("features_planning.pdf"),
        FileURL: gosdk.String("https://example.com/features_planning.pdf"),
        Uploader: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        TicketID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CommentID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingAttachmentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `xConnectionToken`                                                                                       | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The connection token                                                                                     |
| `unifiedTicketingAttachmentInput`                                                                        | [components.UnifiedTicketingAttachmentInput](../../models/components/unifiedticketingattachmentinput.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `remoteData`                                                                                             | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | Set to true to include data from the original Ticketing software.                                        |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateTicketingAttachmentResponse](../../models/operations/createticketingattachmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieve Attachments from any connected Ticketing software

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
    res, err := s.Ticketing.Attachments.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingAttachmentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `xConnectionToken`                                                | *string*                                                          | :heavy_check_mark:                                                | The connection token                                              |                                                                   |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | id of the attachment you want to retrive.                         | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                              |
| `remoteData`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | Set to true to include data from the original Ticketing software. | false                                                             |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |

### Response

**[*operations.RetrieveTicketingAttachmentResponse](../../models/operations/retrieveticketingattachmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |