package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

type ProfileRepository struct {
	db    *sqlx.DB
	lRepo *LinkRepository
}

func NewProfileRepository(db *sqlx.DB, linkRepo *LinkRepository) *ProfileRepository {
	return &ProfileRepository{
		db:    db,
		lRepo: linkRepo,
	}
}

func (r *ProfileRepository) CreateProfile(ctx context.Context, p *models.Profile) (*models.Profile, error) {
	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		p, err := createProfile(ctx, tx, p)
		if err != nil {
			return fmt.Errorf("error creating profile: %w", err)
		}

		for i := range p.Links {
			p.Links[i].ProfileID = p.ID
			if err := r.lRepo.CreateLinkTx(ctx, tx, &p.Links[i]); err != nil {
				return fmt.Errorf("error creating link: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error creating profile: %w", err)
	}

	np, err := r.GetProfileByID(ctx, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating profile: %w", err)
	}
	return np, nil
}

func createProfile(ctx context.Context, tx *sqlx.Tx, p *models.Profile) (*models.Profile, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO profiles (user_id, first_name, last_name, email, avatar_url) VALUES (:user_id, :first_name, :last_name, :email, :avatar_url)", p)
	if err != nil {
		return nil, fmt.Errorf("error inserting profile: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	p.ID = int(id)

	return p, nil
}

func (r *ProfileRepository) GetProfileByID(ctx context.Context, id int) (*models.Profile, error) {
	var p models.Profile
	err := r.db.GetContext(ctx, &p, "SELECT * FROM profiles WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting profile: %w", err)
	}

	links, err := r.lRepo.GetLinksByProfileID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting profile: %w", err)
	}
	p.Links = links
	return &p, nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, p *models.Profile) (*models.Profile, error) {
	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, "UPDATE profiles SET first_name = :first_name, last_name = :last_name, email = :email, avatar_url = :avatar_url WHERE id = :id", p)
		if err != nil {
			return fmt.Errorf("error updating profile: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error updating profile: %w", err)
	}

	up, err := r.GetProfileByID(ctx, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating profile: %w", err)
	}
	return up, nil
}

func (r *ProfileRepository) DeleteProfile(ctx context.Context, id int) error {
	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM profiles WHERE id = ?", id)
		if err != nil {
			return fmt.Errorf("error deleting profile: %w", err)
		}

		_, err = tx.ExecContext(ctx, "DELETE FROM links WHERE profile_id = ?", id)
		if err != nil {
			return fmt.Errorf("error deleting links: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting profile: %w", err)
	}
	return nil
}

func (r *ProfileRepository) execTx(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", rbErr)
		}
		return fmt.Errorf("error executing transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
