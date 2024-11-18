# Retryid
(*Passthrough.{retryid}*)

## Overview

### Available Operations

* [GetRetriedRequestResponse](#getretriedrequestresponse) - Retrieve response of a failed passthrough request due to rate limits

## GetRetriedRequestResponse

Retrieve response of a failed passthrough request due to rate limits

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
    res, err := s.Passthrough.{retryid}.GetRetriedRequestResponse(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `retryID`                                                             | *string*                                                              | :heavy_check_mark:                                                    | id of the retryJob returned when you initiated a passthrough request. |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.GetRetriedRequestResponseResponse](../../models/operations/getretriedrequestresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |