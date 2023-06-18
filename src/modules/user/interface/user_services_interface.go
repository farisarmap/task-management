package interfaces

import (
	"context"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
)

type ICreateUserService interface {
	Create(ctx context.Context, req request.UserRequest) response.CreateUserResponse
}

type IUserService interface {
	ICreateUserService
}
