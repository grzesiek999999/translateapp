package dbcache

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Connector interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Repo struct {
	connection Connector
	logger     *zap.Logger
}

func NewRepo(conn Connector, logger *zap.Logger) *Repo {
	return &Repo{
		connection: conn,
		logger:     logger,
	}
}

func (r Repo) Get(ctx context.Context, key string) (*Item, error) {
	var item Item
	err := r.connection.QueryRow(ctx, "SELECT value, timeout from cache WHERE key=$1", key).Scan(&item.data, &item.ttl)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r Repo) Set(ctx context.Context, key string, value string, timeout time.Time) error {
	_, err := r.connection.Exec(ctx, "INSERT INTO cache(key, value, timeout) VALUES($1, $2, $3)", key, value, timeout)
	if err != nil {
		return err
	}
	r.logger.Info("Data Inserted")
	return nil
}

func (r Repo) Delete(ctx context.Context, key string) error {
	_, err := r.connection.Exec(ctx, "DELETE FROM cache WHERE key = $1", key)
	if err != nil {
		return err
	}
	return nil
}
