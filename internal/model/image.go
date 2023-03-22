package model

type Image struct {
	Id        int64
	UserId    int64  `db:"user_id"`
	ImagePath string `db:"image_path"`
	ImageURL  string `db:"image_url"`
}

func NewImage(
	userId int64,
	imagePath string,
	imageUrl string,
) *Image {
	return &Image{
		UserId:    userId,
		ImagePath: imagePath,
		ImageURL:  imageUrl,
	}
}
