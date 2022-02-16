package translateapp

func (a *App) routes() {
	a.HandleFunc("/api/languages", a.getLanguages()).Methods("GET")
	a.HandleFunc("/api/translate", a.translate()).Methods("POST")
	a.HandleFunc("/api/batchtranslate", a.BatchTranslate()).Methods("POST")
}
