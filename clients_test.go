package billomat

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestFromClient(t *testing.T) {
	c := &Client{
		Name:        "Musterfirma",
		Salutation:  "Herr",
		FirstName:   "Max",
		LastName:    "Muster",
		Street:      "Musterstraße 123",
		Zip:         "12345",
		State:       "Bundesland",
		CountryCode: "DE",
		Phone:       "+49 123456789",
		Mobile:      "+49 123456789",
		Email:       "info@example.com",
		WWW:         "www.example.com",
	}
	_, err := fromClient(c)
	assert.NoError(t, err)
}

func TestToClient(t *testing.T) {
	expected := `<?xml version="1.0" encoding="UTF-8"?>
	<client>
		<id type="integer">1</id>
		<created type="datetime">2007-12-13T12:12:00+01:00</created>
		<archived>0</archived>
		<client_number>KD123</client_number>
		<number type="integer">123</number>
		<number_pre>KD</number_pre>
		<number_length type="integer">0</number_length>
		<name>Musterfirma</name>
		<salutation>Herr</salutation>
		<first_name>Max</first_name>
		<last_name>Muster</last_name>
		<street>Musterstraße 123</street>
		<zip>12345</zip>
		<city>Musterstadt</city>
		<state>Bundesland</state>
		<country_code>DE</country_code>
		<address>Herr Max Muster
	Musterstraße 123
	12345 Musterstadt</address> <!-- read only -->
		<phone>+49 123456789</phone>
		<fax>+49 123456789</fax>
		<mobile>+49 123456789</mobile>
		<email>info@example.com</email>
		<www>www.example.com</www>
		<tax_number>12/3456/789</tax_number>
		<vat_number>DE123456789</vat_number>
		<bank_account_owner>Kontoinhaber</bank_account_owner>
		<bank_number>123456789</bank_number>
		<bank_name>Deutsche Bank</bank_name>
		<bank_account_number>123456789</bank_account_number>
		<bank_swift>SWIFT/BIC</bank_swift>
		<bank_iban>IBAN</bank_iban>
		<enable_customerportal>1</enable_customerportal>
		<customerportal_url>https://mybillomatid.billomat.net/customerportal/auth/autologin/entityId/123?hash=123456789aabbcc</customerportal_url>
		<sepa_mandate>MN123</sepa_mandate>
		<sepa_mandate_date>2013-12-10</sepa_mandate_date>
		<tax_rule>COUNTRY</tax_rule>
		<net_gross>SETTINGS</net_gross>
		<default_payment_types>CASH,PAPAL</default_payment_types>
		<reduction>10</reduction>
		<discount_rate_type>SETTINGS</discount_rate_type>
		<discount_rate>2</discount_rate>
		<discount_days_type>SETTINGS</discount_days_type>
		<discount_days>7</discount_days>
		<due_days_type>SETTINGS</due_days_type>
		<due_days>14</due_days>
		<reminder_due_days_type>RELATIVE</reminder_due_days_type>
		<reminder_due_days>-2</reminder_due_days>
		<offer_validity_days_type>ABSOLUTE</offer_validity_days_type>
		<offer_validity_days>30</offer_validity_days>
		<currency_code>CHF</currency_code>
		<price_group>2</price_group>
		<debitor_account_number>10000</debitor_account_number>
		<dunning_run>0</dunning_run>
		<note>Sehr netter Kunde</note>
		<revenue_gross type="float">119</revenue_gross> <!-- read only -->
		<revenue_net type="float">100</revenue_net> <!-- read only -->
	</client>`
	c, err := toClient([]byte(expected[:]))
	assert.NoError(t, err)

	assert.EqualStrings(t, "Musterfirma", c.Name)
	assert.EqualStrings(t, "info@example.com", c.Email)
}
