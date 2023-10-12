package types

import (
	"database/sql"
	"time"
)

type User struct {
	ID int64			`json:"id"`
	Email string 		`json:"email"`
	Password string 	`json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ScanRow(row *sql.Row) error {
	return row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
}