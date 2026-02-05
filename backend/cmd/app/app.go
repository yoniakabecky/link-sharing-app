package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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
	errCh := make(chan error, 1)

	go func() {
		log.Println("Server started on", a.server.Addr)
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		// when server fails
		return err
	case <-ctx.Done():
		// when signal is received, shutdown the server gracefully
		return a.Shutdown(ctx)
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server shutdown complete")
	return nil
}
