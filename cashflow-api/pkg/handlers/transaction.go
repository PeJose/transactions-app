package handlers

import (
	"cashflow/pkg/helpers"
	"cashflow/pkg/middleware"
	"cashflow/pkg/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) TransactionGetSepa(c *fiber.Ctx) error {
	iban, err := helpers.GetIban(c)
	if err != nil {
		return err
	}

	target := c.Query("target", "from")

	if target != "to" && target != "from" {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest)
	}

	if target == "from" {
		target = "payer"
	} else if target == "to" {
		target = "receiver"
	}

	limit := c.Query("limit", "10")
	cursor := c.Query("cursor", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid limit value")
	}

	cursorParts := strings.Split(cursor, "$$")
	afterUUID := ""
	afterTimestamp := ""
	if len(cursorParts) > 0 {
		afterUUID = cursorParts[0]
	}
	if len(cursorParts) > 1 {
		afterTimestamp = cursorParts[1]
	}

	params := map[string]string{
		target:            iban,
		"limit":           strconv.Itoa(limitInt + 1),
		"after-uuid":      afterUUID,
		"after-timestamp": afterTimestamp,
	}

	transactions, err := h.transactionStore.GetSEPA(params)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if len(transactions) == 0 {
		return helpers.SuccessResponse(c, fiber.Map{
			"hasNext":      false,
			"transactions": "",
		})
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"hasNext":      len(transactions) > limitInt,
		"transactions": transactions[:len(transactions)-1],
		"lastCursor": func() string {
			if len(transactions) > 1 {
				return transactions[len(transactions)-1].ID + "$$" + transactions[len(transactions)-1].Timestamp.Format(time.RFC3339)
			}
			return ""
		}(),
	})
}

func (h *Handler) TransactionGetSwift(c *fiber.Ctx) error {
	iban, err := helpers.GetIban(c)
	if err != nil {
		return err
	}

	target := c.Query("target", "from")

	if target != "from" && target != "to" {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest)
	}

	if target == "from" {
		target = "sender"
	} else if target == "to" {
		target = "beneficiary"
	}

	limit := c.Query("limit", "10")
	cursor := c.Query("cursor", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid limit value")
	}

	cursorParts := strings.Split(cursor, "$$")
	afterUUID := ""
	afterTimestamp := ""
	if len(cursorParts) > 0 {
		afterUUID = cursorParts[0]
	}
	if len(cursorParts) > 1 {
		afterTimestamp = cursorParts[1]
	}

	params := map[string]string{
		target:            iban,
		"limit":           strconv.Itoa(limitInt + 1),
		"after-uuid":      afterUUID,
		"after-timestamp": afterTimestamp,
	}

	transactions, err := h.transactionStore.GetSWIFT(params)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	if len(transactions) == 0 {
		return helpers.SuccessResponse(c, fiber.Map{
			"hasNext":      false,
			"transactions": "",
		})
	}

	return helpers.SuccessResponse(c, fiber.Map{
		"hasNext":      len(transactions) > limitInt,
		"transactions": transactions[:len(transactions)-1],
		"lastCursor":   transactions[len(transactions)-1].ID + "$$" + transactions[len(transactions)-1].Timestamp.Format(time.RFC3339),
	})
}

func (h *Handler) TransactionGetBalance(c *fiber.Ctx) error {
	iban, err := helpers.GetIban(c)
	if err != nil {
		return err
	}

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

func (h *Handler) TransactionGetPerCountry(c *fiber.Ctx) error {
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

	transactionsMap, err := h.transactionStore.GetPerCountries(iban, rates)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, transactionsMap)
}

func (h *Handler) TransactionGetBalanceChange(c *fiber.Ctx) error {
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
