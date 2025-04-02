package routes

import (
	"cashflow/pkg/handlers"
	"cashflow/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	validator := validators.NewValidator()
	handler := handlers.NewHandler(
		validator,
	)

	registerAuthRoutes(app, handler)
	registerCompanyRoutes(app, handler)
	registerRatesRoutes(app, handler)
	registerTransactionsRoutes(app, handler)
}
