package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"translateapp/internal/translateapp"
)

type Servicer struct {
	mock.Mock
}

func (s *Servicer) GetLanguages(ctx context.Context) (*translateapp.Response, error) {
	args := s.Called(ctx)
	return args.Get(0).(*translateapp.Response), args.Error(1)
}

func (s *Servicer) Translate(ctx context.Context, translate translateapp.WordToTranslate) (*translateapp.TranslateResponse, error) {
	args := s.Called(ctx)
	return args.Get(0).(*translateapp.TranslateResponse), args.Error(1)
}

func (s *Servicer) BatchTranslate(ctx context.Context, word translateapp.WordToTranslate) (*translateapp.BatchTranslateResponse, error) {
	args := s.Called(ctx, word)
	return args.Get(0).(*translateapp.BatchTranslateResponse), args.Error(1)
}
