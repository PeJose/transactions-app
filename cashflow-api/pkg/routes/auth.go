package routes

import (
	"cashflow/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func registerAuth(app *fiber.App, handler *handlers.Handler) {
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
