package main

import (
	"github.com/yoniakabecky/link-sharing-app/backend/db"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/handlers"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/services"
)

type Dependencies struct {
	DB       *db.Database
	Handlers *handlers.Handlers
}

func InitializeDependencies(cfg *config.Config) (*Dependencies, error) {
	dbConn, err := db.NewDatabase(cfg.Database.DSN())
	if err != nil {
		return nil, err
	}

	pltRepo := repositories.NewPlatformRepository(dbConn.GetDB())
	linkRepo := repositories.NewLinkRepository(dbConn.GetDB())
	prfRepo := repositories.NewProfileRepository(dbConn.GetDB(), linkRepo)

	pltSrv := services.NewPlatformServices(pltRepo)
	prfSrv := services.NewProfileServices(prfRepo)
	linkSrv := services.NewLinkServices(linkRepo)

	pltHdl := handlers.NewPlatformHandler(pltSrv)
	prfHdl := handlers.NewProfileHandler(prfSrv)
	linkHdl := handlers.NewLinkHandler(linkSrv)

	h := &handlers.Handlers{
		Platform: pltHdl,
		Profile:  prfHdl,
		Link:     linkHdl,
	}

	return &Dependencies{
		DB:       dbConn,
		Handlers: h,
	}, nil
}
