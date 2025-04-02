package handlers

import (
	"cashflow/pkg/helpers"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RatesGetAll(c *fiber.Ctx) error {
	rates, err := h.rateStore.GetAll()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if len(rates) == 0 {
		return helpers.ErrorResponse(c, fiber.ErrNotFound)
	}

	return helpers.SuccessResponse(c, rates)
}

func (h *Handler) RatesGetOne(c *fiber.Ctx) error {
	currency := c.Params("currency")
	if len(currency) == 0 {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest)
	}

	rate, err := h.rateStore.GetOne(currency)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrNotFound)
	}

	return helpers.SuccessResponse(c, rate)
}
