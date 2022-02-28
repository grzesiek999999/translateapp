package dbcache

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"time"
)

var ErrDoesNotExist = errors.New("ErrDoesNotExist")

type DBCache struct {
	repo   Repository
	logger *zap.Logger
}

func NewDBCache(repo Repository, logger *zap.Logger) *DBCache {
	return &DBCache{
		repo:   repo,
		logger: logger,
	}
}

func (c *DBCache) Get(key string) (interface{}, error) {
	res, err := c.repo.Get(context.Background(), key)
	if err != nil {
		return "", err
	}
	if res.ttl.After(time.Now().UTC()) {
		c.logger.Info("Getting value from cache")
		if err != nil {
			return res.data, err
		}
		item := &Item{
			data: res.data,
		}
		return item.data, nil
	} else if res.ttl.Before(time.Now().UTC()) {
		c.logger.Info("Deleting value from cache due to timeout")
		c.repo.Delete(context.Background(), key)
		return "", ErrDoesNotExist
	}
	return "", ErrDoesNotExist
}

func (c *DBCache) Set(key string, value interface{}, timeout *time.Duration) {
	c.logger.Info("Inserting new value to cache")
	convertedValue, ok := value.(string)
	if !ok {
		c.logger.Info("Incorrect conversion")
	}
	if timeout != nil {
		expirationTime := time.Now().UTC().Add(*timeout)
		c.repo.Set(context.Background(), key, convertedValue, expirationTime)
		c.logger.Info("Insert Succesful")
	} else {
		c.logger.Info("time duration is nil value won't be cached")
	}

}
