package routes

func (s *Server) RegisterRoutes() {
	// Transaction routes
	s.Router.POST("/transaction", s.TransactionHandler.PerformTransaction)
	s.Router.GET("/transactions/:account_id", s.TransactionHandler.GetTransactionHistory)

	// Account routes
	s.Router.POST("/account", s.AccountHandler.CreateAccount)

	// Customer routes
	s.Router.POST("/customers", s.CustomerHandler.CreateCustomer)
}
