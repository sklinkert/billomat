package billomat

import (
	"fmt"
	"net/http"
	"time"
)

// Billomat contains the API config
type Billomat struct {
	billomatID string
	appSecret  string
	appID      string
	apiKey     string
	httpClient http.Client
}

func (b *Billomat) setAuthHeader(req *http.Request) {
	req.Header["X-AppId"] = []string{b.appID}
	req.Header["X-AppSecret"] = []string{b.appSecret}
	req.Header["X-BillomatApiKey"] = []string{b.apiKey}
	req.Header["Content-Type"] = []string{"application/xml"}

}

func (b *Billomat) generateURL(endpoint string) string {
	return fmt.Sprintf("https://%s.billomat.net/api/%s", b.billomatID, endpoint)
}

// New creates a new object
func New(billomatID, appID, appSecret, apiKey string) *Billomat {
	return &Billomat{
		billomatID: billomatID,
		apiKey:     apiKey,
		appSecret:  appSecret,
		appID:      appID,
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}
