<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
)

func main() {
	s := gosdk.New()

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.String != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->