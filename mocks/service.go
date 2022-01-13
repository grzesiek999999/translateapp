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

func (s *Servicer) WordTranslate(ctx context.Context) (*translateapp.Response, error) {
	args := s.Called(ctx)
	return args.Get(0).(*translateapp.Response), args.Error(1)
}
