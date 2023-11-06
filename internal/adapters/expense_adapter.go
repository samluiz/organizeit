package adapters

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/samluiz/organizeit/internal/repositories"
	t "github.com/samluiz/organizeit/internal/types"
)

type ExpenseAdapter struct {
	db *sql.DB
}

func NewExpenseAdapter(db *sql.DB) *ExpenseAdapter {
	return &ExpenseAdapter{db}
}

var ErrEmptyAmount = errors.New("Amount is a required attribute.")
var ErrEmptyDescription = errors.New("Description is a required attribute.")

func (adapter *ExpenseAdapter) getAllExpenses() ([]*t.Expense, error) {
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

	return expenses, nil
}

func (adapter *ExpenseAdapter) getExpenseById(id int64) (*t.User, error) {
	query, args, err := sq.Select("*").From("expenses").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	row := adapter.db.QueryRow(query, args...)

	user := &t.User{}

	err = user.ScanRow(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &BusinessRuleError{StatusCode: http.StatusNotFound, Err: ErrUserNotFound}
		}
		return nil, err
	}

	return user, nil
}

func (adapter *ExpenseAdapter) CreateExpense(expense *t.Expense, userId int64, userAdapter *UserAdapter) (*t.Expense, error) {
	expense.CreatedAt = time.Now()
	expense.UpdatedAt = time.Now()

	user, err := userAdapter.GetUserById(userId)

	if err != nil {
		return nil, err
	}
	
	expense.UserId = int(user.ID)

	sql, args, err := sq.Insert("expenses").
		Columns("amount", "description", "user_id", "created_at", "created_at").
		Values(expense.Amount, expense.Description, expense.UserId, expense.CreatedAt, expense.UpdatedAt).
		ToSql()

	if err != nil {
		return nil, err
	}


	err = repositories.NewRepository(adapter.db).CreateOrMerge(adapter.db, "expenses", sql, expense, args...)

	if err != nil {
		return nil, err
	}

	return expense, nil
}