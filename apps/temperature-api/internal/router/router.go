package router

import (
	"temperature-api/internal/application"
	"temperature-api/internal/handler"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func New(app *application.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chi_middleware.Logger)

	app.Log.Info("Server started")

	tmprHandler := handler.NewTemperature(app)

	r.Get("/temperature", tmprHandler.Get)

	return r
}
