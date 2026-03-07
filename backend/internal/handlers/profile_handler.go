package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/jwt"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

const (
	avatarFormField = "avatar"
	avatarMaxBytes  = 2 * 1024 * 1024 // 2MB fallback if config not set
)

var allowedAvatarTypes = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/bmp":  ".bmp",
}

type ProfileHandler struct {
	ctx       context.Context
	services  *services.ProfileServices
	uploadCfg *config.UploadConfig
}

func NewProfileHandler(services *services.ProfileServices, uploadCfg *config.UploadConfig) *ProfileHandler {
	return &ProfileHandler{
		ctx:       context.Background(),
		services:  services,
		uploadCfg: uploadCfg,
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

func (h *ProfileHandler) GetProfilesByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(jwt.UserCtxKey).(string)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profiles, err := h.services.GetProfilesByUserID(h.ctx, userIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profiles)
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(jwt.UserCtxKey).(string)
	userIDInt, err := strconv.Atoi(userID)
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
		UserID:    userIDInt,
		Nickname:  body.Nickname,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		AvatarURL: body.AvatarURL,
		Links:     nil,
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

	existing, err := h.services.GetProfileByID(h.ctx, idInt)
	if err != nil || existing == nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}
	avatarURL := body.AvatarURL
	if avatarURL == "" {
		avatarURL = existing.AvatarURL
	}

	prof := models.Profile{
		ID:        idInt,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		AvatarURL: avatarURL,
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

func (h *ProfileHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(jwt.UserCtxKey).(string)
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := h.services.GetProfileByID(h.ctx, idInt)
	if err != nil || profile == nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}
	if profile.UserID != userIDInt {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	if err := r.ParseMultipartForm(avatarMaxBytes); err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile(avatarFormField)
	if err != nil {
		http.Error(w, "missing or invalid avatar file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")
	ext, ok := allowedAvatarTypes[contentType]
	if !ok {
		http.Error(w, "invalid file type: use PNG, JPG, or BMP", http.StatusBadRequest)
		return
	}

	maxBytes := h.uploadCfg.MaxBytes
	if maxBytes <= 0 {
		maxBytes = avatarMaxBytes
	}
	if header.Size > maxBytes {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	uploadDir := h.uploadCfg.Dir
	if uploadDir == "" {
		uploadDir = "uploads"
	}
	avatarsDir := filepath.Join(uploadDir, "avatars")
	if err := os.MkdirAll(avatarsDir, 0755); err != nil {
		http.Error(w, "failed to create upload directory", http.StatusInternalServerError)
		return
	}

	filename := strconv.Itoa(idInt) + ext
	dstPath := filepath.Join(avatarsDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(dstPath)
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}

	baseURL := strings.TrimSuffix(h.uploadCfg.BaseURL, "/")
	avatarURL := baseURL + "/uploads/avatars/" + filename

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"avatar_url": avatarURL})
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
