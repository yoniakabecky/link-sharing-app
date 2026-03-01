package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/jwt"
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

func (h *UserHandler) authResponse(w http.ResponseWriter, status int, accessToken, refreshToken string, user *models.ResponseUser) {
	cfg := config.Load()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    cfg.JWT.Exp,
		"user":          user,
	})
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

	accessToken, refreshToken, err := h.services.IssueTokens(h.ctx, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.authResponse(w, http.StatusCreated, accessToken, refreshToken, u)
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

	accessToken, refreshToken, err := h.services.IssueTokens(h.ctx, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.authResponse(w, http.StatusOK, accessToken, refreshToken, u)
}

func (h *UserHandler) Session(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(jwt.UserCtxKey)
	if userIDVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var userID int
	switch v := userIDVal.(type) {
	case string:
		if _, err := fmt.Sscanf(v, "%d", &userID); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	case float64:
		userID = int(v)
	default:
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := h.services.GetUserByID(h.ctx, userID)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"user": user})
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	body := struct {
		RefreshToken string `json:"refresh_token"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.RefreshToken == "" {
		http.Error(w, "refresh_token required", http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, user, err := h.services.ValidateRefreshAndIssue(h.ctx, body.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	h.authResponse(w, http.StatusOK, accessToken, refreshToken, user)
}
