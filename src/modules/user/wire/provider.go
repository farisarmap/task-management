package wires

import (
	"task-management-be/src/modules/user/handler"
	interfaces "task-management-be/src/modules/user/interface"
	"task-management-be/src/modules/user/repository"
	services "task-management-be/src/modules/user/service"

	"github.com/google/wire"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	handler.NewUserController,
	services.NewUserService,
	repository.NewUserRepository,

	wire.Bind(new(interfaces.IUserController), new(*handler.UserController)),
	wire.Bind(new(interfaces.IUserService), new(*services.UserService)),
	wire.Bind(new(interfaces.IUserRepository), new(*repository.UserRepository)),
)
