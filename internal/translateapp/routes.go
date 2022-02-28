package translateapp

func (a *App) routes() {
	a.HandleFunc("/api/languages", a.GetLanguages()).Methods("GET")
	a.HandleFunc("/api/translate", a.Translate()).Methods("POST")
	a.HandleFunc("/api/batchtranslate", a.BatchTranslate()).Methods("POST")
}
