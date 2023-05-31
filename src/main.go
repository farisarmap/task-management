package main

import (
	"fmt"
	"net/http"
	"task-management-be/src/config"
	"task-management-be/src/controller"
	"task-management-be/src/repository"
	"task-management-be/src/services"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Welcome to task management api")

	config.DatabaseConnection()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, config.DB)
	userController := controller.NewUserController(userService)

	router := chi.NewRouter()

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Post("/users", userController.Create)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
