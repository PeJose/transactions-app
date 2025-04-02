package handlers

import (
	"cashflow/pkg/helpers"
	"cashflow/pkg/middleware"
	"cashflow/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) TransactionsGetSepa(c *fiber.Ctx) error {
	// Get the user ID from the JWT token
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized)
	}

	target := c.Query("target", "payer")

	if target != "payer" && target != "receiver" {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest)
	}

	limit := c.Query("limit", "5000")
	afterTimestamp := c.Query("after-timestamp", "")
	afterUUID := c.Query("after-uuid", "")

	params := map[string]string{
		target:            userClaims["iban"].(string),
		"limit":           limit,
		"after-timestamp": afterTimestamp,
		"after-uuid":      afterUUID,
	}

	transactions, err := h.transactionStore.GetSEPA(params)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, transactions)
}

func (h *Handler) TransactionsGetSwift(c *fiber.Ctx) error {
	// Get the user ID from the JWT token
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized)
	}

	target := c.Query("target", "sender")

	if target != "sender" && target != "beneficiary" {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest)
	}

	limit := c.Query("limit", "5000")
	afterTimestamp := c.Query("after-timestamp", "")
	afterUUID := c.Query("after-uuid", "")

	params := map[string]string{
		target:            userClaims["iban"].(string),
		"limit":           limit,
		"after-timestamp": afterTimestamp,
		"after-uuid":      afterUUID,
	}

	transactions, err := h.transactionStore.GetSWIFT(params)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError)
	}

	return helpers.SuccessResponse(c, transactions)
}

func (h *Handler) TransactionsGetBalance(c *fiber.Ctx) error {
	// Get the user ID from the JWT token
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized)
	}

	iban := userClaims["iban"].(string)

	currency := c.Query("currency")

	if !models.IsValidCurrency(currency) {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid currency")
	}

	rates, err := h.rateStore.GetAll()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	balance, err := h.transactionStore.GetBalance(iban, rates, models.Currency(currency))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, balance)
}

func (h *Handler) TransactionsGetTransactionsPreCountry(c *fiber.Ctx) error {
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized)
	}

	userIban := userClaims["iban"].(string)
	iban := c.Query("iban", userIban)

	rates, err := h.rateStore.GetAll()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	transactionsMap, err := h.transactionStore.GetTransactionPerCountries(iban, rates)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, transactionsMap)
}

func (h *Handler) TransactionsGetBalanceChange(c *fiber.Ctx) error {
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized)
	}

	userIban := userClaims["iban"].(string)
	iban := c.Query("iban", userIban)

	rates, err := h.rateStore.GetAll()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	balanceChanges, err := h.transactionStore.GetBalanceChange(iban, rates)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, balanceChanges)
}
