package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

type RefreshTokenRepository struct {
	db *sqlx.DB
}

func NewRefreshTokenRepository(db *sqlx.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, userID int, tokenHash string, expiresAt string) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)",
		userID, tokenHash, expiresAt)
	if err != nil {
		return fmt.Errorf("error creating refresh token: %w", err)
	}
	return nil
}

func (r *RefreshTokenRepository) GetByTokenHash(ctx context.Context, tokenHash string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.GetContext(ctx, &rt, "SELECT * FROM refresh_tokens WHERE token_hash = ?", tokenHash)
	if err != nil {
		return nil, fmt.Errorf("refresh token not found: %w", err)
	}
	return &rt, nil
}

func (r *RefreshTokenRepository) DeleteByID(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM refresh_tokens WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting refresh token: %w", err)
	}
	return nil
}
