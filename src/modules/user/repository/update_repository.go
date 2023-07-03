package repository

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (repo *UserRepository) Update(ctx context.Context, id string, user domain.User) error {
	tx, err := repo.DB.BeginTx(ctx, nil)

	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return fmt.Errorf(helper.ErrInternalServerError)
	}

	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`

	_, err = tx.ExecContext(ctx, query, user.Name, user.Email, id)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return fmt.Errorf(helper.ErrInternalServerError)
		}
		return fmt.Errorf(helper.ErrInternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback() // Rollback the transaction if the commit fails
		helper.Logger.Warn().Msg("Rollback the transaction because commit fails")
		return fmt.Errorf(helper.ErrInternalServerError)
	}
	return nil
}
