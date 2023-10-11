package repositories

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	t "github.com/samluiz/organizeit/internal/types"
)

const TABLE_NAME = "users"

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(email string, password string) (*t.User, error) {
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

	err = create(r.db, TABLE_NAME, sql, user, args...)

	if err != nil {
		return nil, err
	}

	return user, nil
}