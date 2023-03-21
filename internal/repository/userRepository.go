package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/timofeyka25/test-jungle/internal/model"
)

type UserRepositoryInterface interface {
	GetByUsername(username string) (*model.User, error)
}

type UserRepositoryMySQL struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryMySQL {
	return UserRepositoryMySQL{db: db}
}

func (r UserRepositoryMySQL) GetByUsername(username string) (*model.User, error) {
	var user model.User
	query := "select * from users where username = ?"
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
