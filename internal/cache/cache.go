package cache

import (
	"errors"
	"go.uber.org/zap"
	"time"
)

var ErrDoesNotExist = errors.New("DoesNotExist")

type Cache struct {
	logger *zap.Logger
	m      map[string]*Item
}

type Cacher interface {
	Set(key string, value interface{}, timeout *time.Duration)
	Get(key string) (interface{}, error)
}

// Set adds value to the cache under specified key and timeout
//Passing in -1 for timeout will cache the value forever
// A timeout of 0 won't cache the value
func (c *Cache) Set(key string, data string, timeout time.Duration) {
	if timeout > 0 || timeout == -1 {
		item, exists := c.m[key]
		if exists {
			c.logger.Info("Value added to cache")
			item.ttl = time.Now().Add(timeout)
		} else {
			c.m[key] = &Item{
				data:  data,
				ttl:   time.Now().Add(timeout),
				timer: timeout,
			}
			c.logger.Info("Value added to cache")
		}
	}
}

func (c *Cache) Get(key string) (string, error) {
	item, exists := c.m[key]
	if exists && (item.ttl.After(time.Now())) {
		item.ttl = time.Now().Add(item.timer)
		c.logger.Info("Value get from cache")
		return item.data, nil
	} else if exists && item.timer == -1 {
		c.logger.Info("Value get from cache")
		return item.data, nil
	}
	return "nil", ErrDoesNotExist

}

func NewCache(logger *zap.Logger) *Cache {
	return &Cache{
		m:      make(map[string]*Item),
		logger: logger,
	}
}
