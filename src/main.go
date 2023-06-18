package main

import (
	"fmt"
	"net/http"
	"task-management-be/src/config"
	"task-management-be/src/routes"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	}
	fmt.Println("Welcome to task management api")

	config.DatabaseConnection()

	router := chi.NewRouter()

	routes.InitializeMainRoute(router, config.DB)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
