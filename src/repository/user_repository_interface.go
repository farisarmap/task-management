package repository

import (
	"context"
	"database/sql"
	"task-management-be/src/model/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
}
