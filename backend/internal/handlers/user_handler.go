package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

type UserHandler struct {
	ctx      context.Context
	services *services.UserServices
}

func NewUserHandler(services *services.UserServices) *UserHandler {
	return &UserHandler{
		ctx:      context.Background(),
		services: services,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	body := new(models.UserAuthInput)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.services.Register(h.ctx, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	body := new(models.UserAuthInput)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.services.Login(h.ctx, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
