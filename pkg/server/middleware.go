package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timofeyka25/test-jungle/pkg/jwt"
	"strconv"
	"strings"
)

type tokenValidatorMiddleware struct {
	tokenValidator *jwt.TokenValidator
}

func NewTokenValidatorMiddleware(tokenValidator *jwt.TokenValidator) *tokenValidatorMiddleware {
	return &tokenValidatorMiddleware{tokenValidator: tokenValidator}
}
func (m *tokenValidatorMiddleware) validateToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid authorization header")
	}
	token := parts[1]
	parsedParams, err := m.tokenValidator.ValidateToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	ctx.Cookie(&fiber.Cookie{Name: "id", Value: strconv.FormatInt(parsedParams.Id, 10)})

	return ctx.Next()
}
