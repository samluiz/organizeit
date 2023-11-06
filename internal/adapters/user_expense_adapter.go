package adapters

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	t "github.com/samluiz/organizeit/internal/types"
)

type UserExpenseAdapter struct {
	db *sql.DB
}

func NewUserExpenseAdapter(db *sql.DB) *UserExpenseAdapter {
	return &UserExpenseAdapter{db}
}

func (adapter *ExpenseAdapter) getUsersWithExpenses() ([]*t.UserResponse, error) {
	sql, args, err := sq.Select("*").From("expenses").ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := adapter.db.Query(sql, args...)
 
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	expenses := []*t.Expense{}

	for rows.Next() {
		expense := &t.Expense{}
		err := expense.ScanRows(rows)

		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return nil, nil
}