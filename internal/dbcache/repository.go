package dbcache

import (
	"context"
	"time"
)

type Repository interface {
	Get(ctx context.Context, key string) (*Item, error)
	Set(ctx context.Context, key string, value string, timeout time.Time) error
	Delete(ctx context.Context, key string) error
}

type Item struct {
	data string
	ttl  time.Time
}
