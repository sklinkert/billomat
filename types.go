package billomat

import "time"

// Client in order to create invoices, offers, etc.
type Client struct {
	ID                  int       `xml:"id,omitempty"`
	ClientNumber        string    `xml:"client_number,omitempty"`
	Created             time.Time `xml:"created,omitempty"`
	Archived            bool      `xml:"archived,omitempty"`
	NumberPre           string    `xml:"number_pre,omitempty"`
	Number              int       `xml:"number,omitempty"`
	NumberLength        string    `xml:"number_length,omitempty"`
	Name                string    `xml:"name,omitempty"`
	Street              string    `xml:"street,omitempty"`
	Zip                 string    `xml:"zip,omitempty"`
	City                string    `xml:"city,omitempty"`
	State               string    `xml:"state,omitempty"`
	CountryCode         string    `xml:"country_code,omitempty"`
	FirstName           string    `xml:"first_name,omitempty"`
	LastName            string    `xml:"last_name,omitempty"`
	Salutation          string    `xml:"salutation,omitempty"`
	Phone               string    `xml:"phone,omitempty"`
	Fax                 string    `xml:"fax,omitempty"`
	Mobile              string    `xml:"mobile,omitempty"`
	Email               string    `xml:"email,omitempty"`
	WWW                 string    `xml:"www,omitempty"`
	TaxNumber           string    `xml:"tax_number,omitempty"`
	VatNumber           string    `xml:"vat_number,omitempty"`
	BankAccountNumber   string    `xml:"bank_account_number,omitempty"`
	BankAccountOwner    string    `xml:"bank_account_owner,omitempty"`
	BankNumber          string    `xml:"bank_number,omitempty"`
	BankName            string    `xml:"bank_name,omitempty"`
	BankSwift           string    `xml:"bank_swift,omitempty"`
	BankIBAN            string    `xml:"bank_iban,omitempty"`
	SepaMandate         string    `xml:"sepa_mandate,omitempty"`
	SepaMandateDate     string    `xml:"sepa_mandate_date,omitempty"`
	Locale              string    `xml:"locale,omitempty"`
	TaxRule             string    `xml:"tax_rule,omitempty"`
	NetGross            string    `xml:"net_gross,omitempty"`
	DefaultPaymentTypes string    `xml:"default_payment_types,omitempty"`
	Note                string    `xml:"note,omitempty"`
	Reduction           float64   `xml:"reduction,omitempty"`
}

// Offer represents an offer for an Client
type Offer struct {
	ID           int       `xml:"id,omitempty"`
	ClientID     int       `xml:"client_id,omitempty"`
	ContactID    string    `xml:"contact_id,omitempty"`
	Created      time.Time `xml:"created,omitempty"`
	Address      string    `xml:"address,omitempty"`
	NumberPre    string    `xml:"number_pre,omitempty"`
	Number       int       `xml:"number,omitempty"`
	NumberLength string    `xml:"number_length,omitempty"`
	// Date         time.Time `xml:"date,omitempty"`
	Title        string  `xml:"title,omitempty"`
	Label        string  `xml:"label,omitempty"`
	Intro        string  `xml:"intro,omitempty"`
	Note         string  `xml:"note,omitempty"`
	Reduction    float64 `xml:"reduction,omitempty"`
	CountryCode  string  `xml:"country_code,omitempty"`
	NetGross     string  `xml:"net_gross,omitempty"`
	Quote        float64 `xml:"quote,omitempty"`
	ValidityDays int     `xml:"validity_days,omitempty"`
	// ValidityDate time.Time   `xml:"validity_date,omitempty"`
	FreeTextID string      `xml:"free_text_id,omitempty"`
	TemplateID string      `xml:"template_id,omitempty"`
	OfferItems []OfferItem `xml:"offer-items>offer-item,omitempty"`
}

// OfferItem represents an item for an offer
type OfferItem struct {
	Unit        string  `xml:"unit,omitempty"`
	UnitPrice   float64 `xml:"unit_price,omitempty"`
	Quantity    int     `xml:"quantity,omitempty"`
	Title       string  `xml:"title,omitempty"`
	Description string  `xml:"description,omitempty"`
	ArticleID   int     `xml:"article_id,omitempty"`
	Optional    int     `xml:"optional,omitempty"`
	TaxRate     float64 `xml:"tax_rate,omitempty"`
	TaxName     string  `xml:"tax_name,omitempty"`
}
