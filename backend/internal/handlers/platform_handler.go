package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

type PlatformHandler struct {
	ctx      context.Context
	services *services.PlatformServices
}

func NewPlatformHandler(services *services.PlatformServices) *PlatformHandler {
	return &PlatformHandler{
		ctx:      context.Background(),
		services: services,
	}
}

func (h *PlatformHandler) GetAllPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms, err := h.services.GetAll(h.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []models.Platform
	for _, p := range platforms {
		res = append(res, p)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
