package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"device-service/internal/application"

	"github.com/go-chi/chi/v5"
)

type device struct {
	app *application.App
}

type DeviceResponse struct {
	ID           int    `json:"id"`
	DeviceTypeID string `json:"deviceTypeID"`
	HouseID      string `json:"houseID"`
	SerialNumber string `json:"serialNumber"`
	Status       string `json:"status"`
}

func NewDevice(app *application.App) *device {
	return &device{app: app}
}

func (d *device) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (d *device) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := DeviceResponse{ID: id, Status: "enabled"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *device) Update(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
