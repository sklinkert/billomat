package billomat

// Client in order to create invoices, offers, etc.
type Client struct {
	Archived            bool    `xml:"archived,omitempty"`
	NumberPre           string  `xml:"number_pre,omitempty"`
	Number              int     `xml:"number,omitempty"`
	NumberLength        string  `xml:"number_length,omitempty"`
	Name                string  `xml:"name,omitempty"`
	Street              string  `xml:"street,omitempty"`
	Zip                 string  `xml:"zip,omitempty"`
	City                string  `xml:"city,omitempty"`
	State               string  `xml:"state,omitempty"`
	CountryCode         string  `xml:"country_code,omitempty"`
	FirstName           string  `xml:"first_name,omitempty"`
	LastName            string  `xml:"last_name,omitempty"`
	Salutation          string  `xml:"salutation,omitempty"`
	Phone               string  `xml:"phone,omitempty"`
	Fax                 string  `xml:"fax,omitempty"`
	Mobile              string  `xml:"mobile,omitempty"`
	Email               string  `xml:"email,omitempty"`
	WWW                 string  `xml:"www,omitempty"`
	TaxNumber           string  `xml:"tax_number,omitempty"`
	VatNumber           string  `xml:"vat_number,omitempty"`
	BankAccountNumber   string  `xml:"bank_account_number,omitempty"`
	BankAccountOwner    string  `xml:"bank_account_owner,omitempty"`
	BankNumber          string  `xml:"bank_number,omitempty"`
	BankName            string  `xml:"bank_name,omitempty"`
	BankSwift           string  `xml:"bank_swift,omitempty"`
	BankIBAN            string  `xml:"bank_iban,omitempty"`
	SepaMandate         string  `xml:"sepa_mandate,omitempty"`
	SepaMandateDate     string  `xml:"sepa_mandate_date,omitempty"`
	Locale              string  `xml:"locale,omitempty"`
	TaxRule             string  `xml:"tax_rule,omitempty"`
	NetGross            string  `xml:"net_gross,omitempty"`
	DefaultPaymentTypes string  `xml:"default_payment_types,omitempty"`
	Note                string  `xml:"note,omitempty"`
	Reduction           float64 `xml:"reduction,omitempty"`
}
