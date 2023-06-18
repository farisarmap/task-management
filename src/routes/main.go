package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitializeMainRoute(router *chi.Mux, db *sql.DB) {
	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// userInjection := wires.ProvideUserInjection(db)

	router.Route("/users", func(r chi.Router) {
		// routes.InitializeUserRoute(r, userInjection)
	})
}
