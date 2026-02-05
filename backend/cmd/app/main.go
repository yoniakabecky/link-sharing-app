package main

import (
	"context"
	"log"

	"github.com/yoniakabecky/link-sharing-app/backend/db"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/handlers"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

func main() {
	cfg := config.Load()

	// Initialize database
	dbConn, err := db.NewDatabase(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()
	log.Println("Connected to database")

	// Initialize dependencies
	prepo := repositories.NewPlatformRepository(dbConn.GetDB())
	plsrv := services.NewPlatformServices(prepo)
	plhdl := handlers.NewPlatformHandler(plsrv)
	prrepo := repositories.NewProfileRepository(dbConn.GetDB())
	prsrv := services.NewProfileServices(prrepo)
	prhdl := handlers.NewProfileHandler(prsrv)

	// Register routes
	h := &handlers.Handlers{
		Platform: plhdl,
		Profile:  prhdl,
	}
	router := handlers.RegisterRoutes(h)

	// Start server
	addr := ":8080"
	app := New(router, addr)
	log.Println("Starting server on port", addr)
	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
