package routes

import (
	"cashflow/pkg/handlers"
	"cashflow/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerCompanies(app *fiber.App, handler *handlers.Handler) {
	companies := app.Group("/companies", middleware.Protected())
	companies.Get("/", handler.CompanyGetAll)
	companies.Get("/:id", handler.CompanyGetById)
	companies.Get("/iban/:iban", handler.CompanyGetByIban)
}
