package handler

import (
	"net/http"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/service"

	"github.com/gin-gonic/gin"
)



// ICustomerHandler defines the interface for customer handlers
type ICustomerHandler interface {
	CreateCustomer(c *gin.Context)
}

// CustomerHandler implements ICustomerHandler
type CustomerHandler struct {
	customerService service.ICustomerService
}

// NewCustomerHandler creates a new instance of CustomerHandler
func NewCustomerHandler(cs service.ICustomerService) ICustomerHandler {
	return &CustomerHandler{customerService: cs}
}

// CreateCustomer handles customer creation requests
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer

	// Bind JSON request
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create customer
	err := h.customerService.CreateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully", "customer": customer})
}