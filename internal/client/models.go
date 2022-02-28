package client

import "fmt"

type Language struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type ListLanguage struct {
	Languages []Language
}

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ListLanguage `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

type WordTranslate struct {
	Text string `json:"translatedText"`
}

type TranslateResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    WordTranslate `json:"data"`
}

type WordToTranslate struct {
	Word   string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type Errors struct {
	Error string `json:"error"`
}

func (r Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", r.Code, r.Message)
}
