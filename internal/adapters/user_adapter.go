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

type UserAdapter struct {
	db *sql.DB
}

func NewUserAdapter(db *sql.DB) *UserAdapter {
	return &UserAdapter{db}
}

type BusinessRuleError struct {
	StatusCode int
	Err error
}

func (e *BusinessRuleError) Error() string {
	return e.Err.Error()
}

var ErrEmptyEmail = errors.New("Email is a required attribute.")
var ErrEmptyPassword = errors.New("Password is a required attribute.")
var ErrEmptyName = errors.New("Name is a required attribute.")
var ErrEmptyIncome = errors.New("Income is a required attribute.")
var ErrEmailAlreadyExists = errors.New("Email already exists.")
var ErrUserNotFound = errors.New("User not found.")

func (adapter *UserAdapter) GetAllUsers() ([]*t.User, error) {
	sql, args, err := sq.Select("*").From("users").ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := adapter.db.Query(sql, args...)
 
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*t.User{}

	for rows.Next() {
		user := &t.User{}
		err := user.ScanRows(rows)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (adapter *UserAdapter) GetUserById(id int64) (*t.User, error) {
	query, args, err := sq.Select("*").From("users").Where(sq.Eq{"id": id}).ToSql()

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

func (adapter *UserAdapter) GetUserByEmail(email string) (*t.User, error) {
	query, args, err := sq.Select("*").From("users").Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, err
	}

	row := adapter.db.QueryRow(query, args...)

	user := &t.User{}

	err = user.ScanRow(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (adapter *UserAdapter) CreateUser(u *t.User) (*t.User, error) {
	switch {
		case u.Email == "":
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyEmail}
		case u.Password == "":
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyPassword}
		case u.Name == "":
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyName}
		case u.Income == 0:
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyIncome}
	}
	
	foundUser, err := adapter.GetUserByEmail(u.Email)

	if err != nil && err != ErrUserNotFound {
		return nil, err
	}

	if foundUser != nil {
		return nil, &BusinessRuleError{StatusCode: http.StatusConflict, Err: ErrEmailAlreadyExists}
	}

	user := &t.User{
		Name: u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Income: u.Income,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	query, args, err := sq.Insert("users").
	Columns("name", "email", "password", "income", "created_at", "updated_at").
	Values(user.Name, user.Email, user.Password, user.Income, user.CreatedAt, user.UpdatedAt).
	ToSql()

	if err != nil {
		return nil, err
	}

	repo := repositories.NewRepository(adapter.db)

	err = repo.CreateOrMerge(adapter.db, "users", query, user, args...)

	if err != nil {
		return nil, err
	}

	return user, nil
}