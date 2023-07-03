package repository

import (
	"context"
	"fmt"
	"task-management-be/src/helper"
)

func (repo *UserRepository) Delete(ctx context.Context, id string) error {
	tx, err := repo.DB.BeginTx(ctx, nil)

	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return fmt.Errorf(helper.ErrInternalServerError)
	}

	query := `DELETE FROM users WHERE id = ?`
	_, err = tx.ExecContext(ctx, query, id)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return fmt.Errorf(helper.ErrInternalServerError)
		}
		return fmt.Errorf(helper.ErrInternalServerError)
	}

	err = tx.Commit() // Commit the transaction if the delete operation succeeds
	if err != nil {
		tx.Rollback() // Rollback the transaction if the commit fails
		helper.Logger.Warn().Msg("Rollback the transaction because commit fails")
		return fmt.Errorf(helper.ErrInternalServerError)
	}

	return nil
}
