package interfaces

import (
	"context"
	"task-management-be/src/modules/user/model/data/request"
	"task-management-be/src/modules/user/model/data/response"
)

type IUserService interface {
	Create(ctx context.Context, req request.UserRequest) response.CreateUserResponse
}
