package translateapp

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Api struct {
	logger *zap.Logger
	*mux.Router
	languages []Language
	Service   *Service
}

func NewApi(logger *zap.Logger) *Api {
	a := &Api{
		Service: NewService(),
		logger:  logger,
		Router:  mux.NewRouter(),
	}
	a.routes()
	return a
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}
