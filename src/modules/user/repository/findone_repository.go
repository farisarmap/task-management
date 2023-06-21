package repository

import (
	"context"
	"task-management-be/src/helper"
	"task-management-be/src/modules/user/model/entity"
)

func (repo *UserRepository) FindById(ctx context.Context, id string) entity.User {
	tx, err := repo.DB.BeginTx(ctx, nil)
	helper.ErrorNotNil(err, "Failed start tx db")

	query := `SELECT * FROM users WHERE id=$1`

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorNotNil(err, "Failed query ctx user")
	defer rows.Close()

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return entity.User{}
		}
		return entity.User{}
	}

	tx.Commit()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Created_At, &user.Updated_At)
		helper.ErrorNotNil(err, "Failed scan user")
	} else {
		helper.Logger.Info().Msg("No data user available")
	}

	return user
}
