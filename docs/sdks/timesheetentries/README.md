# Timesheetentries
(*Hris.Timesheetentries*)

## Overview

### Available Operations

* [List](#list) - List Timesheetentries
* [Create](#create) - Create Timesheetentrys
* [Retrieve](#retrieve) - Retrieve Timesheetentry

## List

List Timesheetentries

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
    res, err := s.Hris.Timesheetentries.List(ctx, "<value>", gosdk.Bool(true), gosdk.Float64(10), gosdk.String("1b8b05bb-5273-4012-b520-8657b0b90874"))
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

**[*operations.ListHrisTimesheetentriesResponse](../../models/operations/listhristimesheetentriesresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Create

Create Timesheetentrys in any supported Hris software

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
    res, err := s.Hris.Timesheetentries.Create(ctx, "<value>", components.UnifiedHrisTimesheetEntryInput{
        HoursWorked: gosdk.Float64(40),
        StartTime: types.MustNewTimeFromString("2024-10-01T08:00:00Z"),
        EndTime: types.MustNewTimeFromString("2024-10-01T16:00:00Z"),
        EmployeeID: gosdk.String("801f9ede-c698-4e66-a7fc-48d19eebaa4f"),
        RemoteWasDeleted: gosdk.Bool(false),
        FieldMappings: &components.UnifiedHrisTimesheetEntryInputFieldMappings{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisTimesheetEntryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `xConnectionToken`                                                                                     | *string*                                                                                               | :heavy_check_mark:                                                                                     | The connection token                                                                                   |
| `unifiedHrisTimesheetEntryInput`                                                                       | [components.UnifiedHrisTimesheetEntryInput](../../models/components/unifiedhristimesheetentryinput.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `remoteData`                                                                                           | **bool*                                                                                                | :heavy_minus_sign:                                                                                     | Set to true to include data from the original Hris software.                                           |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateHrisTimesheetentryResponse](../../models/operations/createhristimesheetentryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## Retrieve

Retrieve an Timesheetentry from any connected Hris software

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
    res, err := s.Hris.Timesheetentries.Retrieve(ctx, "<value>", "801f9ede-c698-4e66-a7fc-48d19eebaa4f", gosdk.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnifiedHrisTimesheetEntryOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |                                                              |
| `xConnectionToken`                                           | *string*                                                     | :heavy_check_mark:                                           | The connection token                                         |                                                              |
| `id`                                                         | *string*                                                     | :heavy_check_mark:                                           | id of the timesheetentry you want to retrieve.               | 801f9ede-c698-4e66-a7fc-48d19eebaa4f                         |
| `remoteData`                                                 | **bool*                                                      | :heavy_minus_sign:                                           | Set to true to include data from the original Hris software. | false                                                        |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |                                                              |

### Response

**[*operations.RetrieveHrisTimesheetentryResponse](../../models/operations/retrievehristimesheetentryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
