package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	Platform *PlatformHandler
	Profile  *ProfileHandler
}

func RegisterRoutes(h *Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/platforms", func(r chi.Router) {
		r.Get("/", h.Platform.GetAllPlatforms)
	})

	r.Route("/profiles", func(r chi.Router) {
		r.Get("/{id}", h.Profile.GetProfileByID)
		r.Post("/", h.Profile.CreateProfile)
		r.Put("/{id}", h.Profile.UpdateProfile)
		r.Delete("/{id}", h.Profile.DeleteProfile)
	})

	return r
}
