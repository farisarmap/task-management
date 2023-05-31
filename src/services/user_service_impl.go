package services

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-be/src/helper"
	"task-management-be/src/model/data/request"
	"task-management-be/src/model/data/response"
	"task-management-be/src/model/entity"
	"task-management-be/src/repository"
	"time"
)

type UserService struct {
	Repository repository.IUserRepository
	DB         *sql.DB
}

func NewUserService(repository repository.IUserRepository, db *sql.DB) IUserService {
	return &UserService{
		Repository: repository,
		DB:         db,
	}
}

func (service *UserService) CreateUser(ctx context.Context, req request.UserRequest) response.CreateUserResponse {
	var response response.CreateUserResponse

	hashPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		fmt.Println("failed hash password", err)
		return response
	}

	tx, err := service.DB.Begin()
	if err != nil {
		fmt.Println("failed begin trx database", err)
		return response
	}

	newRequest := entity.User{
		Name:       req.Email,
		Email:      req.Email,
		Password:   hashPassword,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}

	repoResponse := service.Repository.Create(ctx, tx, newRequest)

	response.Email = repoResponse.Email
	response.Name = repoResponse.Name

	return response
}
