package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func withTestDB(t *testing.T, fn func(*sqlx.DB, sqlmock.Sqlmock)) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	fn(db, mock)
}

func TestGetAll(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *PlatformRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *PlatformRepository, mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000").AddRow(2, "Facebook", "facebook", "#000000")
				mock.ExpectQuery("SELECT * FROM platforms").WillReturnRows(rows)

				ps, err := repo.GetAll(context.Background())
				require.NoError(t, err)
				require.Len(t, ps, 2)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get all platforms",
			test: func(t *testing.T, repo *PlatformRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM platforms").WillReturnError(errors.New("failed to get all platforms"))

				_, err := repo.GetAll(context.Background())
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewPlatformRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}
