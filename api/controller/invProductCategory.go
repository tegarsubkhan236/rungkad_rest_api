package controller

import (
	"example/api/service"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetProductCategories(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	data, count, err := service.GetAllProductCategory(offset, intLimit)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func GetProductCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := service.GetProductCategoryById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No ProductCategory find with ID "+id, nil)
	}
	return config.ResponseHandler(c, fiber.StatusFound, "ProductCategory Found", item)
}

func CreateProductCategory(c *fiber.Ctx) error {
	var payload *model.InvProductCategory
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	result, err := service.CreateProductCategory(payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Couldn't create supplier", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusCreated, "Created ProductCategory", result)
}

func UpdateProductCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *model.InvProductCategory
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	item, err := service.GetProductCategoryById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No ProductCategory find with ID "+id, nil)
	}
	result, err := service.UpdateProductCategory(item, payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Failed to update data", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "Success update data", fiber.Map{"user": result})
}

func DeleteProductCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	result := service.DestroyProductCategory(id)
	if result.RowsAffected == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No ProductCategory find with ID "+id, nil)
	} else if result.Error != nil {
		return config.ResponseHandler(c, fiber.StatusBadGateway, "error", result.Error.Error())
	}
	return config.ResponseHandler(c, fiber.StatusNoContent, "", nil)
}
