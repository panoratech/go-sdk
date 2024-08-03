<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(os.Getenv("API_KEY")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->