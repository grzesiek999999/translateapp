package translateapp

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type App struct {
	*mux.Router

	logger  *zap.Logger
	Service Servicer
}

type Servicer interface {
	GetLanguages(context.Context) (*Response, error)
	Translate(ctx context.Context, word WordToTranslate) (*TranslateResponse, error)
}

func NewApp(service Servicer, logger *zap.Logger) *App {
	a := &App{
		Service: service,
		logger:  logger,
		Router:  mux.NewRouter(),
	}
	a.routes()
	return a
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}
