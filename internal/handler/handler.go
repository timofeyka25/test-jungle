package handler

import "github.com/gofiber/fiber/v2"

type Handler struct {
	imageHandler ImageHandler
	userHandler  UserHandler
}

func NewHandler(
	userHandler UserHandler,
	imageHandler ImageHandler) *Handler {
	return &Handler{
		userHandler:  userHandler,
		imageHandler: imageHandler,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Post("/login", h.userHandler.Login)
	app.Post("/upload-picture", h.imageHandler.UploadPicture)
	app.Get("/images", h.imageHandler.GetImages)
}
