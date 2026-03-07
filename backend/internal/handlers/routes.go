package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/jwt"
)

type Handlers struct {
	Platform *PlatformHandler
	Profile  *ProfileHandler
	Link     *LinkHandler
	User     *UserHandler
}

func RegisterRoutes(h *Handlers, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	if cfg != nil && cfg.Upload.Dir != "" {
		r.Mount("/uploads", http.FileServer(http.Dir(cfg.Upload.Dir)))
	}

	// Public routes: no token required
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.User.Register)
		r.Post("/login", h.User.Login)
		r.Post("/refresh", h.User.Refresh)
	})
	r.Route("/public", func(r chi.Router) {
		r.Get("/profiles/{id}", h.Profile.GetProfileByID)
	})

	// Protected routes: require valid JWT
	r.Group(func(r chi.Router) {
		r.Use(jwt.Middleware)
		r.Get("/auth/session", h.User.Session)
		r.Route("/platforms", func(r chi.Router) {
			r.Get("/", h.Platform.GetAllPlatforms)
		})
		r.Route("/profiles", func(r chi.Router) {
			r.Get("/", h.Profile.GetProfilesByUserID)
			r.Get("/{id}", h.Profile.GetProfileByID)
			r.Post("/", h.Profile.CreateProfile)
			r.Put("/{id}", h.Profile.UpdateProfile)
			r.Post("/{id}/avatar", h.Profile.UploadAvatar)
			r.Delete("/{id}", h.Profile.DeleteProfile)
		})
		r.Route("/links", func(r chi.Router) {
			r.Get("/{profile_id}", h.Link.GetLinksByProfileID)
			r.Put("/{profile_id}", h.Link.UpdateLinks)
		})
	})

	return r
}
