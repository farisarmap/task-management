package wires

import (
	"task-management-be/src/modules/user/handler"
	"task-management-be/src/modules/user/repository"
	services "task-management-be/src/modules/user/service"

	"github.com/google/wire"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(interfaces.IUserController), new(*handler.UserController)),
	services.NewUserService,
	wire.Bind(new(interfaces.IUserService), new(services.UserService)),
	services.NewCreateUserService,
	wire.Bind(new(interfaces.ICreateUserService), new(*services.CreateUserService)),
	repository.NewUserRepository,
	wire.Bind(new(interfaces.IUserRepository), new(*repository.UserRepository)),
)

// func test(db *sql.DB) {
// 	userRepository := repository.NewUserRepository()
// 	userCreateService := services.NewCreateUserService(userRepository, db)
// 	userService := services.NewUserService(userCreateService)
// 	userHandler := handler.NewUserController(userService)
// }
