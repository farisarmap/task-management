package services

import (
	"context"
	"strconv"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (us *UserService) Update(ctx context.Context, id string, req domain.UserRequest) error {
	newId, _ := strconv.Atoi(id)
	user := domain.User{
		Id:    newId,
		Name:  req.Name,
		Email: req.Email,
	}
	err := us.Repository.Update(ctx, id, user)
	if err != nil {
		return helper.NewAppError(err.Error())
	}
	return nil
}
