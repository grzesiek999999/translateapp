package translateapp_test

import (
	"testing"
)

func TestBatchTranslate(t *testing.T) {
	//expected := []translateapp.BatchTranslateResponse{
	//	{
	//		WordToTranslate: "computer",
	//		WordTranslated:  "komputer",
	//	},
	//}
	//body := &bytes.Buffer{}
	//writer := multipart.NewWriter(body)
	//service := mocks.Servicer{}
	//app := translateapp.NewApp(&service, logging.DefaultLogger().Desugar())
	//w := httptest.NewRecorder()
	//req := httptest.NewRequest(http.MethodPost, "/batchTranslate", body)
	//req.Header.Add("Content-Type", writer.FormDataContentType())
	//translation := app.BatchTranslate()
	//translation.ServeHTTP(w, req)
	//assert.Equal(t, w.Body, expected)
}
