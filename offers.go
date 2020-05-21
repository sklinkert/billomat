package billomat

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiEndpointOffers = "offers"

func toOffer(data []byte) (*Offer, error) {
	var offer Offer
	err := xml.Unmarshal(data, &offer)
	if err != nil {
		return nil, fmt.Errorf("xml.Marshal() failed: %w", err)
	}
	return &offer, err
}

func fromOffer(offer *Offer) ([]byte, error) {
	tmp := struct {
		*Offer
		XMLName struct{} `xml:"offer"`
	}{Offer: offer}

	data, err := xml.Marshal(tmp)
	if err != nil {
		return data, fmt.Errorf("xml.Marshal() failed: %w", err)
	}
	return data, err
}

// OfferAdd creates a new offer via API
func (b *Billomat) OfferAdd(offer *Offer) (*Offer, error) {
	var createdOffer *Offer
	body, err := fromOffer(offer)
	if err != nil {
		return createdOffer, err
	}

	url := b.generateURL(apiEndpointOffers, 0)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return createdOffer, fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return createdOffer, fmt.Errorf("billomat.httpOffer.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return createdOffer, fmt.Errorf("error while reading response body: %w", err)
	}
	if httpResp.StatusCode != http.StatusCreated {
		return createdOffer, fmt.Errorf("unexpected HTTP status code %d (body: %q)", httpResp.StatusCode, string(respBody))
	}

	createdOffer, err = toOffer(respBody)
	if err != nil {
		return createdOffer, fmt.Errorf("offer created, but cannot parse response: %w (body: %q)", err, string(respBody))
	}
	return createdOffer, nil
}

// OfferDelete delete an existing offer via API
func (b *Billomat) OfferDelete(offer *Offer) error {
	url := b.generateURL(apiEndpointOffers, offer.ID)
	httpReq, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("billomat.httpClient.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("error while reading response body: %w", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status code %d (body: %q)", httpResp.StatusCode, string(respBody))
	}
	return nil
}
