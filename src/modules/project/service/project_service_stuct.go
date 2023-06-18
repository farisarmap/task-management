package services

import (
	"database/sql"
	"task-management-be/src/modules/user/repository"
)

type UserService struct {
	Repository *repository.UserRepository
	DB         *sql.DB
}

func NewUserService(repository *repository.UserRepository, db *sql.DB) *UserService {
	return &UserService{
		Repository: repository,
		DB:         db,
	}
}
