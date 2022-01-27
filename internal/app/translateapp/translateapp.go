package translateapp

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"log"
	"translateapp/internal/client"
	dbCache2 "translateapp/internal/dbcache"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
)

func Run(logger *zap.Logger) error {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@db:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	repo := dbCache2.NewRepo(conn, logger)
	dbCacher := dbCache2.NewDBCache(repo, logger)
	translator := client.NewLibreTranslateClient(logger)
	cashedTranslator := translateapp.NewCacheWrapper(translator, dbCacher, logger)
	service := translateapp.NewService(logger, cashedTranslator)
	app := translateapp.NewApp(service, logger)
	srv := server.GetServer(app, logger)
	return server.Run(srv, logger)
}
