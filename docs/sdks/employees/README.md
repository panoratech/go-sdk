# Employees
(*Hris.Employees*)

## Overview

### Available Operations

* [List](#list) - List Employees
* [Create](#create) - Create Employees
* [Retrieve](#retrieve) - Retrieve Employee

## List

List Employees

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
    res, err := s.Hris.Employees.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListHrisEmployeesResponse](../../models/operations/listhrisemployeesresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Employees in any supported Hris software

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
    res, err := s.Hris.Employees.Create(ctx, "<value>", components.UnifiedHrisEmployeeInput{
        Groups: []string{
            "Group1",
            "Group2",
        },
        Locations: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        EmployeeNumber: gosdk.String("EMP001"),
        CompanyID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FirstName: gosdk.String("John"),
        LastName: gosdk.String("Doe"),
        PreferredName: gosdk.String("Johnny"),
        DisplayFullName: gosdk.String("John Doe"),
        Username: gosdk.String("johndoe"),
        WorkEmail: gosdk.String("john.doe@company.com"),
        PersonalEmail: gosdk.String("john.doe@personal.com"),
        MobilePhoneNumber: gosdk.String("+1234567890"),
        Employments: []string{
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
            "801f9ede-c698-4e66-a7fc-48d19eebaa4f",
        },
        Ssn: gosdk.String("123-45-6789"),
        Gender: gosdk.String("MALE"),
        Ethnicity: gosdk.String("AMERICAN_INDIAN_OR_ALASKA_NATIVE"),
        MaritalStatus: gosdk.String("Married"),
        DateOfBirth: types.MustNewTimeFromString("1990-01-01"),
        StartDate: types.MustNewTimeFromString("2020-01-01"),
        EmploymentStatus: gosdk.String("ACTIVE"),
        TerminationDate: types.MustNewTimeFromString("2025-01-01"),
        AvatarURL: gosdk.String("https://example.com/avatar.jpg"),
        ManagerID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        FieldMappings: &components.UnifiedHrisEmployeeInputFieldMappings{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisEmployeeOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `xConnectionToken`                                                                         | *string*                                                                                   | :heavy_check_mark:                                                                         | The connection token                                                                       |
| `unifiedHrisEmployeeInput`                                                                 | [components.UnifiedHrisEmployeeInput](../../models/components/unifiedhrisemployeeinput.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `remoteData`                                                                               | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Set to true to include data from the original Hris software.                               |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateHrisEmployeeResponse](../../models/operations/createhrisemployeeresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve an Employee from any connected Hris software

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
    res, err := s.Hris.Employees.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisEmployeeOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |                                                              |
| `xConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |                                                              |
| `id`                                                         | *string*                                                     | :heavy_check_mark:                                           | id of the employee you want to retrieve.                     | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                         |
| `remoteData`                                                 | **bool*                                                      | :heavy_minus_sign:                                           | Set to true to include data from the original Hris software. | false                                                        |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |                                                              |

### Response

**[*operations.RetrieveHrisEmployeeResponse](../../models/operations/retrievehrisemployeeresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
