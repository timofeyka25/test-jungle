package repository

import "github.com/jmoiron/sqlx"

type ImageRepositoryInterface interface {
}

type ImageRepositoryMySQL struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) ImageRepositoryMySQL {
	return ImageRepositoryMySQL{db: db}
}
