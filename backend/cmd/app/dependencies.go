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
	linkRepo := repositories.NewLinkRepository(dbConn.GetDB(), pltRepo)
	prfRepo := repositories.NewProfileRepository(dbConn.GetDB(), linkRepo)
	userRepo := repositories.NewUserRepository(dbConn.GetDB())
	rtRepo := repositories.NewRefreshTokenRepository(dbConn.GetDB())

	pltSrv := services.NewPlatformServices(pltRepo)
	prfSrv := services.NewProfileServices(prfRepo)
	linkSrv := services.NewLinkServices(linkRepo)
	userSrv := services.NewUserServices(userRepo, rtRepo)

	pltHdl := handlers.NewPlatformHandler(pltSrv)
	prfHdl := handlers.NewProfileHandler(prfSrv, &cfg.Upload)
	linkHdl := handlers.NewLinkHandler(linkSrv)
	userHdl := handlers.NewUserHandler(userSrv)

	h := &handlers.Handlers{
		Platform: pltHdl,
		Profile:  prfHdl,
		Link:     linkHdl,
		User:     userHdl,
	}

	return &Dependencies{
		DB:       dbConn,
		Handlers: h,
	}, nil
}
