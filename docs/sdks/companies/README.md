# Companies
(*Crm.Companies*)

### Available Operations

* [List](#list) - List Companies
* [Create](#create) - Create Companies
* [Retrieve](#retrieve) - Retrieve Companies

## List

List Companies

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

    unifiedCrmCompanyInput := components.UnifiedCrmCompanyInput{
        Name: gosdk.String("Acme"),
        Industry: components.UnifiedCrmCompanyInputIndustryAccounting.ToPointer(),
        NumberOfEmployees: gosdk.Float64(10),
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        EmailAddresses: []components.Email{
            components.Email{
                EmailAddress: gosdk.String("acme@gmail.com"),
                EmailAddressType: components.EmailAddressTypeWork.ToPointer(),
            },
        },
        Addresses: []components.Address{
            components.Address{
                Street1: gosdk.String("5th Avenue"),
                Street2: gosdk.String("<value>"),
                City: gosdk.String("New York"),
                State: gosdk.String("NY"),
                PostalCode: gosdk.String("46842"),
                Country: gosdk.String("USA"),
                AddressType: components.AddressTypeWork.ToPointer(),
                OwnerType: gosdk.String("<value>"),
            },
        },
        PhoneNumbers: []components.Phone{
            components.Phone{
                PhoneNumber: gosdk.String("+33660606067"),
                PhoneType: components.PhoneTypeWork.ToPointer(),
            },
        },
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }

    var remoteData *bool = gosdk.Bool(false)
    ctx := context.Background()
    res, err := s.Crm.Companies.Create(ctx, xConnectionToken, unifiedCrmCompanyInput, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmCompanyOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `xConnectionToken`                                                                     | *string*                                                                               | :heavy_check_mark:                                                                     | The connection token                                                                   |                                                                                        |
| `unifiedCrmCompanyInput`                                                               | [components.UnifiedCrmCompanyInput](../../models/components/unifiedcrmcompanyinput.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `remoteData`                                                                           | **bool*                                                                                | :heavy_minus_sign:                                                                     | Set to true to include data from the original CRM software.                            | false                                                                                  |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |


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
    res, err := s.Crm.Companies.Retrieve(ctx, xConnectionToken, id, remoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmCompanyOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the company you want to retrieve.                     | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original Crm software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |


### Response

**[*operations.RetrieveCrmCompanyResponse](../../models/operations/retrievecrmcompanyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
