# Interviews
(*Ats.Interviews*)

### Available Operations

* [List](#list) - List  Interviews
* [Create](#create) - Create Interviews
* [Retrieve](#retrieve) - Retrieve Interviews

## List

List  Interviews

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
    res, err := s.Ats.Interviews.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAtsInterviewResponse](../../models/operations/listatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Interviews in any supported Ats software

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

    unifiedAtsInterviewInput := components.UnifiedAtsInterviewInput{
        Status: components.UnifiedAtsInterviewInputStatusScheduled.ToPointer(),
        ApplicationID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        JobInterviewStageID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        OrganizedBy: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        Interviewers: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Location: gosdk.String("San Francisco"),
        StartAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        EndAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        RemoteCreatedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        RemoteUpdatedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ats.Interviews.Create(ctx, xConnectionToken, unifiedAtsInterviewInput, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsInterviewOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `xConnectionToken`                                                                         | *string*                                                                                   | :heavy_check_mark:                                                                         | The connection token                                                                       |                                                                                            |
| `unifiedAtsInterviewInput`                                                                 | [components.UnifiedAtsInterviewInput](../../models/components/unifiedatsinterviewinput.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `remoteData`                                                                               | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Set to true to include data from the original Ats software.                                | false                                                                                      |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |


### Response

**[*operations.CreateAtsInterviewResponse](../../models/operations/createatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Interviews from any connected Ats software

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
    res, err := s.Ats.Interviews.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsInterviewOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the interview you want to retrieve.                   | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Ats software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |


### Response

**[*operations.RetrieveAtsInterviewResponse](../../models/operations/retrieveatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
