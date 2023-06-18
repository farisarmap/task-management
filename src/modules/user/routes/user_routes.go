package routes

import (
	"task-management-be/src/modules/user/handler"

	"github.com/go-chi/chi/v5"
)

func InitializeUserRoute(router chi.Router, userInjection *handler.UserController) {
	router.Post("/create", userInjection.Create)
}
