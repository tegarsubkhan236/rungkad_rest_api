package controller

import (
	"example/api/service"
	"example/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func BalanceCheck(c *fiber.Ctx) error {
	body, err := service.HitBalanceSCheck()
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Error on server", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", body.Data)
}

func PriceList(c *fiber.Ctx) error {
	var code = c.Query("code", "")
	var priceType = c.Query("type", "")
	var groupBy = c.Query("groupBy", "")

	body, err := service.HitPriceList(priceType, code)
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Error on server", err.Error())
	}

	if groupBy != "" {
		body, err := service.PriceListGrouped(body.Data, groupBy)
		if err != nil {
			return config.ResponseHandler(c, fiber.StatusBadRequest, "Bad request", err.Error())
		} else {
			return config.ResponseHandler(c, fiber.StatusOK, "", body)
		}
	}

	return config.ResponseHandler(c, fiber.StatusOK, "", body.Data)
}

func Deposit(c *fiber.Ctx) error {
	body, err := service.HitDeposit()
	if err != nil {
		return config.ResponseHandler(c, fiber.StatusInternalServerError, "Error on server", err.Error())
	}
	return config.ResponseHandler(c, fiber.StatusOK, "", body.Data)
}
