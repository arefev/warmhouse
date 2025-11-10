package middleware

import (
	"temperature-api/internal/application"
)

type Middleware struct {
	app *application.App
}

func NewMiddleware(app *application.App) Middleware {
	return Middleware{
		app: app,
	}
}
