package main

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
	addr   string
}

// New creates a new App instance with the provided router and address
func New(router http.Handler, addr string) *App {
	app := &App{
		router: router,
		addr:   addr,
	}
	return app
}

// Start starts the HTTP server and blocks until the context is cancelled
func (a *App) Start(ctx context.Context) error {
	s := &http.Server{
		Addr:    a.addr,
		Handler: a.router,
	}

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
