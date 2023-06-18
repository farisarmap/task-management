package repository

import (
	"context"
	"database/sql"
	"fmt"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	insertQuery := `insert into users("name", "email", "password", "created_at", "updated_at") values ($1, $2, $3, $4, $5)`

	_, err := tx.ExecContext(ctx, insertQuery, user.Name, user.Email, user.Password, user.Created_At, user.Updated_At)

	if err != nil {
		fmt.Println("error create user:", err)
		return entity.User{}
	}

	return user
}
