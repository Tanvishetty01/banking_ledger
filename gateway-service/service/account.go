package service

import (

	"banking_ledger/gateway-service/models"
	repo "banking_ledger/gateway-service/repo"

	uuid "github.com/satori/go.uuid"
)

// IAccountService defines the methods that the account service should implement
type IAccountService interface {
	CreateAccount(account *model.Account) error
}

// ServiceConfig holds dependencies for the service
type AccountService struct {
	accountRepo repo.AccountRepository
}

// NewAccountService creates a new instance of IAccountService
func NewAccountService(ar repo.AccountRepository) IAccountService {
	return &AccountService{
		accountRepo: ar,
	}	
}

// CreateAccount creates a new account
func (s *AccountService) CreateAccount(account *model.Account) error {
	account.AccountID = uuid.NewV4()
	return s.accountRepo.CreateAccount(account)
}