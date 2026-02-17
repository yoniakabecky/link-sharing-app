package services

import (
	"context"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
)

type LinkServices struct {
	repo *repositories.LinkRepository
}

func NewLinkServices(repo *repositories.LinkRepository) *LinkServices {
	return &LinkServices{
		repo: repo,
	}
}

func (s *LinkServices) GetLinksByProfileID(ctx context.Context, prfID int) ([]models.Link, error) {
	return s.repo.GetLinksByProfileID(ctx, prfID)
}

func (s *LinkServices) UpdateLinks(ctx context.Context, prfID int, ls []models.Link) error {
	return s.repo.UpdateLinks(ctx, prfID, ls)
}
