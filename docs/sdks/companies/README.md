# Companies
(*Crm.Companies*)

### Available Operations

* [List](#list) - List  Companies
* [Create](#create) - Create Companies
* [Retrieve](#retrieve) - Retrieve Companies

## List

List  Companies

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
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Crm.Companies.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListCrmCompanyResponse](../../models/operations/listcrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Companies in any supported CRM software

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
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"

    unifiedCrmCompanyInput := components.UnifiedCrmCompanyInput{
        Name: "<value>",
        FieldMappings: components.UnifiedCrmCompanyInputFieldMappings{},
    }
    ctx := context.Background()
    res, err := s.Crm.Companies.Create(ctx, xConnectionToken, unifiedCrmCompanyInput, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmCompanyOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `xConnectionToken`                                                                     | *string*                                                                               | :heavy_check_mark:                                                                     | The connection token                                                                   |
| `unifiedCrmCompanyInput`                                                               | [components.UnifiedCrmCompanyInput](../../models/components/unifiedcrmcompanyinput.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `remoteData`                                                                           | **bool*                                                                                | :heavy_minus_sign:                                                                     | Set to true to include data from the original CRM software.                            |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.CreateCrmCompanyResponse](../../models/operations/createcrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Companies from any connected Crm software

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
        gosdk.WithSecurity(os.Getenv("BEARER")),
    )
    var xConnectionToken string = "<value>"

    var id string = "<value>"
    ctx := context.Background()
    res, err := s.Crm.Companies.Retrieve(ctx, xConnectionToken, id, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmCompanyOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the company you want to retrieve.                     |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Crm software. |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |


### Response

**[*operations.RetrieveCrmCompanyResponse](../../models/operations/retrievecrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
