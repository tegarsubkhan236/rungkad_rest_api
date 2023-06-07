package controller

import (
	"example/api/service"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetSuppliers(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	data, count, err := service.GetAllSupplier(offset, intLimit)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func GetSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := service.GetSupplierById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No Supplier find with ID "+id, nil)
	}
	return config.ResponseHandler(c, fiber.StatusFound, "Supplier Found", item)
}

func GetSupplierByColumn(c *fiber.Ctx) error {
	var payload *model.InvSupplier
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	if payload.Name == "" {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "name cannot be empty", nil)
	}
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	data, count, err := service.GetSupplierByColumn(payload, offset, intLimit)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	if len(data) == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Data not found", nil)
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func CreateSupplier(c *fiber.Ctx) error {
	var payload *model.InvSupplier
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	result, err := service.CreateSupplier(payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Couldn't create supplier", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusCreated, "Created Supplier", result)
}

func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var payload *model.InvSupplier
	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	item, err := service.GetSupplierById(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No Supplier find with ID "+id, nil)
	}
	result, err := service.UpdateSupplier(item, payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Failed to update data", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "Success update data", fiber.Map{"item": result})
}

func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	result := service.DestroySupplier(id)
	if result.RowsAffected == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No Supplier find with ID "+id, nil)
	} else if result.Error != nil {
		return config.ResponseHandler(c, fiber.StatusBadGateway, "error", result.Error.Error())
	}
	return config.ResponseHandler(c, fiber.StatusNoContent, "", nil)
}
