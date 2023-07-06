package services

import (
	interfaces "task-management-be/src/modules/user/interface"
)

type UserService struct {
	Repository interfaces.IUserRepository
}

func NewUserService(repository interfaces.IUserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}
