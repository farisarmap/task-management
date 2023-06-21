package services

import (
	"context"
	"task-management-be/src/helper"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
	"task-management-be/src/modules/user/model/entity"
	"time"
)

func (service *UserService) Create(ctx context.Context, req request.UserRequest) response.CreateUserResponse {
	var response response.CreateUserResponse

	hashPassword, err := helper.HashPassword(req.Password)
	helper.ErrorNotNil(err, "Failed to hash password")

	newRequest := entity.User{
		Name:       req.Email,
		Email:      req.Email,
		Password:   hashPassword,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}

	repoResponse := service.Repository.Create(ctx, newRequest)

	response.Email = repoResponse.Email
	response.Name = repoResponse.Name

	return response
}
