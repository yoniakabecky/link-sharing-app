package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	input := &models.UserAuthInput{Email: "test@example.com", Password: "hashed"}

	tsc := []struct {
		name string
		test func(*testing.T, *UserRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").
					WithArgs(input.Email, input.Password).
					WillReturnResult(sqlmock.NewResult(1, 1))

				got, err := repo.Register(ctx, input)
				require.NoError(t, err)
				require.Equal(t, 1, got.ID)
				require.Equal(t, "test@example.com", got.Email)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "failed to register user",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").
					WithArgs(input.Email, input.Password).
					WillReturnError(errors.New("failed to register user"))

				_, err := repo.Register(ctx, input)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error registering user")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "failed to get last insert ID",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").
					WithArgs(input.Email, input.Password).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert id not supported")))

				_, err := repo.Register(ctx, input)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error getting last insert ID")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewUserRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	ctx := context.Background()
	email := "test@example.com"

	tsc := []struct {
		name string
		test func(*testing.T, *UserRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM users WHERE email = ?").
					WithArgs(email).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
						AddRow(1, "test@example.com", "hashed", time.Time{}, nil))

				user, err := repo.GetUserByEmail(ctx, email)
				require.NoError(t, err)
				require.Equal(t, 1, user.ID)
				require.Equal(t, "test@example.com", user.Email)
				require.Equal(t, "hashed", user.Password)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "failed to get user by email",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM users WHERE email = ?").
					WithArgs(email).
					WillReturnError(errors.New("failed to get user by email"))

				_, err := repo.GetUserByEmail(ctx, email)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error getting user by email")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewUserRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}
