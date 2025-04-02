package handlers

import (
	"cashflow/pkg/helpers"
	"cashflow/pkg/middleware"
	"cashflow/pkg/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	requestUser := new(models.LoginUser)
	if err := c.BodyParser(requestUser); err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid request")
	}

	println(requestUser.Email)

	if err := h.validator.Validate(requestUser); err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, h.validator.Format(err))
	}

	user, err := h.userStore.GetByEmail(requestUser.Email)
	if err != nil || user == nil {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized, "Invalid credentials")
	}

	if !user.CheckPassword(requestUser.Password) {
		return helpers.ErrorResponse(c, fiber.ErrUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, "Failed to generate token")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt-token",
		Value: token,
	})

	return helpers.SuccessResponse(c, token, helpers.SuccessMessageKeys.LoginMessage)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "Invalid request")
	}

	if err := h.validator.Validate(newUser); err != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, h.validator.Format(err))
	}

	existingUser, err := h.userStore.GetByEmail(newUser.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, "Failed to check existing user")
	}
	if existingUser != nil {
		return helpers.ErrorResponse(c, fiber.ErrBadRequest, "User already exists")
	}

	if err := newUser.HashPassword(); err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, "Failed to hash password")
	}

	user, err := h.userStore.Create(newUser)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.ErrInternalServerError, "Failed creating new user")
	}

	return helpers.SuccessResponse(c, user, helpers.SuccessMessageKeys.RegisterMessage)
}
