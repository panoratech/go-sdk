# Files
(*Filestorage.Files*)

### Available Operations

* [List](#list) - List  Files
* [Create](#create) - Create Files
* [Retrieve](#retrieve) - Retrieve Files

## List

List  Files

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
    res, err := s.Filestorage.Files.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListFilestorageFileResponse](../../models/operations/listfilestoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Files in any supported Filestorage software

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

    var remoteData bool = false

    unifiedFilestorageFileInput := components.UnifiedFilestorageFileInput{
        Name: gosdk.String("my_paris_photo.png"),
        FileURL: gosdk.String("https://example.com/my_paris_photo.png"),
        MimeType: gosdk.String("application/pdf"),
        Size: gosdk.String("1024"),
        FolderID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Permission: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        SharedLink: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }
    ctx := context.Background()
    res, err := s.Filestorage.Files.Create(ctx, xConnectionToken, remoteData, unifiedFilestorageFileInput)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedFilestorageFileOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `xConnectionToken`                                                                               | *string*                                                                                         | :heavy_check_mark:                                                                               | The connection token                                                                             |
| `remoteData`                                                                                     | *bool*                                                                                           | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `unifiedFilestorageFileInput`                                                                    | [components.UnifiedFilestorageFileInput](../../models/components/unifiedfilestoragefileinput.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |


### Response

**[*operations.CreateFilestorageFileResponse](../../models/operations/createfilestoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Files from any connected Filestorage software

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
    res, err := s.Filestorage.Files.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedFilestorageFileOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `xConnectionToken`                                                   | *string*                                                             | :heavy_check_mark:                                                   | The connection token                                                 |                                                                      |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | id of the file you want to retrieve.                                 | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                                 |
| `remoteData`                                                         | **bool*                                                              | :heavy_minus_sign:                                                   | Set to true to include data from the original File Storage software. | false                                                                |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |


### Response

**[*operations.RetrieveFilestorageFileResponse](../../models/operations/retrievefilestoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
