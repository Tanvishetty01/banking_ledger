package handler

import (
	"net/http"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/service"

	"github.com/gin-gonic/gin"
)

// ITransactionHandler defines the interface for transaction handlers
type ITransactionHandler interface {
	PerformTransaction(c *gin.Context)
	GetTransactionHistory(c *gin.Context)
}

// TransactionHandler implements ITransactionHandler
type TransactionHandler struct {
	transactionService service.ITransactionService
}

// NewTransactionHandler creates a new instance of TransactionHandler
func NewTransactionHandler(ts service.ITransactionService) ITransactionHandler {
	return &TransactionHandler{transactionService: ts}
}

// PerformTransaction handles transaction processing
func (h *TransactionHandler) PerformTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.transactionService.PerformTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetTransactionHistory handles retrieving transactions for an account
func (h *TransactionHandler) GetTransactionHistory(c *gin.Context) {
	accountID := c.Param("account_id")
	transactions, err := h.transactionService.GetTransactionHistory(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
