package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

type ProfileHandler struct {
	ctx      context.Context
	services *services.ProfileServices
}

func NewProfileHandler(services *services.ProfileServices) *ProfileHandler {
	return &ProfileHandler{
		ctx:      context.Background(),
		services: services,
	}
}

func (h *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.services.GetProfileByID(h.ctx, idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	body := new(models.Profile)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prof := models.Profile{
		UserID:    body.UserID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		AvatarURL: body.AvatarURL,
		Links:     body.Links,
	}

	p, err := h.services.CreateProfile(h.ctx, &prof)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body := new(models.Profile)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prof := models.Profile{
		ID:        idInt,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		AvatarURL: body.AvatarURL,
		Links:     body.Links,
	}

	p, err := h.services.UpdateProfile(h.ctx, &prof)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.DeleteProfile(h.ctx, idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
