package models

import "time"

type Link struct {
	ID         int        `json:"id" db:"id"`
	ProfileID  int        `json:"profile_id" db:"profile_id"`
	PlatformID int        `json:"platform_id" db:"platform_id"`
	URL        string     `json:"url" db:"url"`
	Position   int        `json:"position" db:"position"`
	Platform   *Platform  `json:"platform" db:"-"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}
