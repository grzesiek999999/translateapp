package translateapp_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"
	"translateapp/internal/logging"
	"translateapp/internal/translateapp"
	"translateapp/mocks"
)

func TestBatchTranslate(t *testing.T) {
	input := []translateapp.WordToTranslate{
		{
			Word:   "computer",
			Source: "",
			Target: "",
		},
		{
			Word:   "night",
			Source: "",
			Target: "",
		},
	}
	expected := []translateapp.TranslateResponse{
		{
			Code:    200,
			Message: "succes",
			Data: translateapp.WordTranslate{
				Text: "komputer",
			},
		},
		{
			Code:    200,
			Message: "succes",
			Data: translateapp.WordTranslate{
				Text: "noc",
			},
		},
	}
	testFile := path.Join("testuploaddata", "text")
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(file)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("text", filepath.Base(testFile))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}
	service := mocks.Servicer{}
	service.On("Translate", mock.Anything, input).Return(&expected, nil)
	app := translateapp.NewApp(&service, logging.DefaultLogger().Desugar())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/batchTranslate", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	translation := app.BatchTranslate()
	translation.ServeHTTP(w, req)
	var response []translateapp.TranslateResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, expected, response)
}

func TestTranslate(t *testing.T) {
	input := []translateapp.WordToTranslate{
		{
			Word:   "computer",
			Source: "",
			Target: "",
		},
	}
	expected := []translateapp.TranslateResponse{
		{
			Code:    200,
			Message: "succes",
			Data: translateapp.WordTranslate{
				Text: "komputer",
			},
		},
	}
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(input[0])
	require.NoError(t, err)
	service := mocks.Servicer{}
	service.On("Translate", mock.Anything, input).Return(&expected, nil)
	app := translateapp.NewApp(&service, logging.DefaultLogger().Desugar())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/Translate", &body)
	translation := app.Translate()
	translation.ServeHTTP(w, req)
	var response []translateapp.TranslateResponse
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, expected, response)
}

func TestGetLanguages(t *testing.T) {
	expected := translateapp.Response{
		Code:    200,
		Message: "success",
		Data: translateapp.ListLanguage{
			Languages: []translateapp.Language{
				{
					Name: "polish",
					Code: "pl",
				},
			},
		},
	}
	service := mocks.Servicer{}
	service.On("GetLanguages", mock.Anything).Return(&expected, nil)
	app := translateapp.NewApp(&service, logging.DefaultLogger().Desugar())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/GetLanguages", nil)
	translation := app.GetLanguages()
	translation.ServeHTTP(w, req)
	var response translateapp.Response
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, expected, response)
}
