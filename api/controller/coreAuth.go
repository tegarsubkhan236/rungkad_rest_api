package controller

import (
	"example/api/service"
	"example/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Error on login request", err)
	}
	identity := input.Identity
	pass := input.Password

	email, err := service.GetUserByEmail(identity)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusUnauthorized, "Error on email", err)
	}

	user, err := service.GetUserByUsername(identity)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusUnauthorized, "Error on username", err)
	}

	if email == nil && user == nil {
		return config.ResponseHandler(c, fiber.StatusUnauthorized, "User not found", nil)
	}

	if email.ID != 0 {
		ud = UserData{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	} else if user.ID != 0 {
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		return config.ResponseHandler(c, fiber.StatusNotFound, "User not found", nil)
	}

	if !service.CheckPasswordHash(pass, ud.Password) {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Invalid password", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	t, err := token.SignedString([]byte(config.GetEnv("SECRET")))
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Error on server", err)
	}

	return config.ResponseHandler(c, fiber.StatusOK, "Success login", t)
}
