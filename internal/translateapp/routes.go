package translateapp

func (a *Api) routes() {
	a.HandleFunc("/api/languages", a.getLanguages()).Methods("GET")
	a.HandleFunc("/api/languages/{code}", a.getLanguage()).Methods("GET")
	a.HandleFunc("/api/translate", a.translate()).Methods("POST")
}
