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

type LinkHandler struct {
	ctx      context.Context
	services *services.LinkServices
}

func NewLinkHandler(services *services.LinkServices) *LinkHandler {
	return &LinkHandler{
		ctx:      context.Background(),
		services: services,
	}
}

func (h *LinkHandler) GetLinksByProfileID(w http.ResponseWriter, r *http.Request) {
	prfID := chi.URLParam(r, "profile_id")
	prfIDInt, err := strconv.Atoi(prfID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links, err := h.services.GetLinksByProfileID(h.ctx, prfIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

func (h *LinkHandler) UpdateLinks(w http.ResponseWriter, r *http.Request) {
	prfIDStr := chi.URLParam(r, "profile_id")
	prfID, err := strconv.Atoi(prfIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var links []models.Link
	if err := json.NewDecoder(r.Body).Decode(&links); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range links {
		links[i].ProfileID = prfID
	}

	if err := h.services.UpdateLinks(h.ctx, prfID, links); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
