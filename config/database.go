package config

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "config/database/organizeit.db")
	if err != nil {
		return nil, err
	}

	err = createTables(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	sqlScript, err := os.ReadFile("config/database/scripts/create_tables.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlScript))
	if err != nil {
		return err
	}

	return nil
}