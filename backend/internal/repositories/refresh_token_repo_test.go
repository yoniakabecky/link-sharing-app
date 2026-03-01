package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestNewRefreshTokenRepository(t *testing.T) {
	withTestDB(t, func(db *sqlx.DB, _ sqlmock.Sqlmock) {
		repo := NewRefreshTokenRepository(db)
		require.NotNil(t, repo)
		require.NotNil(t, repo.db)
	})
}

func TestRefreshTokenRepository_Create(t *testing.T) {
	ctx := context.Background()
	userID := 1
	tokenHash := "abc123hash"
	expiresAt := "2026-12-31T23:59:59Z"

	tsc := []struct {
		name string
		test func(*testing.T, *RefreshTokenRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)").
					WithArgs(userID, tokenHash, expiresAt).
					WillReturnResult(sqlmock.NewResult(1, 1))

				err := repo.Create(ctx, userID, tokenHash, expiresAt)
				require.NoError(t, err)
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "failed to create refresh token",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)").
					WithArgs(userID, tokenHash, expiresAt).
					WillReturnError(errors.New("db error"))

				err := repo.Create(ctx, userID, tokenHash, expiresAt)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error creating refresh token")
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewRefreshTokenRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestRefreshTokenRepository_GetByTokenHash(t *testing.T) {
	ctx := context.Background()
	tokenHash := "abc123hash"
	expiresAt := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)
	createdAt := time.Now()

	tsc := []struct {
		name string
		test func(*testing.T, *RefreshTokenRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM refresh_tokens WHERE token_hash = ?").
					WithArgs(tokenHash).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "token_hash", "expires_at", "created_at"}).
						AddRow(1, 42, tokenHash, expiresAt, createdAt))

				rt, err := repo.GetByTokenHash(ctx, tokenHash)
				require.NoError(t, err)
				require.NotNil(t, rt)
				require.Equal(t, 1, rt.ID)
				require.Equal(t, 42, rt.UserID)
				require.Equal(t, tokenHash, rt.TokenHash)
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "refresh token not found",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM refresh_tokens WHERE token_hash = ?").
					WithArgs(tokenHash).
					WillReturnError(errors.New("sql: no rows in result set"))

				rt, err := repo.GetByTokenHash(ctx, tokenHash)
				require.Error(t, err)
				require.Nil(t, rt)
				require.Contains(t, err.Error(), "refresh token not found")
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewRefreshTokenRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestRefreshTokenRepository_DeleteByID(t *testing.T) {
	ctx := context.Background()
	tokenID := 1

	tsc := []struct {
		name string
		test func(*testing.T, *RefreshTokenRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM refresh_tokens WHERE id = ?").
					WithArgs(tokenID).
					WillReturnResult(sqlmock.NewResult(0, 1))

				err := repo.DeleteByID(ctx, tokenID)
				require.NoError(t, err)
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "failed to delete refresh token",
			test: func(t *testing.T, repo *RefreshTokenRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM refresh_tokens WHERE id = ?").
					WithArgs(tokenID).
					WillReturnError(errors.New("db error"))

				err := repo.DeleteByID(ctx, tokenID)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error deleting refresh token")
				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewRefreshTokenRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}
