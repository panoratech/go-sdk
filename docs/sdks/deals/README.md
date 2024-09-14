# Deals
(*Crm.Deals*)

## Overview

### Available Operations

* [List](#list) - List Deals
* [Create](#create) - Create Deals
* [Retrieve](#retrieve) - Retrieve Deals

## List

List Deals

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
    res, err := s.Crm.Deals.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListCrmDealsResponse](../../models/operations/listcrmdealsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Deals in any supported Crm software

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
    res, err := s.Crm.Deals.Create(ctx, "<value>", components.UnifiedCrmDealInput{
        Name: gosdk.String("Huge Contract with Acme"),
        Description: gosdk.String("Contract with Sales Operations Team"),
        Amount: gosdk.Float64(1000),
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        StageID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CompanyID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmDealOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `xConnectionToken`                                                               | *string*                                                                         | :heavy_check_mark:                                                               | The connection token                                                             |
| `unifiedCrmDealInput`                                                            | [components.UnifiedCrmDealInput](../../models/components/unifiedcrmdealinput.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `remoteData`                                                                     | **bool*                                                                          | :heavy_minus_sign:                                                               | Set to true to include data from the original Crm software.                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateCrmDealResponse](../../models/operations/createcrmdealresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve Deals from any connected Crm software

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
    res, err := s.Crm.Deals.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmDealOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the deal you want to retrieve.                        | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Crm software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |

### Response

**[*operations.RetrieveCrmDealResponse](../../models/operations/retrievecrmdealresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
