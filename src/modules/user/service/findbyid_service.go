package services

import (
	"context"
	"fmt"
	"task-management-be/src/modules/user/model/data/response"
)

func (service *UserService) FindById(ctx context.Context, id string) response.FindByIdUserResponse {
	user := service.Repository.FindById(ctx, id)
	fmt.Println(user.Password, "<< pass")
	response := response.FindByIdUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	return response
}
