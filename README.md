# Golang library to connect to Billit

## Billit API client for Go

## Swagger file
https://api.billit.be/swagger/docs/v1

## Postman collection
https://www.postman.com/billit/billit-quickstart-guide/request/aciw1ik/01-create-sales-invoice

## Billit API documentation
https://docs.billit.be/docs/getting-started
https://docs.billit.be/reference/

## Using swagger
```bash
go install github.com/go-swagger/go-swagger/cmd/swagger@latest 
```

```bash
swagger generate client -f billit.v1.json --client-package=client --model-package=model
```

## Example
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/becoded/billit/client/account"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/becoded/billit/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	r := httptransport.New(
		"api.sandbox.billit.be",
		apiclient.DefaultBasePath,
		apiclient.DefaultSchemes,
	)
	r.DefaultAuthentication = httptransport.APIKeyAuth(
		"apiKey",
		"header",
		os.Getenv("API_KEY"),
	)

	client := apiclient.New(r, strfmt.Default)

	// make the authenticated request to get all items
	params := account.NewAccountGetAccountInformationParams()
	resp, err := client.Account.AccountGetAccountInformation(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.Payload)
}
```