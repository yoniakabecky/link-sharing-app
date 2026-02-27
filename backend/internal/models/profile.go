package models

import "time"

type Profile struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	Nickname  string     `json:"nickname" db:"nickname" validate:"required,min=2,max=255"`
	FirstName string     `json:"first_name" db:"first_name" validate:"required"`
	LastName  string     `json:"last_name" db:"last_name" validate:"required"`
	Email     string     `json:"email" db:"email" validate:"omitempty,email"`
	AvatarURL string     `json:"avatar_url" db:"avatar_url" validate:"omitempty,url"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	Links     []Link     `json:"links" validate:"omitempty,dive"`
}
