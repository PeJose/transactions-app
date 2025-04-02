package routes

import (
	"cashflow/pkg/handlers"
	"cashflow/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerTransactionsRoutes(app *fiber.App, handler *handlers.Handler) {
	transactions := app.Group("/transactions", middleware.Protected())
	transactions.Get("/balance", handler.TransactionGetBalance)
	transactions.Get("/sepa", handler.TransactionGetSepa)
	transactions.Get("/swift", handler.TransactionGetSwift)

	charts := transactions.Group("/charts")
	charts.Get("/balance", handler.TransactionGetBalanceChange)
	charts.Get("/map", handler.TransactionGetPerCountry)
}
