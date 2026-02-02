package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers struct {
	Platform *PlatformHandler
}

func RegisterRoutes(h *Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/platforms", func(r chi.Router) {
		r.Get("/", h.Platform.GetAllPlatforms)
	})

	return r
}
