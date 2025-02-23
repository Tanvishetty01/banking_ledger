package main

import (
	"banking_ledger/gateway-service/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize server
	server := routes.NewServer()

	err := godotenv.Load() // Load .env file
	if err != nil {
		fmt.Println("⚠️ Warning: No .env file found, using system environment variables")
	}
	// Start the server
	server.Router.Run(":8080")
}
