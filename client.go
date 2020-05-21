package billomat

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiEndpointClients = "clients"

func toClient(data []byte) (*Client, error) {
	var client Client
	err := xml.Unmarshal(data, &client)
	if err != nil {
		return nil, fmt.Errorf("xml.Marshal() failed: %w", err)
	}
	return &client, err
}

func fromClient(client *Client) ([]byte, error) {
	tmp := struct {
		*Client
		XMLName struct{} `xml:"client"`
	}{Client: client}

	data, err := xml.Marshal(tmp)
	if err != nil {
		return data, fmt.Errorf("xml.Marshal() failed: %w", err)
	}
	return data, err
}

// ClientAdd creates a new client via API
func (b *Billomat) ClientAdd(client *Client) (*Client, error) {
	var createdClient *Client
	body, err := fromClient(client)
	if err != nil {
		return createdClient, err
	}

	url := b.generateURL(apiEndpointClients)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return createdClient, fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return createdClient, fmt.Errorf("billomat.httpClient.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return createdClient, fmt.Errorf("error while reading response body: %v", err)
	}
	if httpResp.StatusCode != http.StatusCreated {
		return createdClient, fmt.Errorf("unexpected HTTP status code %d (body: %q)", httpResp.StatusCode, string(respBody))
	}

	createdClient, err = toClient(respBody)
	if err != nil {
		return createdClient, fmt.Errorf("client created, but cannot parse response: %v (body: %q)", err, string(respBody))
	}
	return createdClient, nil
}
