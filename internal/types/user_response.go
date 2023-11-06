package types

import (
	"database/sql"
	"time"
)

type UserResponse struct {
	ID int64				`json:"id"`
	Name string 			`json:"name"`
	Email string 			`json:"email"`
	Income float64 			`json:"income"`
	Expenses []*Expense 	`json:"expenses"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}

func (u *UserResponse) ScanRow(row *sql.Row) error {
	return row.Scan(&u.ID, &u.Name, &u.Email, &u.Income, &u.Expenses, &u.CreatedAt, &u.UpdatedAt)
}

func (u *UserResponse) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&u.ID, &u.Name, &u.Email, &u.Income, &u.Expenses, &u.CreatedAt, &u.UpdatedAt)
}