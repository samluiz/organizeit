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

const TABLE_NAME = "users"

type UserAdapter struct {
	db *sql.DB
}

func NewAdapter(db *sql.DB) *UserAdapter {
	return &UserAdapter{db}
}

type BusinessRuleError struct {
	StatusCode int
	Err error
}

func (e *BusinessRuleError) Error() string {
	return e.Err.Error()
}

var ErrEmptyEmail = errors.New("Email is a required attribute")
var ErrEmptyPassword = errors.New("Password is a required attribute")
var ErrEmailAlreadyExists = errors.New("Email already exists")

func (adapter *UserAdapter) GetAllUsers() ([]*t.User, error) {
	sql, args, err := sq.Select("*").From(TABLE_NAME).ToSql()

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
	sql, args, err := sq.Select("*").From(TABLE_NAME).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	row := adapter.db.QueryRow(sql, args...)

	user := &t.User{}

	err = user.ScanRow(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (adapter *UserAdapter) GetUserByEmail(email string) (*t.User, error) {
	sql, args, err := sq.Select("*").From(TABLE_NAME).Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, err
	}

	row := adapter.db.QueryRow(sql, args...)

	user := &t.User{}

	err = user.ScanRow(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (adapter *UserAdapter) CreateUser(email string, password string) (*t.User, error) {
	switch {
		case email == "":
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyEmail}
		case password == "":
			return nil, &BusinessRuleError{StatusCode: http.StatusBadRequest, Err: ErrEmptyPassword}
	}
	
	foundUser, err := adapter.GetUserByEmail(email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if foundUser != nil {
		return nil, &BusinessRuleError{StatusCode: http.StatusConflict, Err: ErrEmailAlreadyExists}
	}

	user := &t.User{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sql, args, err := sq.Insert(TABLE_NAME).
	Columns("email", "password", "created_at", "updated_at").
	Values(user.Email, user.Password, user.CreatedAt, user.UpdatedAt).
	ToSql()

	if err != nil {
		return nil, err
	}

	repo := repositories.NewRepository(adapter.db)

	err = repo.Create(adapter.db, TABLE_NAME, sql, user, args...)

	if err != nil {
		return nil, err
	}

	return user, nil
}