package server

import (
	"cashflow/pkg/config"
	"cashflow/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Start() {
	app := fiber.New()

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		println("Request:", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(cors.New())
	app.Use(helmet.New())

	// Register routes
	routes.RegisterRoutes(app)

	// List all routes on startup
	listRoutes(app)

	// Start the server
	if err := app.Listen(":" + config.Port); err != nil {
		panic(err)
	}
}

func listRoutes(app *fiber.App) {
	println("Registered Routes:")
	for _, route := range app.Stack() {
		for _, r := range route {
			println(r.Method, r.Path)
		}
	}
}
