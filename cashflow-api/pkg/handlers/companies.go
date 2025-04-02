package handlers

import (
	"cashflow/pkg/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jbub/banking/iban"
)

func (h *Handler) CompanyGetAll(c *fiber.Ctx) error {
	afterId, err := strconv.Atoi(c.Query("after-id", "0"))
	if err != nil || afterId < 0 {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid page parameter")
	}

	limit, err := strconv.Atoi(c.Query("limit", "0"))
	if err != nil || limit < 0 {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid page parameter")
	}

	companies, err := h.companyStore.GetAll(afterId, limit)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if len(companies) == 0 {
		return helpers.ErrorResponse(c, fiber.ErrNotFound, "No companies found")
	}

	return helpers.SuccessResponse(c, companies)
}

func (h *Handler) CompanyGetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid ID parameter")
	}

	company, err := h.companyStore.GetById(id)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if company == nil {
		return helpers.ErrorResponse(c, fiber.ErrNotFound, "Company not found")
	}

	return helpers.SuccessResponse(c, company)
}

func (h *Handler) CompanyGetByIban(c *fiber.Ctx) error {
	ibanN := c.Params("iban", "")
	err := iban.Validate(ibanN)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid iban parameter")
	}

	company, err := h.companyStore.GetByIban(ibanN)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if company == nil {
		return helpers.ErrorResponse(c, fiber.ErrNotFound, "Company not found")
	}

	return helpers.SuccessResponse(c, company)
}
