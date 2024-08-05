# Applications
(*Ats.Applications*)

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
    var xConnectionToken string = "<value>"
    ctx := context.Background()
    res, err := s.Ats.Applications.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAtsApplicationResponse](../../models/operations/listatsapplicationresponse.md), error**
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
	"github.com/panoratech/go-sdk/models/components"
	"github.com/panoratech/go-sdk/types"
	"context"
	"log"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )
    var xConnectionToken string = "<value>"

    unifiedAtsApplicationInput := components.UnifiedAtsApplicationInput{
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
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ats.Applications.Create(ctx, xConnectionToken, unifiedAtsApplicationInput, remoteData)
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
    var xConnectionToken string = "<value>"

    var id string = "801f9ede-c698-4e66-a7fc-48d19eebaa4f"

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ats.Applications.Retrieve(ctx, xConnectionToken, id, remoteData)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
