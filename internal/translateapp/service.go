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

func (s *Service) Translate(ctx context.Context, words []WordToTranslate) (*[]TranslateResponse, error) {
	var responseList []TranslateResponse
	if len(words) == 1 {
		translation, err := s.translator.Translate(ctx, words[0])
		if err != nil {
			return nil, err
		}
		response := make([]TranslateResponse, 1, 1)
		response[0].Data = *translation
		response[0].Code = 200
		response[0].Message = "success"
		return &response, nil
	}

	for _, v := range words {
		translation, err := s.translator.Translate(ctx, v)
		if err != nil {
			return nil, err
		}
		var response TranslateResponse
		response.Data = *translation
		response.Code = 200
		response.Message = "success"
		responseList = append(responseList, response)
	}
	return &responseList, nil
}
