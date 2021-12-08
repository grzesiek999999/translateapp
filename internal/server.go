package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*mux.Router
	languages []Language
}

func NewServer() *Server {
	s := &Server{
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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
