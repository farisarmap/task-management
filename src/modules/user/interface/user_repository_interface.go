package interfaces

import (
	"context"
	"task-management-be/src/modules/user/model/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, user entity.User) entity.User
	Update(ctx context.Context, id string, user entity.User) entity.User
	FindById(ctx context.Context, id string) entity.User
	Delete(ctx context.Context, id string) bool
	List(ctx context.Context) []entity.User
}
