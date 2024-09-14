# PanoraContacts
(*Crm.Contacts*)

## Overview

### Available Operations

* [List](#list) - List CRM Contacts
* [Create](#create) - Create Contacts
* [Retrieve](#retrieve) - Retrieve Contacts

## List

List CRM Contacts

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
    res, err := s.Crm.Contacts.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListCrmContactsResponse](../../models/operations/listcrmcontactsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Contacts in any supported CRM

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
    res, err := s.Crm.Contacts.Create(ctx, "<value>", components.UnifiedCrmContactInput{
        FirstName: gosdk.String("John"),
        LastName: gosdk.String("Doe"),
        EmailAddresses: []components.Email{
            components.Email{
                EmailAddress: gosdk.String("Jena.Nienow28@yahoo.com"),
                EmailAddressType: gosdk.String("<value>"),
            },
        },
        PhoneNumbers: []components.Phone{
            components.Phone{
                PhoneNumber: gosdk.String("1-809-839-8041"),
                PhoneType: gosdk.String("<value>"),
            },
        },
        Addresses: []components.Address{
            components.Address{
                Street1: gosdk.String("5th Avenue"),
                Street2: gosdk.String("Street 2"),
                City: gosdk.String("Anytown"),
                State: gosdk.String("CA"),
                PostalCode: gosdk.String("10001"),
                Country: gosdk.String("USA"),
                AddressType: gosdk.String("PERSONAL"),
                OwnerType: gosdk.String("<value>"),
            },
        },
        UserID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: map[string]any{
            "fav_dish": "broccoli",
            "fav_color": "red",
        },
    }, gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmContactOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `xConnectionToken`                                                                     | *string*                                                                               | :heavy_check_mark:                                                                     | The connection token                                                                   |                                                                                        |
| `unifiedCrmContactInput`                                                               | [components.UnifiedCrmContactInput](../../models/components/unifiedcrmcontactinput.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `remoteData`                                                                           | **bool*                                                                                | :heavy_minus_sign:                                                                     | Set to true to include data from the original CRM software.                            | false                                                                                  |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.CreateCrmContactResponse](../../models/operations/createcrmcontactresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve Contacts from any connected CRM

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
    res, err := s.Crm.Contacts.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedCrmContactOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |                                                             |
| `xConnectionToken`                                          | *string*                                                    | :heavy_check_mark:                                          | The connection token                                        |                                                             |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | id of the `contact` you want to retrive.                    | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                        |
| `remoteData`                                                | **bool*                                                     | :heavy_minus_sign:                                          | Set to true to include data from the original CRM software. | false                                                       |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |                                                             |

### Response

**[*operations.RetrieveCrmContactResponse](../../models/operations/retrievecrmcontactresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
