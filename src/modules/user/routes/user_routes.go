package routes

import (
	"task-management-be/src/modules/user/handler"

	"github.com/go-chi/chi/v5"
)

func InitializeUserRoute(router chi.Router, userInjection *handler.UserHandler) {
	router.Post("/create", userInjection.Create)
	router.Get("/detail/{id}", userInjection.FindById)
}
