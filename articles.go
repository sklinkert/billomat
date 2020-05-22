package billomat

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiEndpointArticles = "articles"

func toArticles(data []byte) (*[]Article, error) {
	tmp := struct {
		Articles []Article `xml:"article"`
	}{}
	if err := xml.Unmarshal(data, &tmp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal() failed: %w", err)
	}
	return &tmp.Articles, nil
}

func toArticle(data []byte) (*Article, error) {
	var article Article
	if err := xml.Unmarshal(data, &article); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal() failed: %w", err)
	}
	return &article, nil
}

func fromArticle(article *Article) ([]byte, error) {
	tmp := struct {
		*Article
		XMLName struct{} `xml:"article"`
	}{Article: article}

	data, err := xml.Marshal(tmp)
	if err != nil {
		return data, fmt.Errorf("xml.Marshal() failed: %w", err)
	}
	return data, err
}

// GetArticleByID fetch article by article ID
func (b *Billomat) GetArticleByID(articleID int) (*Article, error) {
	url := b.generateURL(apiEndpointArticles, articleID)
	return b.getArticle(url)
}

// GetArticlesByNumber fetch article by article number
func (b *Billomat) GetArticlesByNumber(articleNumber string) (*[]Article, error) {
	url := b.generateURL(apiEndpointArticles, 0)
	url = fmt.Sprintf("%s?article_number=%s", url, articleNumber)
	return b.getArticles(url)
}

func (b *Billomat) getArticles(url string) (*[]Article, error) {
	var articles *[]Article
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return articles, fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return articles, fmt.Errorf("billomat.httpClient.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return articles, fmt.Errorf("error while reading response body: %w", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		return articles, fmt.Errorf("unexpected HTTP status code %d (body: %q)",
			httpResp.StatusCode, string(respBody))
	}

	articles, err = toArticles(respBody)
	if err != nil {
		return articles, fmt.Errorf("article created, but cannot parse response: %w (body: %q)",
			err, string(respBody))
	}
	return articles, nil
}

func (b *Billomat) getArticle(url string) (*Article, error) {
	var article *Article
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return article, fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return article, fmt.Errorf("billomat.httpClient.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return article, fmt.Errorf("error while reading response body: %w", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		return article, fmt.Errorf("unexpected HTTP status code %d (body: %q)",
			httpResp.StatusCode, string(respBody))
	}

	article, err = toArticle(respBody)
	if err != nil {
		return article, fmt.Errorf("article created, but cannot parse response: %w (body: %q)",
			err, string(respBody))
	}
	return article, nil
}

// CreateArticle creates a new article via API
func (b *Billomat) CreateArticle(article *Article) (*Article, error) {
	var createdArticle *Article
	body, err := fromArticle(article)
	if err != nil {
		return createdArticle, err
	}

	url := b.generateURL(apiEndpointArticles, 0)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return createdArticle, fmt.Errorf("http.NewRequest() failed: %w", err)
	}
	b.setAuthHeader(httpReq)

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return createdArticle, fmt.Errorf("billomat.httpClient.Do() failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return createdArticle, fmt.Errorf("error while reading response body: %w", err)
	}
	if httpResp.StatusCode != http.StatusCreated {
		return createdArticle, fmt.Errorf("unexpected HTTP status code %d (body: %q)",
			httpResp.StatusCode, string(respBody))
	}

	createdArticle, err = toArticle(respBody)
	if err != nil {
		return createdArticle, fmt.Errorf("article created, but cannot parse response: %w (body: %q)",
			err, string(respBody))
	}
	return createdArticle, nil
}

// DeleteArticle delete an existing article via API
func (b *Billomat) DeleteArticle(article *Article) error {
	url := b.generateURL(apiEndpointArticles, article.ID)
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
		return fmt.Errorf("unexpected HTTP status code %d (body: %q)",
			httpResp.StatusCode, string(respBody))
	}
	return nil
}
