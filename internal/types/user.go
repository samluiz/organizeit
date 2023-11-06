package types

import (
	"database/sql"
	"time"
)

type User struct {
	ID int64				`json:"id"`
	Name string 			`json:"name"`
	Email string 			`json:"email"`
	Password string 		`json:"password"`
	Income float64 			`json:"income"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}

func (u *User) ScanRow(row *sql.Row) error {
	return row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Income, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Income, &u.CreatedAt, &u.UpdatedAt)
}