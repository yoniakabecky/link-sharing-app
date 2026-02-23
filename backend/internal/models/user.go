package models

import "time"

type User struct {
	ID        int        `json:"id" db:"id"`
	Email     string     `json:"email" db:"email" validate:"required,email"`
	Password  string     `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterUser struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}

type ResponseUser struct {
	ID    int    `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}
