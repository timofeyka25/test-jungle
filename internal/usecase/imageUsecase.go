package usecase

import "github.com/timofeyka25/test-jungle/internal/repository"

type ImageUseCaseInterface interface {
}

type ImageUseCase struct {
	repository repository.ImageRepositoryInterface
}

func NewImageUseCase(repository repository.ImageRepositoryInterface) ImageUseCase {
	return ImageUseCase{repository: repository}
}
