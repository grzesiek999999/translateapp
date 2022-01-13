package cache

import (
	"context"
	"go.uber.org/zap"
	"time"
	"translateapp/internal/translateapp"
)

type InMemoryCache struct {
	cache      Cacher
	translator translateapp.LibreTranslator
	logger     *zap.Logger
}

func NewInMemoryCache(libreTranslator translateapp.LibreTranslator, cache Cacher, logger *zap.Logger) *InMemoryCache {
	return &InMemoryCache{
		translator: libreTranslator,
		cache:      cache,
		logger:     logger,
	}
}

func (c *InMemoryCache) Translate(ctx context.Context, word translateapp.WordToTranslate) (*translateapp.WordTranslate, error) {
	res, err := c.cache.Get(word.Word)
	if err != nil {
		c.logger.Info("value not in cache, connecting to client")
		response, err := c.translator.Translate(context.Background(), word)
		c.cache.Set(word.Word, response.Text, time.Second*5)
		return response, err
	}
	item := &translateapp.WordTranslate{
		Text: res,
	}

	return item, nil
}

func (c *InMemoryCache) GetLanguages(ctx context.Context) (*translateapp.ListLanguage, error) {
	response, err := c.translator.GetLanguages(context.Background())
	if err != nil {
		return nil, err
	}
	return response, nil
}
