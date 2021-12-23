package mocks

import (
	"context"
	"translateapp/internal/translateapp"

	"github.com/stretchr/testify/mock"
)

type LibreTranslator struct {
	mock.Mock
}

func (t *LibreTranslator) GetLanguages(ctx context.Context) (*translateapp.ListLanguage, error) {
	args := t.Called(ctx)
	return args.Get(0).(*translateapp.ListLanguage), args.Error(1)
}

func (t *LibreTranslator) Translate(ctx context.Context, word translateapp.WordToTranslate) (*translateapp.WordTranslate, error) {
	args := t.Called(ctx, word)
	return args.Get(0).(*translateapp.WordTranslate), args.Error(1)
}
