package dbcache

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
)

var ErrDoesNotExist = errors.New("ErrDoesNotExist")

type DBCache struct {
	repo   Repository
	logger *zap.Logger
	l      sync.Mutex
}

func NewDBCache(repo Repository, logger *zap.Logger) *DBCache {
	return &DBCache{
		repo:   repo,
		logger: logger,
	}
}

func (c *DBCache) Get(key string) (interface{}, error) {
	c.l.Lock()
	defer c.l.Unlock()
	res, err := c.repo.Get(context.Background(), key)
	if err != nil {
		return "", err
	}
	if res.ttl.After(time.Now().Truncate(0)) {
		c.logger.Info("Getting value from cache")
		if err != nil {
			return res.data, err
		}
		item := &Item{
			data: res.data,
		}
		return item.data, nil
	}
	if res.ttl.Before(time.Now().Truncate(0)) {
		c.logger.Info("Deleting value from cache due to timeout")
		c.repo.Delete(context.Background(), key)
	}
	return "", ErrDoesNotExist
}

func (c *DBCache) Set(key string, value interface{}, timeout *time.Duration) {
	c.l.Lock()
	defer c.l.Unlock()
	c.logger.Info("Inserting new value to cache")
	convertedValue, ok := value.(string)
	if !ok {
		c.logger.Info("Incorrect conversion")
	}
	if timeout != nil {
		expirationTime := time.Now().Add(*timeout)
		c.repo.Set(context.Background(), key, convertedValue, expirationTime)
		c.logger.Info("Insert Succesful")
	} else {
		c.logger.Info("time duration is nil value won't be cached")
	}

}
