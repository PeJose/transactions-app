package routes

import (
	"cashflow/pkg/handlers"
	"cashflow/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerTransactions(app *fiber.App, handler *handlers.Handler) {
	transactions := app.Group("/transactions", middleware.Protected())
	transactions.Get("/balance", handler.TransactionsGetBalance)
	transactions.Get("/map", handler.TransactionsGetTransactionsPreCountry)
	transactions.Get("/sepa", handler.TransactionsGetSepa)
	transactions.Get("/swift", handler.TransactionsGetSwift)
	transactions.Get("/chart/balance", handler.TransactionsGetBalanceChange)
}
