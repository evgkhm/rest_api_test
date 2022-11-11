package service

import (
	"rest_api_test"
	"rest_api_test/internal/repository"
)

type CreateService struct {
	repo repository.Creating
}

func NewCreateService(repo repository.Creating) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateUser(user rest_api_test.User) error {
	return s.repo.CreateUser(user)
}
