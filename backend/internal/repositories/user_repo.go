package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(ctx context.Context, u *models.UserAuthInput) (*models.ResponseUser, error) {
	res, err := r.db.NamedExecContext(ctx, "INSERT INTO users (email, password) VALUES (:email, :password)", u)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}
	ru := models.ResponseUser{
		ID:    int(id),
		Email: u.Email,
	}
	return &ru, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}
	err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user := models.User{}
	err := r.db.GetContext(ctx, &user, "SELECT id, email, password, created_at, updated_at FROM users WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}
	return &user, nil
}
