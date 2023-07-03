package services

import (
	"context"
	"task-management-be/src/helper"
)

func (us *UserService) Delete(ctx context.Context, id string) error {
	err := us.Repository.Delete(ctx, id)
	if err != nil {
		return helper.NewAppError(err.Error())
	}

	return nil
}
