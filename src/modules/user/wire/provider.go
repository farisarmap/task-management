package wires

import (
	"task-management-be/src/modules/user/handler"
	"task-management-be/src/modules/user/repository"
	services "task-management-be/src/modules/user/service"

	"github.com/google/wire"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	handler.NewUserHandler,
	services.NewUserService,
	repository.NewUserRepository,

	// wire.Bind(new(interfaces.IUserHandler), new(*handler.UserHandler)),
	// wire.Bind(new(interfaces.IUserService), new(*services.UserService)),
	// wire.Bind(new(interfaces.IUserRepository), new(*repository.UserRepository)),
)
