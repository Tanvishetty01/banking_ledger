package repo

import (
	"banking_ledger/gateway-service/config"
	model "banking_ledger/gateway-service/models"

	"gorm.io/gorm"
)

// AccountRepository defines the interface for account data operations
type AccountRepository interface {
	CreateAccount(account *model.Account) error
	GetAccountByID(accountID string) (*model.Account, error)
	UpdateAccount(account *model.Account, tx *gorm.DB) error
}

// AccountRepo implements AccountRepository
type AccountRepo struct{}

// NewAccountRepo creates a new instance of AccountRepo
func NewAccountRepo() AccountRepository {
	return &AccountRepo{}
}

// CreateAccount saves a new account in the database
func (r *AccountRepo) CreateAccount(account *model.Account) error {
	return config.DB.Create(account).Error
}

// GetAccountByID fetches an account by its ID
func (r *AccountRepo) GetAccountByID(accountID string) (*model.Account, error) {
	var account model.Account
	err := config.DB.First(&account, "account_id", accountID).Error
	return &account, err
}

// UpdateAccount updates the account balance
func (r *AccountRepo) UpdateAccount(account *model.Account, tx *gorm.DB) error {
	err := config.DB.Save(account).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
