package handler

import (
	"github.com/timofeyka25/test-jungle/internal/dto"
	"github.com/timofeyka25/test-jungle/internal/model"
	"github.com/timofeyka25/test-jungle/internal/usecase"
)

func toLoginParams(dto *dto.LoginRequestDTO) usecase.LoginParams {
	return usecase.LoginParams{
		Username: dto.Username,
		Password: dto.Password,
	}
}

func toImageParams(userId int64, imagePath, imageUrl string) usecase.ImageParams {
	return usecase.ImageParams{
		UserId:    userId,
		ImagePath: imagePath,
		ImageUrl:  imageUrl,
	}
}

func mapDtoImage(image model.Image) dto.ImageDTO {
	return dto.ImageDTO{
		Id:        image.Id,
		UserId:    image.UserId,
		ImagePath: image.ImagePath,
		ImageURL:  image.ImageURL,
	}
}

func mapDtoImages(images []model.Image) []dto.ImageDTO {
	var dtoImages []dto.ImageDTO
	for _, image := range images {
		dtoImages = append(dtoImages, mapDtoImage(image))
	}
	return dtoImages
}
