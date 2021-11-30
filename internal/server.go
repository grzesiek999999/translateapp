package internal

import "github.com/gorilla/mux"

type server struct{
	*mux.Router

}

func NewServer() *server{
	s := &server{
		Router:	mux.NewRouter(),
	}
	s.routes()
	return s
}



