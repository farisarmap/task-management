package services

import (
	"context"
	"task-management-be/src/model/data/request"
	"task-management-be/src/model/data/response"
)

type IUserService interface {
	CreateUser(ctx context.Context, req request.UserRequest) response.CreateUserResponse
}
