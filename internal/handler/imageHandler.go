package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timofeyka25/test-jungle/internal/usecase"
)

type ImageHandler struct {
	useCase usecase.ImageUseCaseInterface
}

func NewImageHandler(useCase usecase.ImageUseCaseInterface) ImageHandler {
	return ImageHandler{useCase: useCase}
}

func (h *ImageHandler) UploadPicture(ctx *fiber.Ctx) error {
	return nil
}

func (h *ImageHandler) GetImages(ctx *fiber.Ctx) error {
	return nil
}
