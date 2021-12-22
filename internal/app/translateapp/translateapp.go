package translateapp

import (
	"go.uber.org/zap"
	"translateapp/internal/client"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
)

func Run(logger *zap.Logger) error {
	client := client.NewLibreTranslateClient()
	service := translateapp.NewService(logger, client)
	app := translateapp.NewApp(service, logger)
	srv := server.GetServer(app, logger)

	return server.Run(srv, logger)
}
