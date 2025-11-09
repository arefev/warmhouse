package router

import (
	"user-service/internal/application"
	"user-service/internal/handler"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func New(app *application.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chi_middleware.Logger)

	app.Log.Info("Server started")

	userHandler := handler.NewUser(app)

	r.Post("/users", userHandler.Create)
	r.Get("/users/{id}", userHandler.Get)
	r.Put("/users/{id}", userHandler.Update)

	return r
}
