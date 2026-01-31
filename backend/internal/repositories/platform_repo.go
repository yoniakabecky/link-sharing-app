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
