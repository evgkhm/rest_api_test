package service

import (
	"rest_api_test"
	"rest_api_test/internal/repository"
)

type Creating interface {
	CreateUser(user rest_api_test.User) error
}

type Service struct {
	Creating
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Creating: NewCreateService(repos.Creating),
	}
}
