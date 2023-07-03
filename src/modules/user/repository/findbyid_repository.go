package repository

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (repo *UserRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)

	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	query := `SELECT * FROM users WHERE id=$1`

	rows, err := tx.QueryContext(ctx, query, id)

	if err != nil {
		helper.Logger.Warn().Msg("failed query ctx user")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}
	defer rows.Close()

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("failed to create data transaction")
			return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
		}
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Created_At, &user.Updated_At)

		if err != nil {
			helper.Logger.Warn().Msg("failed scan user")
			return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
		}
	} else {
		helper.Logger.Warn().Msg(helper.ErrUserNotFound)
		return domain.User{}, fmt.Errorf(helper.ErrUserNotFound)
	}

	return user, nil
}
