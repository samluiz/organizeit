package types

import (
	"database/sql"
	"time"
)

type Expense struct {
	ID int 				`json:"id"`
	Amount float64 		`json:"amount"`
	Description string 	`json:"description"`
	UserId int 			`json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Expense) ScanRow(row *sql.Row) error {
	return row.Scan(&e.ID, &e.Amount, &e.Description, &e.UserId, &e.CreatedAt, &e.UpdatedAt)
}

func (e *Expense) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Amount, &e.Description, &e.UserId, &e.CreatedAt, &e.UpdatedAt)
}