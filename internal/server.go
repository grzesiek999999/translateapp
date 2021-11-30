package internal

import "github.com/gorilla/mux"

type server struct {
	*mux.Router
	languages []Language
}

func NewServer() *server {
	s := &server{
		Router: mux.NewRouter(),
		languages: []Language{
			{
				Code:     "pl",
				Language: "polish",
			},
			{
				Code:     "en",
				Language: "english",
			},
		},
	}
	s.routes()
	return s
}
