# Folders
(*Filestorage.Folders*)

### Available Operations

* [List](#list) - List  Folders
* [Create](#create) - Create Folders
* [Retrieve](#retrieve) - Retrieve Folders

## List

List  Folders

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
    res, err := s.Filestorage.Folders.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListFilestorageFolderResponse](../../models/operations/listfilestoragefolderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Folders in any supported Filestorage software

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

    var remoteData bool = false

    unifiedFilestorageFolderInput := components.UnifiedFilestorageFolderInput{
        Name: gosdk.String("<value>"),
        Size: gosdk.String("<value>"),
        FolderURL: gosdk.String("<value>"),
        Description: "Multi-tiered human-resource model",
        DriveID: gosdk.String("<value>"),
        ParentFolderID: gosdk.String("<value>"),
        SharedLink: gosdk.String("<value>"),
        Permission: gosdk.String("<value>"),
    }
    ctx := context.Background()
    res, err := s.Filestorage.Folders.Create(ctx, xConnectionToken, remoteData, unifiedFilestorageFolderInput)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedFilestorageFolderOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `xConnectionToken`                                                                                   | *string*                                                                                             | :heavy_check_mark:                                                                                   | The connection token                                                                                 |
| `remoteData`                                                                                         | *bool*                                                                                               | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `unifiedFilestorageFolderInput`                                                                      | [components.UnifiedFilestorageFolderInput](../../models/components/unifiedfilestoragefolderinput.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.CreateFilestorageFolderResponse](../../models/operations/createfilestoragefolderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Folders from any connected Filestorage software

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
    res, err := s.Filestorage.Folders.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedFilestorageFolderOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `xConnectionToken`                                                   | *string*                                                             | :heavy_check_mark:                                                   | The connection token                                                 |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | id of the folder you want to retrieve.                               |
| `remoteData`                                                         | **bool*                                                              | :heavy_minus_sign:                                                   | Set to true to include data from the original File Storage software. |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |


### Response

**[*operations.RetrieveFilestorageFolderResponse](../../models/operations/retrievefilestoragefolderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
