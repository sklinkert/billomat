# billomat
Unofficial billomat.com API Go/Golang SDK

https://www.billomat.com/api/

## Examples

### Creating new Client

```go
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/sklinkert/billomat"
)

func main() {
	const apiKey = "abcdefgh"
  const billomatID = "mycompany"
  const appID = "12345"
  const appSecret = "ejkrfierfnebnfuernferufrbef/ejenir"
	b := billomat.New(billomatID, appID, appSecret, apiKey)

	client := &billomat.Client{
		Name:        "Musterfirma",
		Salutation:  "Herr",
		FirstName:   "Max",
		LastName:    "Muster",
		Street:      "Musterstraße 123",
		Zip:         "12345",
		State:       "Bundesland",
		CountryCode: "DE",
		Phone:       "+49 123456789",
		Email:       "info@example.com",
	}
	if err := b.ClientAdd(client); err != nil {
		log.WithError(err).Error("Adding client failed")
	}
}

```

