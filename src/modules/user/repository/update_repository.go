package repository

import (
	"context"
	"database/sql"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) Update(ctx context.Context, tx *sql.Tx, id string, user entity.User) entity.User {
	panic("not implemented") // TODO: Implement
}
