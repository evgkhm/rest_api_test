package service

import "rest_api_test/internal/repository"

type Auth interface {
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
