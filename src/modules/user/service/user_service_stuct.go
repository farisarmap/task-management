package services

import (
	"task-management-be/src/modules/user/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}
