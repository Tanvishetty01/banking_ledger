package handler

import (
	"errors"
	"net/http"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/service"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// IAccountHandler defines the interface for account handlers
type IAccountHandler interface {
	CreateAccount(c *gin.Context)
}

// AccountHandler implements IAccountHandler
type AccountHandler struct {
	accountService service.IAccountService
}

// NewAccountHandler creates a new instance of AccountHandler
func NewAccountHandler(as service.IAccountService) IAccountHandler {
	return &AccountHandler{accountService: as}
}

// CreateAccount handles the HTTP request for creating an account
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate account data
	if err := validateAccount(account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.accountService.CreateAccount(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created", "account": account})
}


func validateAccount(account model.Account) error {
	if account.AccountID == uuid.Nil {
		return errors.New("AccountID is required")
	}
	if account.CustomerID == uuid.Nil {
		return errors.New("CustomerID is required")
	}
	if len(account.NickName) < 3 {
		return errors.New("NickName must be at least 3 characters long")
	}
	if account.Amount < 0 {
		return errors.New("Amount cannot be negative")
	}
	return nil
}
