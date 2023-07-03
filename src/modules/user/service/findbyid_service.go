package services

import (
	"context"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (service *UserService) FindById(ctx context.Context, id string) (domain.FindByIdUserResponse, error) {
	user, err := service.Repository.FindById(ctx, id)
	if err != nil {
		return domain.FindByIdUserResponse{}, helper.NewAppError(err.Error())
	}

	response := domain.FindByIdUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return response, nil
}
