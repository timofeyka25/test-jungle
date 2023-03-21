package server

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type Config struct {
	ReadTimeoutSeconds string
}

func FiberConfig(config Config) fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(config.ReadTimeoutSeconds)

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
