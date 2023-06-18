package services

import (
	"context"
	"fmt"
	"task-management-be/src/helper"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
	"task-management-be/src/modules/user/model/entity"
	"time"
)

func (service *UserService) Create(ctx context.Context, req request.UserRequest) response.CreateUserResponse {
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
