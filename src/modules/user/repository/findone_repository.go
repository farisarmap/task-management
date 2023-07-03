package repository

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (r *UserRepository) FindOne(ctx context.Context, colName string, value string) (domain.User, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		helper.Logger.Warn().Msg("failed to start tx db")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1", colName)
	rows, err := tx.QueryContext(ctx, query, value)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			helper.Logger.Warn().Msg("Failed to create data transaction")
			return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
		}
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}

	err = tx.Commit() // Commit the transaction if the delete operation succeeds
	if err != nil {
		tx.Rollback() // Rollback the transaction if the commit fails
		helper.Logger.Warn().Msg("Rollback the transaction because commit fails")
		return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
	}
	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_At, &user.Updated_At)
		if err != nil {
			helper.Logger.Warn().Msg("failed scan user")
			return domain.User{}, fmt.Errorf(helper.ErrInternalServerError)
		}
	} else {
		helper.Logger.Warn().Msg("User not found")
		return domain.User{}, fmt.Errorf(helper.ErrUserNotFound)
	}

	return user, nil
}
