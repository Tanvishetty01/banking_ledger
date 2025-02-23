package routes

import (
	"banking_ledger/gateway-service/config"
	"banking_ledger/gateway-service/handler"
	"banking_ledger/gateway-service/repo"
	"banking_ledger/gateway-service/service"

	"github.com/gin-gonic/gin"
)

// Server struct holds all dependencies
type Server struct {
	Router             *gin.Engine
	TransactionHandler handler.ITransactionHandler
	AccountHandler     handler.IAccountHandler
	CustomerHandler    handler.ICustomerHandler
}

// NewServer initializes all dependencies and returns a server instance
func NewServer() *Server {
	// Initialize database connection
	config.ConnectDB()
	kafkaProducer, _ := config.NewKafkaProducer()

	// Initialize repositories
	transactionRepo := repo.NewTransactionRepo()
	accountRepo := repo.NewAccountRepo()
	customerRepo := repo.NewCustomerRepo()

	// Initialize services
	transactionService := service.NewTransactionService(transactionRepo, accountRepo, kafkaProducer)
	accountService := service.NewAccountService(accountRepo)
	customerService := service.NewCustomerService(customerRepo)

	// Initialize handlers
	transactionHandler := handler.NewTransactionHandler(transactionService)
	accountHandler := handler.NewAccountHandler(accountService)
	customerHandler := handler.NewCustomerHandler(customerService)

	// Initialize Gin router
	router := gin.Default()

	// Create server instance
	server := &Server{
		Router:             router,
		TransactionHandler: transactionHandler,
		AccountHandler:     accountHandler,
		CustomerHandler:    customerHandler,
	}

	// Register routes
	server.RegisterRoutes()

	return server
}
