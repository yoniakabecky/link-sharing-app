package models

type Platform struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Icon  string `json:"icon" db:"icon"`
	Color string `json:"color" db:"color"`
}
