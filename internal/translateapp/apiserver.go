package translateapp

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	logger *zap.Logger
	*mux.Router
	Service Servicer
}

type Servicer interface {
	GetLanguages(context.Context) (*Response, error)
	Translate(ctx context.Context, word WordToTranslate) (*TranslateResponse, error)
	BatchTranslate(ctx context.Context, word WordToTranslate) (*BatchTranslateResponse, error)
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
