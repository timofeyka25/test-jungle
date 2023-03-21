package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/timofeyka25/test-jungle/internal/dto"
	"github.com/timofeyka25/test-jungle/internal/usecase"
)

type UserHandler struct {
	useCase usecase.UserUseCaseInterface
}

func NewUserHandler(useCase usecase.UserUseCaseInterface) UserHandler {
	return UserHandler{useCase: useCase}
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	loginDto := new(dto.LoginRequestDTO)
	if err := ctx.BodyParser(&loginDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Error body parsing %s", err.Error()))
	}
	token, err := h.useCase.Login(toLoginParams(loginDto))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(dto.LoginResponseDTO{Token: token})
}
