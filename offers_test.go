package billomat

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestFromOffer(t *testing.T) {
	o := &Offer{
		ClientID: 123,
		OfferItems: []OfferItem{
			{
				Title:     "Product XYZ",
				UnitPrice: 18.99,
				Quantity:  5,
			},
		},
	}
	_, err := fromOffer(o)
	assert.NoError(t, err)
}

func TestToOffer(t *testing.T) {
	expected := `<?xml version="1.0" encoding="UTF-8"?>
	<offer>
	<id type="integer">1234</id>
    <client_id>5</client_id>
    <date>2009-11-17</date>
    <offer-items>
        <offer-item>
            <unit>Stück</unit>
            <unit_price>1.23</unit_price>
            <quantity>1.5</quantity>
            <title>Muster</title>
        </offer-item>
        <offer-item>
            <unit>Stück</unit>
            <unit_price>1.23</unit_price>
            <quantity>5</quantity>
            <title>Muster 2</title>
        </offer-item>
    </offer-items>
	</offer>`
	o, err := toClient([]byte(expected[:]))
	assert.NoError(t, err)

	assert.EqualInt(t, 1234, o.ID)
}
