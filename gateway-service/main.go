// package main

// import (
// 	"banking_ledger/gateway-service/config"
// 	"banking_ledger/gateway-service/routes"
// 	"fmt"
// )

// func main() {
// 		// Connect to Database
// 		config.ConnectDB()
// 	// Initialize router
// 	router := route.SetupRoutes()

// 	// Start server
// 	fmt.Println("🚀 Server running on port 8080")
// 	router.Run(":8080") // Gin handles HTTP server
// }

package main

import "banking_ledger/gateway-service/routes"

func main() {
	// Initialize server
	server := routes.NewServer()

	// Start the server
	server.Router.Run(":8080")
}
