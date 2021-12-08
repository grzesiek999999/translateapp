package internal

func (s *Server) routes() {
	s.HandleFunc("/api/languages", s.getLanguages()).Methods("GET")
	s.HandleFunc("/api/languages/{code}", s.getLanguage()).Methods("GET")
	s.HandleFunc("/api/translate", s.translate()).Methods("POST")
}
