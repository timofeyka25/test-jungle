package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/timofeyka25/test-jungle/internal/model"
)

type ImageRepositoryInterface interface {
	AddImage(image *model.Image) error
	GetImages(userId int64) ([]model.Image, error)
}

type ImageRepositoryMySQL struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) ImageRepositoryMySQL {
	return ImageRepositoryMySQL{db: db}
}

func (r ImageRepositoryMySQL) AddImage(image *model.Image) error {
	query := "insert into images (user_id, image_path, image_url) values (?, ?, ?)"
	res, err := r.db.Exec(query, image.UserId, image.ImagePath, image.ImageURL)
	if err != nil {
		return err
	}
	image.Id, err = res.LastInsertId()
	return err
}

func (r ImageRepositoryMySQL) GetImages(userId int64) ([]model.Image, error) {
	var images []model.Image
	query := "select * from images where user_id = ?"
	err := r.db.Select(&images, query, userId)
	if err != nil {
		return nil, err
	}
	return images, nil
}
