package handler

import (
	"math/rand/v2"
	"net/http"
	"strconv"

	"temperature-api/internal/application"
)

type temperature struct {
	app *application.App
}

func NewTemperature(app *application.App) *temperature {
	return &temperature{app: app}
}

func (t *temperature) Get(w http.ResponseWriter, r *http.Request) {
	const max int = 50
	tmpr := rand.IntN(max)
	if rand.IntN(2) > 0 {
		tmpr *= -1
	}

	w.Write([]byte(strconv.Itoa(tmpr)))
	w.WriteHeader(http.StatusOK)
}
