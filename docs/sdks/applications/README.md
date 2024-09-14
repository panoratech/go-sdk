# Applications
(*Ats.Applications*)

## Overview

### Available Operations

* [List](#list) - List  Applications
* [Create](#create) - Create Applications
* [Retrieve](#retrieve) - Retrieve Applications

## List

List  Applications

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
    res, err := s.Ats.Applications.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListAtsApplicationResponse](../../models/operations/listatsapplicationresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Applications in any supported Ats software

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
    res, err := s.Ats.Applications.Create(ctx, "<value>", components.UnifiedAtsApplicationInput{
        AppliedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        RejectedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Offers: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            "12345678-1234-1234-1234-123456789012",
        },
        Source: gosdk.String("Source Name"),
        CreditedTo: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        CurrentStage: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        RejectReason: gosdk.String("Candidate not experienced enough"),
        CandidateID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        JobID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsApplicationOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `xConnectionToken`                                                                             | *string*                                                                                       | :heavy_check_mark:                                                                             | The connection token                                                                           |                                                                                                |
| `unifiedAtsApplicationInput`                                                                   | [components.UnifiedAtsApplicationInput](../../models/components/unifiedatsapplicationinput.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `remoteData`                                                                                   | **bool*                                                                                        | :heavy_minus_sign:                                                                             | Set to true to include data from the original Ats software.                                    | false                                                                                          |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.CreateAtsApplicationResponse](../../models/operations/createatsapplicationresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve Applications from any connected Ats software

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
    res, err := s.Ats.Applications.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsApplicationOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the application you want to retrieve.                 | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Ats software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |

### Response

**[*operations.RetrieveAtsApplicationResponse](../../models/operations/retrieveatsapplicationresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
