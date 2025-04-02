package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessMessageKey string

var SuccessMessageKeys = struct {
	DefaultMessage  SuccessMessageKey
	LoginMessage    SuccessMessageKey
	RegisterMessage SuccessMessageKey
}{
	DefaultMessage:  "default",
	LoginMessage:    "login",
	RegisterMessage: "register",
}

var SuccessMessages = map[SuccessMessageKey]string{
	SuccessMessageKeys.DefaultMessage:  "Operation successful",
	SuccessMessageKeys.LoginMessage:    "Login successful",
	SuccessMessageKeys.RegisterMessage: "Registration successful",
}

func SuccessResponse(c *fiber.Ctx, data interface{}, optionalMsg ...SuccessMessageKey) error {
	message := SuccessMessages["default"]
	if len(optionalMsg) > 0 && optionalMsg[0] != "" {
		message = SuccessMessages[optionalMsg[0]]
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *fiber.Ctx, err *fiber.Error, optionalErrMsg ...string) error {
	message := err.Message
	if len(optionalErrMsg) > 0 && optionalErrMsg[0] != "" {
		message = optionalErrMsg[0]
	}

	return c.Status(err.Code).JSON(fiber.Map{
		"status": "error",
		"error":  message,
	})
}
