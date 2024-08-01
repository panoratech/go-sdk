# PanoraTicketingAttachments
(*Ticketing.Attachments*)

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
    s := gosdk.New()
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Ticketing.Attachments.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListTicketingAttachmentsResponse](../../models/operations/listticketingattachmentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Attachments in any supported Ticketing software

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
    s := gosdk.New()
    var xConnectionToken string = "<value>"

    unifiedTicketingAttachmentInput := components.UnifiedTicketingAttachmentInput{
        FileName: "your_file_here",
        FileURL: "<value>",
        Uploader: "<value>",
        FieldMappings: components.UnifiedTicketingAttachmentInputFieldMappings{},
    }
    ctx := context.Background()
    res, err := s.Ticketing.Attachments.Create(ctx, xConnectionToken, unifiedTicketingAttachmentInput, nil)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
    s := gosdk.New()
    var xConnectionToken string = "<value>"

    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Ticketing.Attachments.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedTicketingAttachmentOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `xConnectionToken`                                                | *string*                                                          | :heavy_check_mark:                                                | The connection token                                              |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | id of the attachment you want to retrive.                         |
| `remoteData`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | Set to true to include data from the original Ticketing software. |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |


### Response

**[*operations.RetrieveTicketingAttachmentResponse](../../models/operations/retrieveticketingattachmentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
