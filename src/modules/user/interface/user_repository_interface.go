package interfaces

import (
	"context"
	"task-management-be/src/domain"
)

type IUserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, id string, user domain.User) error
	FindById(ctx context.Context, id string) (domain.User, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.User, error)
	Exist(ctx context.Context, colName string, value string) error
	FindOne(ctx context.Context, colName string, value string) (domain.User, error)
}
