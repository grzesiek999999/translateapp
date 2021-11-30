package internal

type Language struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type WordToTranslate struct {
	Word   string `json:"word"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type WordResponse struct {
	TranslatedWord string
}
