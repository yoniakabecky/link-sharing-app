package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	server *http.Server
}

// New creates a new App instance with the provided router and address
func New(router http.Handler, addr string) *App {
	app := &App{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
	return app
}

// Start starts the HTTP server and blocks until the context is cancelled
func (a *App) Start(ctx context.Context) error {
	log.Println("Server started on", a.server.Addr)
	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
