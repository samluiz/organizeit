package repositories

import (
	"database/sql"
)

type Scanner interface {
	ScanRow(*sql.Row) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateOrMerge(db *sql.DB, tableName string, sql string, s Scanner, args ...interface{}) error {
	rs, err := db.Exec(sql, args...)

	if err != nil {
		return err
	}

	lastInsertId, err := rs.LastInsertId()

	if err != nil {
		return err
	}

	return s.ScanRow(db.QueryRow("SELECT * FROM " + tableName + " WHERE id = ?", lastInsertId))
}

func (r *Repository) QueryRow(db *sql.DB, sql string, s Scanner, args ...interface{}) error {
	return s.ScanRow(db.QueryRow(sql, args...))
}