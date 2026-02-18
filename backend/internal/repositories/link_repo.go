package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
)

type LinkRepository struct {
	db    *sqlx.DB
	pRepo *PlatformRepository
}

func NewLinkRepository(db *sqlx.DB, pRepo *PlatformRepository) *LinkRepository {
	return &LinkRepository{
		db:    db,
		pRepo: pRepo,
	}
}

func (r *LinkRepository) GetLinksByProfileID(ctx context.Context, prfID int) ([]models.Link, error) {
	return r.getLinksByProfileID(ctx, prfID)
}

func (r *LinkRepository) GetLinksByProfileIDTx(ctx context.Context, tx *sqlx.Tx, profileID int) ([]models.Link, error) {
	return r.getLinksByProfileIDTx(ctx, tx, profileID)
}

func (r *LinkRepository) getLinksByProfileID(ctx context.Context, prfID int) ([]models.Link, error) {
	var links []models.Link
	err := r.db.SelectContext(ctx, &links, "SELECT * FROM links WHERE profile_id = ?", prfID)
	if err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}

	for i := range links {
		links[i].Platform, err = r.pRepo.GetPlatformByID(ctx, links[i].PlatformID)
		if err != nil {
			return nil, fmt.Errorf("error getting links: %w", err)
		}
	}
	return links, nil
}

func (r *LinkRepository) getLinksByProfileIDTx(ctx context.Context, tx *sqlx.Tx, prfID int) ([]models.Link, error) {
	var links []models.Link
	err := tx.SelectContext(ctx, &links, "SELECT * FROM links WHERE profile_id = ?", prfID)
	if err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}

	for i := range links {
		links[i].Platform, err = r.pRepo.GetPlatformByID(ctx, links[i].PlatformID)
		if err != nil {
			return nil, fmt.Errorf("error getting links: %w", err)
		}
	}
	return links, nil
}

func (r *LinkRepository) CreateLinkTx(ctx context.Context, tx *sqlx.Tx, link *models.Link) error {
	return createLink(ctx, tx, link)
}

func (r *LinkRepository) UpdateLinkTx(ctx context.Context, tx *sqlx.Tx, l *models.Link) error {
	return updateLink(ctx, tx, l)
}

func (r *LinkRepository) DeleteLinkByIDTx(ctx context.Context, tx *sqlx.Tx, linkID int) error {
	return deleteLinkByID(ctx, tx, linkID)
}

func createLink(ctx context.Context, tx *sqlx.Tx, link *models.Link) error {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO links (profile_id, platform_id, url) VALUES (:profile_id, :platform_id, :url)", link)
	if err != nil {
		return fmt.Errorf("error inserting link: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}
	link.ID = int(id)
	return nil
}

func updateLink(ctx context.Context, tx *sqlx.Tx, l *models.Link) error {
	if l.ID == 0 {
		return createLink(ctx, tx, l)
	}
	_, err := tx.NamedExecContext(ctx, "UPDATE links SET platform_id = :platform_id, url = :url WHERE id = :id", l)
	if err != nil {
		return fmt.Errorf("error updating link: %w", err)
	}
	return nil
}

func deleteLinkByID(ctx context.Context, tx *sqlx.Tx, lID int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM links WHERE id = ?", lID)
	if err != nil {
		return fmt.Errorf("error deleting link: %w", err)
	}
	return nil
}

func (r *LinkRepository) UpdateLinks(ctx context.Context, profileID int, links []models.Link) error {
	return r.execTx(ctx, func(tx *sqlx.Tx) error {
		linkIDs := make(map[int]struct{}, len(links))
		for i := range links {
			links[i].ProfileID = profileID
			if links[i].ID != 0 {
				linkIDs[links[i].ID] = struct{}{}
			}
		}

		current, err := r.getLinksByProfileIDTx(ctx, tx, profileID)
		if err != nil {
			return fmt.Errorf("error getting links: %w", err)
		}

		for _, l := range current {
			if _, keep := linkIDs[l.ID]; !keep {
				if err := deleteLinkByID(ctx, tx, l.ID); err != nil {
					return err
				}
			}
		}
		for i := range links {
			if err := updateLink(ctx, tx, &links[i]); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *LinkRepository) execTx(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil
}
