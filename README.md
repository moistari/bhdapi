# About

BHD API. See: [Go reference](https://pkg.go.dev/github.com/moistari/bhdapi)

Using:

```sh
go get github.com/moistari/bhdapi
```

Example:

```go
// example/example.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/moistari/bhdapi"
)

func main() {
	api := flag.String("api", "", "api key")
	flag.Parse()
	cl := bhdapi.New(bhdapi.WithApiKey(*api))
	res, err := cl.Search(context.Background(), "Fight Club")
	if err != nil {
		log.Fatal(err)
	}
	for i, r := range res.Results {
		fmt.Printf("%02d: %s\n", i, r.Name)
	}
}
```
