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
}

```

### Create Client

```go
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
	createdClient, err := b.CreateClient(client)
	if err != nil {
		log.WithError(err).Fatal("Adding client failed")
	}
	log.Infof("Client ID is %d. Created at %v", createdClient.ID, createdClient.Created)

```

### Delete Client

```go
if err := b.DeleteClient(createdClient); err != nil {
	log.WithError(err).Fatal("deleting client failed")
}
```

### Create Offer

```go
	offer := &billomat.Offer{
		ClientID: createdClient.ID,
		OfferItems: []billomat.OfferItem{
			{
				Title:     "Product XYZ",
				UnitPrice: 18.99,
				Quantity:  5,
			},
		},
	}
	createdOffer, err := b.CreateOffer(offer)
	if err != nil {
		log.WithError(err).Fatal("Adding offer failed")
	}
	log.Infof("Offer ID is %d", createdOffer.ID)
```

### Delete Offer

```go
if err := b.DeleteOffer(createdOffer); err != nil {
	log.WithError(err).Fatal("deleting offer failed")
}
```



### Create Article

```go
article := &billomat.Article{
  Title: "ETF-Sparplan von ETFs24.de"
  SalePrice: 9.99,
  // ...
}
createdArticle, err := b.CreateArticle(article)
if err != nil {
	log.WithError(err).Fatal("Adding article failed")
}
```

### Get Article(s)

```go
// by article number
articles, err := b.GetArticlesByNumber("D-1830-54")
if err != nil {
	log.WithError(err).Fatal("getting article failed")
}

// or by ID
article, err := b.GetArticleByID(287)
if err != nil {
	log.WithError(err).Fatal("getting article failed")
}
```

