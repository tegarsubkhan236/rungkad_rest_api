package controller

import (
	"example/api/service"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProductFilter struct {
	Page            int    `query:"page"`
	Limit           int    `query:"limit"`
	SupplierID      int    `query:"supplier_id"`
	ProductCategory []int  `query:"product_category"`
	SearchText      string `query:"search_text"`
}

func GetProducts(c *fiber.Ctx) error {
	var productFilter = new(ProductFilter)
	if err := c.QueryParser(productFilter); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	if productFilter.Page == 0 {
		productFilter.Page = 1
	}
	if productFilter.Limit == 0 {
		productFilter.Limit = 10
	}
	offset := (productFilter.Page - 1) * productFilter.Limit
	data, count, err := service.GetAllProduct(offset, productFilter.Limit)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func GetProductsByFilter(c *fiber.Ctx) error {
	var productFilter = new(ProductFilter)
	if err := c.QueryParser(productFilter); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	if productFilter.Page == 0 {
		productFilter.Page = 1
	}
	if productFilter.Limit == 0 {
		productFilter.Limit = 10
	}

	offset := (productFilter.Page - 1) * productFilter.Limit
	data, count, err := service.GetAllProductByFilter(offset, productFilter.Limit, productFilter.SupplierID, productFilter.ProductCategory, productFilter.SearchText)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", fiber.Map{"results": data, "total": count})
}

func GetProduct(c *fiber.Ctx) error {
	var id = c.Params("id")
	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}
	item, err := service.GetProductById(uint(num))
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No Product find with ID "+id, nil)
	}
	return config.ResponseHandler(c, fiber.StatusFound, "Product Found", item)
}

func CreateProduct(c *fiber.Ctx) error {
	var payload *model.InvProduct

	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	createResult, err := service.CreateProduct(payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Couldn't create supplier", err.Error())
	}
	findResult, err := service.GetProductById(createResult.ID)
	if err != nil {
		return err
	}
	return config.ResponseHandler(c, fiber.StatusCreated, "Created Product", findResult)
}

func UpdateProduct(c *fiber.Ctx) error {
	var id = c.Params("id")
	var payload *model.InvProduct

	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	}

	if err := c.BodyParser(&payload); err != nil {
		return config.ResponseHandler(c, fiber.StatusBadRequest, "Review your input", err.Error())
	}
	item, err := service.GetProductById(uint(num))
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "", err.Error())
	} else if item.ID == 0 {
		return config.ResponseHandler(c, fiber.StatusNotFound, "No Product find with ID "+id, nil)
	}
	result, err := service.UpdateProduct(item, payload)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusNotFound, "Failed to update data", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "Success update data", fiber.Map{"user": result})
}

func DeleteProduct(c *fiber.Ctx) error {
	var id = c.Params("id")

	err := service.DestroyProduct(id)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusBadGateway, "error", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusNoContent, "", nil)
}
