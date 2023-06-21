package repository

import (
	"context"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) Update(ctx context.Context, id string, user entity.User) entity.User {
	panic("not implemented") // TODO: Implement
}
