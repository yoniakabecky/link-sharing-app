package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {
	dns := "root:root_password@tcp(localhost:3306)/link-sharing-app?parseTime=true"
	db, err := sqlx.Open("mysql", dns)
	if err != nil {
		return nil, fmt.Errorf("failed opening database: %w", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}
