package router

import (
	"device-service/internal/application"
	"device-service/internal/handler"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func New(app *application.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chi_middleware.Logger)

	app.Log.Info("Server started")

	deviceHandler := handler.NewDevice(app)

	r.Post("/devices", deviceHandler.Create)
	r.Get("/devices/{id}", deviceHandler.Get)
	r.Put("/devices/{id}", deviceHandler.Update)

	return r
}
