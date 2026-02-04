package services

import (
	"context"

	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
)

type ProfileServices struct {
	repo *repositories.ProfileRepository
}

func NewProfileServices(repo *repositories.ProfileRepository) *ProfileServices {
	return &ProfileServices{
		repo: repo,
	}
}

func (s *ProfileServices) GetProfileByID(ctx context.Context, id int) (*models.Profile, error) {
	return s.repo.GetProfileByID(ctx, id)
}

func (s *ProfileServices) CreateProfile(ctx context.Context, p *models.Profile) (*models.Profile, error) {
	return s.repo.CreateProfile(ctx, p)
}

func (s *ProfileServices) UpdateProfile(ctx context.Context, p *models.Profile) (*models.Profile, error) {
	return s.repo.UpdateProfile(ctx, p)
}

func (s *ProfileServices) DeleteProfile(ctx context.Context, id int) error {
	return s.repo.DeleteProfile(ctx, id)
}
