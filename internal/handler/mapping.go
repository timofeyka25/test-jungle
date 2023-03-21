package handler

import (
	"github.com/timofeyka25/test-jungle/internal/dto"
	"github.com/timofeyka25/test-jungle/internal/usecase"
)

func toLoginParams(dto *dto.LoginRequestDTO) usecase.LoginParams {
	return usecase.LoginParams{
		Username: dto.Username,
		Password: dto.Password,
	}
}
