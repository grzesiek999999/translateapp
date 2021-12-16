package translateapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const BaseURL = "http://172.18.0.2:5000"

type LibreTranslateClient struct {
	BaseURL    string
	httpClient *http.Client
}

func NewLibreTranslateClient() *LibreTranslateClient {
	return &LibreTranslateClient{
		BaseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c LibreTranslateClient) GetLanguages(ctx context.Context) ([]Language, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/languages", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	req = req.WithContext(ctx)

	if err != nil {
		return nil, err
	}
	var languages []Language
	err = json.NewDecoder(res.Body).Decode(&languages)
	return languages, err
}

//func (c LibreTranslateClient) Translate(ctx context.Context, from, to, text string) (string, error) {
//
//}
//
//func (c LibreTranslateClient) call(ctx context.Context, url string) (string, error) {
//	fullURL := c.baseURL + url
//	return fullURL
//}
