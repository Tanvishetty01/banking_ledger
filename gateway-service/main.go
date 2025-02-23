package main

import "banking_ledger/gateway-service/routes"

func main() {
	// Initialize server
	server := routes.NewServer()

	// Start the server
	server.Router.Run(":8080")
}
