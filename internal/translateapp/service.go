package translateapp

import (
	"context"
	"log"
)

type Service struct {
	client *LibreTranslateClient
}

func NewService() *Service {
	service := &Service{
		client: NewLibreTranslateClient(),
	}
	return service
}

func (s *Service) GetLanguages(ctx context.Context) ([]Language, error) {
	res, err := s.client.GetLanguages(ctx)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	return res, nil
}
