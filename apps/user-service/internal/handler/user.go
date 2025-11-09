package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"user-service/internal/application"

	"github.com/go-chi/chi/v5"
)

type user struct {
	app *application.App
}

type UserResponse struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUser(app *application.App) *user {
	return &user{app: app}
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := UserResponse{ID: id}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *user) Update(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
