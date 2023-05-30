package main

import (
	"fmt"
	"net/http"
	"task-management-be/src/config"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Welcome to task management api")

	config.DatabaseConnection()

	router := chi.NewRouter()

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
