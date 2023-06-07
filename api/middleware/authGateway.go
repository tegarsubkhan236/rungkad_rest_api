package middleware

import (
	"example/pkg/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetEnv("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusUnauthorized, "Invalid or expired JWT", nil)
}
