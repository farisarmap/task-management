package interfaces

import (
	"context"
	"database/sql"
	"task-management-be/src/modules/user/model/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Update(ctx context.Context, tx *sql.Tx, id string, user entity.User) entity.User
	FindById(ctx context.Context, tx *sql.Tx, id string) entity.User
	Delete(ctx context.Context, tx *sql.Tx, id string) bool
	List(ctx context.Context, tx *sql.Tx) []entity.User
}
