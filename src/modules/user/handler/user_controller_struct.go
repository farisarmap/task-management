package handler

import (
	services "task-management-be/src/modules/user/service"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}
