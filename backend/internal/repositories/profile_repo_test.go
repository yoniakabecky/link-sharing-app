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

func TestCreateProfile(t *testing.T) {
	pls := []models.Link{
		{
			ProfileID:  1,
			PlatformID: 1,
			URL:        "https://example.com",
		},
		{
			ProfileID:  1,
			PlatformID: 2,
			URL:        "https://example.org",
		},
	}

	p := &models.Profile{
		UserID:    1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		AvatarURL: "https://example.com/avatar.png",
		Links:     pls,
	}

	tsc := []struct {
		name string
		test func(*testing.T, *ProfileRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(2, 2))
				mock.ExpectCommit()

				now := time.Now()
				prow := sqlmock.NewRows([]string{"id", "user_id", "first_name", "last_name", "email", "avatar_url", "created_at", "updated_at"}).
					AddRow(1, 1, "John", "Doe", "john.doe@example.com", "https://example.com/avatar.png", now, nil)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)
				plrows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com", now, nil).
					AddRow(2, 1, 2, "https://example.org", now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(plrows)

				result, err := repo.CreateProfile(context.Background(), p)
				require.NoError(t, err)
				require.Equal(t, 1, result.ID)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to create profile",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnError(errors.New("failed to create profile"))
				mock.ExpectRollback()

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get last insert ID for profile",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewErrorResult(errors.New("failed to get last insert ID for profile")))
				mock.ExpectRollback()

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to begin transaction",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to commit transaction",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(2, 2))
				mock.ExpectCommit().WillReturnError(errors.New("failed to commit transaction"))

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to rollback transaction",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnError(errors.New("failed to create profile"))
				mock.ExpectRollback().WillReturnError(errors.New("failed to rollback transaction"))

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to create links",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnError(errors.New("failed to create links"))
				mock.ExpectRollback()

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get last insert ID for links",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewErrorResult(errors.New("failed to get last insert ID for links")))
				mock.ExpectRollback()

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed GetProfileByID after create",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(2, 2))
				mock.ExpectCommit()

				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnError(errors.New("failed to get profile"))

				_, err := repo.CreateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewProfileRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestGetProfileByID(t *testing.T) {
	pls := []models.Link{
		{
			ID:         1,
			ProfileID:  1,
			PlatformID: 1,
			URL:        "https://example.com",
		},
		{
			ID:         2,
			ProfileID:  1,
			PlatformID: 2,
			URL:        "https://example.org",
		},
	}

	p := &models.Profile{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		AvatarURL: "https://example.com/avatar.png",
		Links:     pls,
	}

	tsc := []struct {
		name string
		test func(*testing.T, *ProfileRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				prow := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "avatar_url"}).AddRow(p.ID, p.FirstName, p.LastName, p.Email, p.AvatarURL)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)

				plrows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url"}).AddRow(1, 1, 1, "https://example.com").AddRow(2, 1, 2, "https://example.org")
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(plrows)

				p, err := repo.GetProfileByID(context.Background(), 1)
				require.NoError(t, err)
				require.Equal(t, int(1), p.ID)

				for i, l := range p.Links {
					require.Equal(t, pls[i].ID, l.ID)
					require.Equal(t, pls[i].ProfileID, l.ProfileID)
					require.Equal(t, pls[i].PlatformID, l.PlatformID)
					require.Equal(t, pls[i].URL, l.URL)
				}

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get profile",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnError(errors.New("failed to get profile"))

				_, err := repo.GetProfileByID(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get links",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				prow := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "avatar_url"}).AddRow(p.ID, p.FirstName, p.LastName, p.Email, p.AvatarURL)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)

				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnError(errors.New("failed to get links"))

				_, err := repo.GetProfileByID(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewProfileRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	pls := []models.Link{
		{
			ID:         1,
			ProfileID:  1,
			PlatformID: 1,
			URL:        "https://example.com",
		},
	}
	p := &models.Profile{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		AvatarURL: "https://example.com/avatar.png",
		Links:     pls,
	}

	tsc := []struct {
		name string
		test func(*testing.T, *ProfileRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE profiles SET updated_at = CURRENT_TIMESTAMP WHERE id = ?").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				prow := sqlmock.NewRows([]string{"id", "user_id", "first_name", "last_name", "email", "avatar_url", "created_at", "updated_at"}).
					AddRow(1, 0, "John", "Doe", "john.doe@example.com", "https://example.com/avatar.png", now, nil)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)
				plrows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com", now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(plrows)

				result, err := repo.UpdateProfile(context.Background(), p)
				require.NoError(t, err)
				require.Equal(t, 1, result.ID)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to update profile",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnError(errors.New("failed to update profile"))
				mock.ExpectRollback()

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to update links",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnError(errors.New("failed to update links"))
				mock.ExpectRollback()

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to begin transaction for update",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to commit transaction for update",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE profiles SET updated_at = CURRENT_TIMESTAMP WHERE id = ?").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit().WillReturnError(errors.New("failed to commit transaction"))

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to rollback transaction for update",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnError(errors.New("failed to update profile"))
				mock.ExpectRollback().WillReturnError(errors.New("failed to rollback transaction"))

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "success with link deletion",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				// Profile has links 1 and 2 in DB, we update with only link 1 (removing link 2)
				pWithOneLink := &models.Profile{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					AvatarURL: "https://example.com/avatar.png",
					Links: []models.Link{
						{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com/updated"},
					},
				}
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil).
						AddRow(2, 1, 2, "https://example.org", now, nil))
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE profiles SET updated_at = CURRENT_TIMESTAMP WHERE id = ?").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				prow := sqlmock.NewRows([]string{"id", "user_id", "first_name", "last_name", "email", "avatar_url", "created_at", "updated_at"}).
					AddRow(1, 0, "John", "Doe", "john.doe@example.com", "https://example.com/avatar.png", now, nil)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)
				plrows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com/updated", now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(plrows)

				result, err := repo.UpdateProfile(context.Background(), pWithOneLink)
				require.NoError(t, err)
				require.Equal(t, 1, result.ID)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "success with new link",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				pWithNewLink := &models.Profile{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					AvatarURL: "https://example.com/avatar.png",
					Links: []models.Link{
						{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"},
						{ID: 0, ProfileID: 1, PlatformID: 2, URL: "https://example.org/new"},
					},
				}
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links (profile_id, platform_id, url) VALUES (?, ?, ?)").WillReturnResult(sqlmock.NewResult(2, 1))
				mock.ExpectExec("UPDATE profiles SET updated_at = CURRENT_TIMESTAMP WHERE id = ?").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()

				prow := sqlmock.NewRows([]string{"id", "user_id", "first_name", "last_name", "email", "avatar_url", "created_at", "updated_at"}).
					AddRow(1, 0, "John", "Doe", "john.doe@example.com", "https://example.com/avatar.png", now, nil)
				mock.ExpectQuery("SELECT * FROM profiles WHERE id = ?").WillReturnRows(prow)
				plrows := sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "https://example.com", now, nil).
					AddRow(2, 1, 2, "https://example.org/new", now, nil)
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(plrows)

				result, err := repo.UpdateProfile(context.Background(), pWithNewLink)
				require.NoError(t, err)
				require.Equal(t, 1, result.ID)
				require.Len(t, result.Links, 2)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to get links in update",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnError(errors.New("failed to get links"))
				mock.ExpectRollback()

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to touch profile updated_at",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil))
				mock.ExpectExec("UPDATE links SET platform_id = ?, url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("UPDATE profiles SET updated_at = CURRENT_TIMESTAMP WHERE id = ?").WillReturnError(errors.New("failed to touch updated_at"))
				mock.ExpectRollback()

				_, err := repo.UpdateProfile(context.Background(), p)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to delete link",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				now := time.Now()
				pWithOneLink := &models.Profile{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john.doe@example.com",
					AvatarURL: "https://example.com/avatar.png",
					Links: []models.Link{
						{ID: 1, ProfileID: 1, PlatformID: 1, URL: "https://example.com"},
					},
				}
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE profiles SET first_name = ?, last_name = ?, email = ?, avatar_url = ? WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT * FROM links WHERE profile_id = ?").WillReturnRows(
					sqlmock.NewRows([]string{"id", "profile_id", "platform_id", "url", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "https://example.com", now, nil).
						AddRow(2, 1, 2, "https://example.org", now, nil))
				mock.ExpectExec("DELETE FROM links WHERE id = ?").WillReturnError(errors.New("failed to delete link"))
				mock.ExpectRollback()

				_, err := repo.UpdateProfile(context.Background(), pWithOneLink)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewProfileRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}

func TestDeleteProfile(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T, *ProfileRepository, sqlmock.Sqlmock)
	}{
		{
			name: "success",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM profiles WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM links WHERE profile_id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err := repo.DeleteProfile(context.Background(), 1)
				require.NoError(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to delete profile",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM profiles WHERE id = ?").WillReturnError(errors.New("failed to delete profile"))
				mock.ExpectRollback()

				err := repo.DeleteProfile(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to delete links",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM profiles WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM links WHERE profile_id = ?").WillReturnError(errors.New("failed to delete links"))
				mock.ExpectRollback()

				err := repo.DeleteProfile(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to begin transaction for delete",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))

				err := repo.DeleteProfile(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to commit transaction for delete",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM profiles WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM links WHERE profile_id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit().WillReturnError(errors.New("failed to commit transaction"))

				err := repo.DeleteProfile(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
		{
			name: "failed to rollback transaction for delete",
			test: func(t *testing.T, repo *ProfileRepository, mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM profiles WHERE id = ?").WillReturnError(errors.New("failed to delete profile"))
				mock.ExpectRollback().WillReturnError(errors.New("failed to rollback transaction"))

				err := repo.DeleteProfile(context.Background(), 1)
				require.Error(t, err)

				err = mock.ExpectationsWereMet()
				require.NoError(t, err)
			},
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			withTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
				repo := NewProfileRepository(db)
				tc.test(t, repo, mock)
			})
		})
	}
}
