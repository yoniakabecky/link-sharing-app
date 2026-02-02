package main

import (
	"context"
	"log"

	"github.com/yoniakabecky/link-sharing-app/backend/db"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/handlers"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

func main() {
	// Initialize database
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()
	log.Println("Connected to database")

	// Initialize dependencies
	prepo := repositories.NewPlatformRepository(dbConn.GetDB())
	psrv := services.NewPlatformServices(prepo)
	phdl := handlers.NewPlatformHandler(psrv)

	// Register routes
	h := &handlers.Handlers{
		Platform: phdl,
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
