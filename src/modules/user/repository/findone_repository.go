package repository

import (
	"context"
	"database/sql"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) FindById(ctx context.Context, tx *sql.Tx, id string) entity.User {
	panic("not implemented") // TODO: Implement
}
