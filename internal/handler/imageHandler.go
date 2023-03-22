package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/timofeyka25/test-jungle/internal/dto"
	"github.com/timofeyka25/test-jungle/internal/usecase"
	"github.com/timofeyka25/test-jungle/pkg/cloudstorage"
	"github.com/timofeyka25/test-jungle/pkg/drivestorage"
	"strconv"
)

type ImageHandler struct {
	useCase           usecase.ImageUseCaseInterface
	cloudFileUploader *cloudstorage.FileUploader
	driveFileUploader *drivestorage.FileUploader
}

func NewImageHandler(
	useCase usecase.ImageUseCaseInterface,
	cloudFileUploader *cloudstorage.FileUploader,
	driveFileUploader *drivestorage.FileUploader,
) ImageHandler {
	return ImageHandler{
		useCase:           useCase,
		cloudFileUploader: cloudFileUploader,
		driveFileUploader: driveFileUploader,
	}
}

func (h *ImageHandler) UploadPicture(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// get query param to choose targeted storage
	storage := ctx.Query("storage", "google-drive")

	// get user id
	id := ctx.Cookies("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var fileURL string

	// upload file
	switch storage {
	case "google-drive":
		fileURL, err = h.driveFileUploader.Upload(file)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError,
				fmt.Errorf("error while uploading to google drive storage: %s", err.Error()).Error())
		}
	case "google-cloud":
		fileURL, err = h.cloudFileUploader.Upload(ctx.Context(), file)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError,
				fmt.Errorf("error while uploading to google cloud storage: %s", err.Error()).Error())
		}
	}

	// add image info to db
	_, err = h.useCase.SaveImage(toImageParams(userId, file.Filename, fileURL))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError,
			fmt.Errorf("error while writing to db: %s", err.Error()).Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.UploadedFileDTO{
		FileURL: fileURL,
	})
}

func (h *ImageHandler) GetImages(ctx *fiber.Ctx) error {
	// get user id
	id := ctx.Cookies("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	images, err := h.useCase.GetImages(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(mapDtoImages(images))
}
