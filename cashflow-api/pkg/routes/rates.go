package routes

import (
	"cashflow/pkg/handlers"
	"cashflow/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerRatesRoutes(app *fiber.App, handler *handlers.Handler) {
	rates := app.Group("/rates", middleware.Protected())
	rates.Get("/", handler.RatesGetAll)
	rates.Get("/:currency", handler.RatesGetOne)
}
