# Candidates
(*Ats.Candidates*)

### Available Operations

* [List](#list) - List  Candidates
* [Create](#create) - Create Candidates
* [Retrieve](#retrieve) - Retrieve Candidates

## List

List  Candidates

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
    res, err := s.Ats.Candidates.List(ctx, xConnectionToken, nil, nil, nil)
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

**[*operations.ListAtsCandidateResponse](../../models/operations/listatscandidateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create Candidates in any supported Ats software

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

    unifiedAtsCandidateInput := components.UnifiedAtsCandidateInput{
        FirstName: gosdk.String("Joe"),
        LastName: gosdk.String("Doe"),
        Company: gosdk.String("Acme"),
        Title: gosdk.String("Analyst"),
        Locations: gosdk.String("New York"),
        IsPrivate: gosdk.Bool(false),
        EmailReachable: gosdk.Bool(true),
        RemoteCreatedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        RemoteModifiedAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        LastInteractionAt: types.MustNewTimeFromString("2024-10-01T12:00:00Z"),
        Attachments: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Applications: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Tags: []string{
            "tag_1",
            "tag_2",
        },
        Urls: []components.URL{
            components.URL{
                URL: gosdk.String("mywebsite.com"),
                URLType: gosdk.String("WEBSITE"),
            },
        },
        PhoneNumbers: []components.Phone{
            components.Phone{
                PhoneNumber: gosdk.String("+33660688899"),
                PhoneType: components.PhoneTypeWork.ToPointer(),
            },
        },
        EmailAddresses: []components.Email{
            components.Email{
                EmailAddress: gosdk.String("joedoe@gmail.com"),
                EmailAddressType: components.EmailAddressTypeWork.ToPointer(),
            },
        },
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Ats.Candidates.Create(ctx, xConnectionToken, unifiedAtsCandidateInput, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsCandidateOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `xConnectionToken`                                                                         | *string*                                                                                   | :heavy_check_mark:                                                                         | The connection token                                                                       |                                                                                            |
| `unifiedAtsCandidateInput`                                                                 | [components.UnifiedAtsCandidateInput](../../models/components/unifiedatscandidateinput.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `remoteData`                                                                               | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Set to true to include data from the original Ats software.                                | false                                                                                      |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |


### Response

**[*operations.CreateAtsCandidateResponse](../../models/operations/createatscandidateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieve Candidates from any connected Ats software

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
    res, err := s.Ats.Candidates.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedAtsCandidateOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the candidate you want to retrieve.                   | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Ats software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |


### Response

**[*operations.RetrieveAtsCandidateResponse](../../models/operations/retrieveatscandidateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
