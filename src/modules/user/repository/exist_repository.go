package repository

import (
	"context"
	"fmt"
	"task-management-be/src/helper"
)

func (r *UserRepository) Exist(ctx context.Context, colName string, value string) error {
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return fmt.Errorf(helper.ErrInternalServerError)
	}
	var exist bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM users WHERE %s = $1)", colName)
	err = tx.QueryRowContext(ctx, query, value).Scan(&exist)
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

	if !exist {
		return fmt.Errorf(helper.ErrUserNotFound)
	}
	return nil
}
