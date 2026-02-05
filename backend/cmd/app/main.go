package main

import (
	"context"
	"log"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/handlers"
)

func main() {
	cfg := config.Load()

	deps, err := InitializeDependencies(cfg)
	if err != nil {
		log.Fatalf("failed to initialize dependencies: %v", err)
	}
	defer func() {
		if err := deps.DB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()
	log.Println("Connected to database")

	r := handlers.RegisterRoutes(deps.Handlers)

	app := New(r, cfg.Server.Address)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
