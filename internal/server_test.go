package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlergetLanguage(t *testing.T) {
	s := NewServer()
	r := httptest.NewRequest("GET", "/api/languages", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestHandlertranstale(t *testing.T) {
	s := NewServer()
	r := httptest.NewRequest("POST", "/api/translate", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}
