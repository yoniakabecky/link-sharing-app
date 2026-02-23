package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestRegister(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *UserRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed to register user",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").WillReturnError(errors.New("failed to register user"))
				mock.ExpectRollback()
			},
		},
		{
			name: "failed to get last insert ID",
			test: func(t *testing.T, repo *UserRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO users (email, password) VALUES (?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
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
