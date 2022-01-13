package translateapp

import (
	"go.uber.org/zap"
	"translateapp/internal/cache"
	"translateapp/internal/client"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
)

func Run(logger *zap.Logger) error {
	cacher := cache.NewCache(logger)
	translator := client.NewLibreTranslateClient(logger)
	cashedTranslator := cache.NewInMemoryCache(translator, cacher, logger)
	service := translateapp.NewService(logger, cashedTranslator)
	app := translateapp.NewApp(service, logger)
	srv := server.GetServer(app, logger)
	return server.Run(srv, logger)
}
