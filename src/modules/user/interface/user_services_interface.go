package interfaces

import (
	"context"
	"task-management-be/src/domain"
)

type IUserService interface {
	Create(ctx context.Context, req domain.UserRequest) (domain.CreateUserResponse, error)
	FindById(ctx context.Context, id string) (domain.FindByIdUserResponse, error)
	List(ctx context.Context) ([]domain.CreateUserResponse, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, req domain.UserRequest) error
	Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, error)
}
