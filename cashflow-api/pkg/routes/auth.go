package routes

import (
	"cashflow/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(app *fiber.App, handler *handlers.Handler) {
	auth := app.Group("/auth")
	auth.Post("/login", handler.UserLogin)
	auth.Post("/register", handler.UserRegister)
}
