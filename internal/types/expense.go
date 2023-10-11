package types

import (
	"time"
)

type Expense struct {
	ID int 				`json:"id"`
	Amount float64 		`json:"amount"`
	Description string 	`json:"description"`
	Person Person 		`json:"person"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}