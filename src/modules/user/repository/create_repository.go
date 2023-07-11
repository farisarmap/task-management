package repository

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (repo *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	tx, err := repo.DB.Begin()

	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	insertQuery := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.ExecContext(ctx, insertQuery, user.Name, user.Email, user.Password, user.Created_At, user.Updated_At)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
		}
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback() // Rollback the transaction if the commit fails
		helper.Logger.Warn().Msg("Rollback the transaction because commit fails")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	return user, nil
}
