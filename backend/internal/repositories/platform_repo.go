package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

type PlatformRepository struct {
	db *sqlx.DB
}

func NewPlatformRepository(db *sqlx.DB) *PlatformRepository {
	return &PlatformRepository{
		db: db,
	}
}

func (r *PlatformRepository) GetAll(ctx context.Context) ([]models.Platform, error) {
	platforms := []models.Platform{}
	err := r.db.Select(&platforms, "SELECT * FROM platforms")
	if err != nil {
		return nil, fmt.Errorf("failed to get all platforms: %w", err)
	}
	return platforms, nil
}

func (r *PlatformRepository) GetPlatformByID(ctx context.Context, id int) (*models.Platform, error) {
	platform := models.Platform{}
	err := r.db.Get(&platform, "SELECT * FROM platforms WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get platform by ID: %w", err)
	}
	return &platform, nil
}
