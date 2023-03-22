package usecase

import (
	"github.com/timofeyka25/test-jungle/internal/model"
	"github.com/timofeyka25/test-jungle/internal/repository"
)

type ImageUseCaseInterface interface {
	SaveImage(params ImageParams) (int64, error)
	GetImages(userId int64) ([]model.Image, error)
}

type ImageUseCase struct {
	repository repository.ImageRepositoryInterface
}

func NewImageUseCase(repository repository.ImageRepositoryInterface) ImageUseCase {
	return ImageUseCase{repository: repository}
}

func (uc ImageUseCase) SaveImage(params ImageParams) (int64, error) {
	image := model.NewImage(params.UserId, params.ImagePath, params.ImageUrl)
	err := uc.repository.AddImage(image)
	if err != nil {
		return 0, err
	}
	return image.Id, nil
}

func (uc ImageUseCase) GetImages(userId int64) ([]model.Image, error) {
	return uc.repository.GetImages(userId)
}

type ImageParams struct {
	UserId    int64
	ImagePath string
	ImageUrl  string
}
