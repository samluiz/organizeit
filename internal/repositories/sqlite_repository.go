package repositories

import (
	"database/sql"
)

type Scanner interface {
	ScanRow(*sql.Row) error
}

func create(db *sql.DB, tableName string, sql string, s Scanner, args ...interface{}) error {
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