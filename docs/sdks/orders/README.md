# Orders
(*Ecommerce.Orders*)

## Overview

### Available Operations

* [List](#list) - List Orders
* [Create](#create) - Create Orders
* [Retrieve](#retrieve) - Retrieve Orders

## List

List Orders

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
    res, err := s.Ecommerce.Orders.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListEcommerceOrdersResponse](../../models/operations/listecommerceordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create Orders in any supported Ecommerce software

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
    res, err := s.Ecommerce.Orders.Create(ctx, "<value>", components.UnifiedEcommerceOrderInput{
        OrderStatus: gosdk.String("UNSHIPPED"),
        OrderNumber: gosdk.String("19823838833"),
        PaymentStatus: gosdk.String("SUCCESS"),
        Currency: gosdk.String("AUD"),
        TotalPrice: gosdk.Float64(300),
        TotalDiscount: gosdk.Float64(10),
        TotalShipping: gosdk.Float64(120),
        TotalTax: gosdk.Float64(120),
        FulfillmentStatus: gosdk.String("PENDING"),
        CustomerID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Items: []components.LineItem{
            components.LineItem{
                Name: gosdk.String("Net Income"),
                Value: gosdk.Float64(100000),
                Type: gosdk.String("Operating Activities"),
                ParentItem: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
                RemoteID: gosdk.String("12345"),
                RemoteGeneratedAt: types.MustNewTimeFromString("2024-07-01T12:00:00Z"),
                CompanyInfoID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
                CreatedAt: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
                ModifiedAt: types.MustNewTimeFromString("2024-06-15T12:00:00Z"),
            },
        },
        FieldMappings: &components.UnifiedEcommerceOrderInputFieldMappings{},
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedEcommerceOrderOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `xConnectionToken`                                                                             | *string*                                                                                       | :heavy_check_mark:                                                                             | The connection token                                                                           |                                                                                                |
| `unifiedEcommerceOrderInput`                                                                   | [components.UnifiedEcommerceOrderInput](../../models/components/unifiedecommerceorderinput.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `remoteData`                                                                                   | **bool*                                                                                        | :heavy_minus_sign:                                                                             | Set to true to include data from the original Accounting software.                             | false                                                                                          |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.CreateEcommerceOrderResponse](../../models/operations/createecommerceorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Retrieve

Retrieve orders from any connected Ats software

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
    res, err := s.Ecommerce.Orders.Retrieve(ctx, "<value>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedEcommerceOrderOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the order you want to retrieve.                       |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Ats software. |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*operations.RetrieveEcommerceOrderResponse](../../models/operations/retrieveecommerceorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |