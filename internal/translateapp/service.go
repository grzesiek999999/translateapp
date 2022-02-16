package translateapp

import (
	"context"
	"go.uber.org/zap"
)

type Service struct {
	logger     *zap.Logger
	translator LibreTranslator
}

type LibreTranslator interface {
	GetLanguages(ctx context.Context) (*ListLanguage, error)
	Translate(ctx context.Context, word WordToTranslate) (*WordTranslate, error)
}

func NewService(logger *zap.Logger, translator LibreTranslator) *Service {
	service := &Service{
		translator: translator,
	}
	return service
}

func (s *Service) GetLanguages(ctx context.Context) (*Response, error) {
	listLanguages, err := s.translator.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	var response Response
	response.Data = *listLanguages
	response.Code = 200
	response.Message = "success"
	return &response, nil
}

func (s *Service) Translate(ctx context.Context, word WordToTranslate) (*TranslateResponse, error) {
	translation, err := s.translator.Translate(ctx, word)
	if err != nil {
		return nil, err
	}
	var response TranslateResponse
	response.Data = *translation
	response.Code = 200
	response.Message = "success"
	return &response, nil
}

func (s *Service) BatchTranslate(ctx context.Context, word WordToTranslate) (*BatchTranslateResponse, error) {
	translation, err := s.translator.Translate(ctx, word)
	if err != nil {
		return nil, err
	}
	var response BatchTranslateResponse
	response.WordToTranslate = word.Word
	response.WordTranslated = translation.Text

	return &response, nil
}
