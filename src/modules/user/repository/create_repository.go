package repository

import (
	"context"
	"task-management-be/src/helper"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) Create(ctx context.Context, user entity.User) entity.User {
	tx, err := repo.DB.BeginTx(ctx, nil)
	helper.ErrorNotNil(err, "Failed start tx db")

	insertQuery := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.ExecContext(ctx, insertQuery, user.Name, user.Email, user.Password, user.Created_At, user.Updated_At)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return entity.User{}
		}
		return entity.User{}
	}
	tx.Commit()

	return user
}
