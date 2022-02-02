package translateapp

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"translateapp/internal/client"
	"translateapp/internal/dbcache"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
)

func Run(logger *zap.Logger) error {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/postgres")
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	repo := dbcache.NewRepo(conn, logger)
	dbCacher := dbcache.NewDBCache(repo, logger)
	translator := client.NewLibreTranslateClient(logger)
	cashedTranslator := translateapp.NewCacheWrapper(translator, dbCacher, logger)
	service := translateapp.NewService(logger, cashedTranslator)
	app := translateapp.NewApp(service, logger)
	srv := server.GetServer(app, logger)
	return server.Run(srv, logger)
}
