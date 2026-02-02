package models

import "time"

type Profile struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	AvatarURL string     `json:"avatar_url" db:"avatar_url"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	Links     []Link     `json:"links"`
}
