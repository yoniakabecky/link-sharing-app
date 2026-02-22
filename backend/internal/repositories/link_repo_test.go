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

func TestNewLinkRepository(t *testing.T) {
	withTestDB(t, func(db *sqlx.DB, _ sqlmock.Sqlmock) {
		pRepo := NewPlatformRepository(db)
		repo := NewLinkRepository(db, pRepo)
		require.NotNil(t, repo)
		require.NotNil(t, repo.db)
	})
}

func TestGetLinksByProfileID(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com", 0, now, nil).
					AddRow(2, 1, 2, "https://example.org", 1, now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).WillReturnRows(rows)
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(2).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(2, "Facebook", "facebook", "#000000"))

				links, err := repo.GetLinksByProfileID(context.Background(), 1)
				require.NoError(t, err)
				require.Len(t, links, 2)
				require.Equal(t, 1, links[0].ID)
				require.Equal(t, "https://example.com", links[0].URL)
				require.Equal(t, 2, links[1].ID)
				require.Equal(t, "https://example.org", links[1].URL)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success empty",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"})
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).WillReturnRows(rows)

				links, err := repo.GetLinksByProfileID(context.Background(), 1)
				require.NoError(t, err)
				require.Empty(t, links)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error getting links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnError(errors.New("db error"))

				links, err := repo.GetLinksByProfileID(context.Background(), 1)
				require.Error(t, err)
				require.Nil(t, links)
				require.Contains(t, err.Error(), "error getting links")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestGetLinksByProfileIDTx(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				now := time.Now()
				rows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com", 0, now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).WillReturnRows(rows)
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectCommit()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				links, err := repo.GetLinksByProfileIDTx(context.Background(), tx, 1)
				require.NoError(t, err)
				require.Len(t, links, 1)
				require.Equal(t, 1, links[0].ID)
				require.NoError(t, tx.Commit())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error getting links in tx",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnError(errors.New("tx select error"))
				mock.ExpectRollback()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				links, err := repo.GetLinksByProfileIDTx(context.Background(), tx, 1)
				require.Error(t, err)
				require.Nil(t, links)
				require.Contains(t, err.Error(), "error getting links")
				require.NoError(t, tx.Rollback())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestCreateLinkTx(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(1, 1, "https://example.com", 0).
					WillReturnResult(sqlmock.NewResult(5, 1))
				mock.ExpectCommit()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)

				link := &models.Link{ProfileID: 1, PlatformID: 1, URL: "https://example.com"}
				err = repo.CreateLinkTx(context.Background(), tx, link)
				require.NoError(t, err)
				require.Equal(t, 5, link.ID)
				require.NoError(t, tx.Commit())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error inserting link",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(1, 1, "https://example.com", 0).
					WillReturnError(errors.New("insert failed"))
				mock.ExpectRollback()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				link := &models.Link{ProfileID: 1, PlatformID: 1, URL: "https://example.com"}
				err = repo.CreateLinkTx(context.Background(), tx, link)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error inserting link")
				require.NoError(t, tx.Rollback())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error getting last insert ID",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(1, 1, "https://example.com", 0).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert id failed")))
				mock.ExpectRollback()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				link := &models.Link{ProfileID: 1, PlatformID: 1, URL: "https://example.com"}
				err = repo.CreateLinkTx(context.Background(), tx, link)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error getting last insert ID")
				require.NoError(t, tx.Rollback())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestUpdateLinkTx(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success update existing",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(2, "https://example.com/updated", 0, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)

				link := &models.Link{ID: 1, ProfileID: 1, PlatformID: 2, URL: "https://example.com/updated"}
				err = repo.UpdateLinkTx(context.Background(), tx, link)
				require.NoError(t, err)
				require.NoError(t, tx.Commit())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success create when ID is zero",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(1, 2, "https://example.com/new", 0).
					WillReturnResult(sqlmock.NewResult(10, 1))
				mock.ExpectCommit()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)

				link := &models.Link{ID: 0, ProfileID: 1, PlatformID: 2, URL: "https://example.com/new"}
				err = repo.UpdateLinkTx(context.Background(), tx, link)
				require.NoError(t, err)
				require.Equal(t, 10, link.ID)
				require.NoError(t, tx.Commit())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error updating link",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(2, "https://example.com/updated", 0, 1).
					WillReturnError(errors.New("update failed"))
				mock.ExpectRollback()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				link := &models.Link{ID: 1, ProfileID: 1, PlatformID: 2, URL: "https://example.com/updated"}
				err = repo.UpdateLinkTx(context.Background(), tx, link)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error updating link")
				require.NoError(t, tx.Rollback())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestDeleteLinkByIDTx(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)

				err = repo.DeleteLinkByIDTx(context.Background(), tx, 1)
				require.NoError(t, err)
				require.NoError(t, tx.Commit())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error deleting link",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WithArgs(1).WillReturnError(errors.New("delete failed"))
				mock.ExpectRollback()

				tx, err := repo.db.BeginTxx(context.Background(), nil)
				require.NoError(t, err)
				defer tx.Rollback()

				err = repo.DeleteLinkByIDTx(context.Background(), tx, 1)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error deleting link")
				require.NoError(t, tx.Rollback())

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestUpdateLinks(t *testing.T) {
	now := time.Now()
	tsc := []struct {
		name string
		test func(*testing.T, *LinkRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success add new links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{
					{ID: 0, ProfileID: 1, PlatformID: 1, URL: "https://example.com/new"},
				}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(1, 1, "https://example.com/new", 0).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.NoError(t, err)
				require.Equal(t, 1, links[0].ID)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success update existing links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{
					{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com/updated"},
				}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", 0, now, nil))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(1, "https://example.com/updated", 0, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.NoError(t, err)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success delete removed links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{
					{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"},
				}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", 0, now, nil).
						AddRow(2, 1, 2, "https://example.org", 1, now, nil))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(2).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(2, "Facebook", "facebook", "#000000"))
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WithArgs(2).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(1, "https://example.com", 0, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.NoError(t, err)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success empty links no current",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}))
				mock.ExpectCommit()

				err := repo.UpdateLinks(context.Background(), 1, []models.Link{})
				require.NoError(t, err)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error begin transaction",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

				err := repo.UpdateLinks(context.Background(), 1, []models.Link{{ID: 1, URL: "https://example.com"}})
				require.Error(t, err)
				require.Contains(t, err.Error(), "error starting transaction")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error getting current links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnError(errors.New("select failed"))
				mock.ExpectRollback()

				err := repo.UpdateLinks(context.Background(), 1, []models.Link{{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"}})
				require.Error(t, err)
				require.Contains(t, err.Error(), "error getting links")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error deleting link",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"}}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", 0, now, nil).
						AddRow(2, 1, 2, "https://example.org", 1, now, nil))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(2).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(2, "Facebook", "facebook", "#000000"))
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WithArgs(2).WillReturnError(errors.New("delete failed"))
				mock.ExpectRollback()

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error deleting link")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error updating link",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"}}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", 0, now, nil))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(1, "https://example.com", 0, 1).
					WillReturnError(errors.New("update failed"))
				mock.ExpectRollback()

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error updating link")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "error commit transaction",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"}}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", 0, now, nil))
				mock.ExpectQuery("SELECT * FROM platforms WHERE id = ?").WithArgs(1).WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "icon", "color"}).AddRow(1, "Twitter", "twitter", "#000000"))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ?, position = ? WHERE id = ?").
					WithArgs(1, "https://example.com", 0, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit().WillReturnError(errors.New("commit failed"))

				err := repo.UpdateLinks(context.Background(), 1, links)
				require.Error(t, err)
				require.Contains(t, err.Error(), "error committing transaction")

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
		{
			name: "success assigns profile ID to links",
			test: func(t *testing.T, repo *LinkRepository, mock sqlmock.Sqlmock) {
				links := []models.Link{
					{ID: 0, ProfileID: 0, PlatformID: 1, URL: "https://example.com"},
				}
				mock.ExpectBegin()
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WithArgs(99).
					WillReturnRows(sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "position", "created_at", "updated_at"}))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url, position) VALUES (?, ?, ?, ?)").
					WithArgs(99, 1, "https://example.com", 0).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err := repo.UpdateLinks(context.Background(), 99, links)
				require.NoError(t, err)
				require.Equal(t, 99, links[0].ProfileID)

				require.NoError(t, mock.ExpectationsWereMet())
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				pRepo := NewPlatformRepository(db)
				repo := NewLinkRepository(db, pRepo)
				tc.test(t, repo, mock)
			})
		})
	}
}
