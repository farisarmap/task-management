package main

import (
	"net/http"
	"task-management-be/src/config"
	"task-management-be/src/helper"
	"task-management-be/src/routes"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	}

	helper.LoggingFile()
	helper.Logger.Info().Str("module", "main").Msg("Starting the task management application")

	config.DatabaseConnection()

	router := chi.NewRouter()

	routes.InitializeMainRoute(router, config.DB)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		// panic(err)
	}
}
