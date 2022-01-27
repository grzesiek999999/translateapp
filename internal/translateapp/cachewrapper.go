package translateapp

import (
	"context"
	"go.uber.org/zap"
	"time"
	"translateapp/internal/cache"
)

type CacheWrapper struct {
	cache      cache.Cacher
	translator LibreTranslator
	logger     *zap.Logger
}

func NewCacheWrapper(libreTranslator LibreTranslator, cache cache.Cacher, logger *zap.Logger) *CacheWrapper {
	return &CacheWrapper{
		translator: libreTranslator,
		cache:      cache,
		logger:     logger,
	}
}

func (c *CacheWrapper) Translate(ctx context.Context, word WordToTranslate) (*WordTranslate, error) {
	res, err := c.cache.Get(word.Word)
	if err != nil {
		c.logger.Info("value not in cache, connecting to client")
		response, err := c.translator.Translate(context.Background(), word)
		duration := time.Second * 5
		c.cache.Set(word.Word, response.Text, &duration)
		return response, err
	}
	convertedResponse, ok := res.(string)
	if !ok {
		c.logger.Info("Incorrect Conversion")
	}
	item := &WordTranslate{
		Text: convertedResponse,
	}
	c.logger.Info("Value get from cache")
	return item, nil
}

func (c *CacheWrapper) GetLanguages(ctx context.Context) (*ListLanguage, error) {
	response, err := c.translator.GetLanguages(context.Background())
	if err != nil {
		return nil, err
	}
	return response, nil
}
