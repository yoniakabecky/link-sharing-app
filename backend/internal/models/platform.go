package models

type Platform struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Icon  string `db:"icon"`
	Color string `db:"color"`
}
