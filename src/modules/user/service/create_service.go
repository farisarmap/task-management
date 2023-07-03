package services

import (
	"context"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
	"time"
)

func (service *UserService) Create(ctx context.Context, req domain.UserRequest) (domain.CreateUserResponse, error) {
	var response domain.CreateUserResponse

	hashPassword, err := helper.HashPassword(req.Password)
	helper.ErrorNotNil(err, "Failed to hash password")

	err = service.Repository.Exist(ctx, "email", req.Email)

	if err == nil {
		return response, helper.NewAppError("Email already used")
	}

	newRequest := domain.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashPassword,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}

	repoResponse, err := service.Repository.Create(ctx, newRequest)
	if err != nil {
		return response, helper.NewAppError(err.Error())
	}

	response.Email = repoResponse.Email
	response.Name = repoResponse.Name

	return response, nil
}
