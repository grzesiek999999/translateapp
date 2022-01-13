package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
	"translateapp/internal/translateapp"
)

const BaseURL = "http://172.18.0.2:5000/"

type LibreTranslateClient struct {
	logger     *zap.Logger
	BaseURL    string
	httpClient *http.Client
}

func NewLibreTranslateClient(logger *zap.Logger) *LibreTranslateClient {
	return &LibreTranslateClient{
		logger:  logger,
		BaseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c LibreTranslateClient) GetLanguages(ctx context.Context) (*translateapp.ListLanguage, error) {
	var errorRes translateapp.Error
	var languages translateapp.ListLanguage
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/languages", c.BaseURL), nil)
	if err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		errorRes.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&errorRes.Message); err != nil {
			errorRes.Message = err.Error()
			return nil, errorRes
		}
		return nil, errorRes
	}

	if err = json.NewDecoder(res.Body).Decode(&languages.Languages); err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	log.Println(&languages)
	return &languages, nil
}

func (c LibreTranslateClient) Translate(ctx context.Context, word translateapp.WordToTranslate) (*translateapp.WordTranslate, error) {

	var errorRes translateapp.Error
	wordRequestJson, err := json.Marshal(word)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/translate", c.BaseURL), bytes.NewBuffer(wordRequestJson))
	if err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.httpClient.Do(req)

	if err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		errorRes.Code = res.StatusCode
		if err := json.NewDecoder(res.Body).Decode(&errorRes.Message); err != nil {
			errorRes.Message = err.Error()
			return nil, errorRes
		}
		return nil, errorRes
	}
	var translation translateapp.WordTranslate
	if err := json.NewDecoder(res.Body).Decode(&translation); err != nil {
		errorRes.Message = err.Error()
		errorRes.Code = 500
		return nil, errorRes
	}
	c.logger.Info("Value from client")
	return &translation, nil
}
