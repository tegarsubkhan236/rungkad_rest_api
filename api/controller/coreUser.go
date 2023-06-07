package controller

import (
	"example/api/service"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetUsers(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	data, count, err := service.GetAllUser(offset, intLimit)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := service.GetUserById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if user.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No CoreUser find with ID", nil)
	}
	return config.ResponseHandler(c, fiber.StatusFound, "CoreUser Found", user)
}

func CreateUser(c *fiber.Ctx) error {
	var payload *model.CoreUser
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	hash, err := service.HashPassword(payload.Password)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Couldn't hash password", err.Error())
	}
	payload.Password = hash
	result, err := service.CreateUser(payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Couldn't create user", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusCreated, "Created user", result)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *model.CoreUser
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	user, err := service.GetUserById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No CoreUser find with ID", err.Error())
	}
	result, err := service.UpdateUser(user, payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Failed to update data", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "Success update data", fiber.Map{"user": result})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	result := service.DestroyUser(id)
	if result.RowsAffected == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No user with that Id exists", nil)
	} else if result.Error != nil {
		return config.ResponseHandler(c, fiber.StatusBadGateway, "error", result.Error.Error())
	}
	return config.ResponseHandler(c, fiber.StatusNoContent, "", nil)
}
