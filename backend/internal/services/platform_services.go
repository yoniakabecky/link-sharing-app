package services

import (
	"context"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
)

type PlatformServices struct {
	repo *repositories.PlatformRepository
}

func NewPlatformServices(repo *repositories.PlatformRepository) *PlatformServices {
	return &PlatformServices{
		repo: repo,
	}
}

func (s *PlatformServices) GetAll(ctx context.Context) ([]models.Platform, error) {
	return s.repo.GetAll(ctx)
}

func (s *PlatformServices) GetPlatformByID(ctx context.Context, id int) (*models.Platform, error) {
	return s.repo.GetPlatformByID(ctx, id)
}
