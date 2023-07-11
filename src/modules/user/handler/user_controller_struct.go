package handler

import (
	interfaces "task-management-be/src/modules/user/interface"
)

type UserHandler struct {
	Service interfaces.IUserService
}

func NewUserHandler(service interfaces.IUserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
