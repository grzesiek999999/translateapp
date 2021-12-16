package translateapp

//Language struct group data about langauge
type Language struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type WordToTranslate struct {
	Word   string `json:"word"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type WordResponse struct {
	TranslatedWord string
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
