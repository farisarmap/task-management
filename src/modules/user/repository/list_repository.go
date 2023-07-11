package repository

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (repo *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return []domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	insertQuery := `select * from users`
	rows, err := tx.QueryContext(ctx, insertQuery)
	if err != nil {
		helper.Logger.Warn().Msg("failed to start query context")
		return []domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	var users []domain.User

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Created_At, &user.Updated_At)
		if err != nil {
			helper.Logger.Info().Msg("No data user")
			return []domain.User{}, fmt.Errorf(helper.ErrUserNotFound)
		}

		users = append(users, user)
	}
	return users, nil
}
