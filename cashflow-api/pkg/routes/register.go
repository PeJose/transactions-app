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

	registerAuth(app, handler)
	registerCompanies(app, handler)
	registerRates(app, handler)
	registerTransactions(app, handler)
}
