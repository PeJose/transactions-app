package main

import (
	"cashflow/pkg/config"
	"cashflow/pkg/db"
	"cashflow/pkg/server"
)

func main() {
	// Load configuration
	config.Load()

	// Connect to the database
	db.Connect()

	// Start the server
	server.Start()
}
