package handler

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"time"

	"temperature-api/internal/application"
)

type temperature struct {
	app *application.App
}

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
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

	resp := TemperatureResponse{
		Value: float64(tmpr),
		Timestamp: time.Now(),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
