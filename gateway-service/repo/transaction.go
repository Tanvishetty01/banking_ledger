package repo

import (
	"banking_ledger/gateway-service/config"
	"banking_ledger/gateway-service/models"
	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction data operations
type TransactionRepository interface {
	SaveTransaction(transaction *model.Transaction, tx *gorm.DB) error
	GetTransactionsByAccount(accountID string) ([]model.Transaction, error)
}

// TransactionRepo implements TransactionRepository
type TransactionRepo struct{}

// NewTransactionRepo creates a new instance of TransactionRepo
func NewTransactionRepo() TransactionRepository {
	return &TransactionRepo{}
}

// SaveTransaction inserts a new transaction into the database
func (r *TransactionRepo) SaveTransaction(transaction *model.Transaction, tx *gorm.DB) error {
	err := config.DB.Create(transaction).Error
	if err != nil {	
		tx.Rollback()
        return err
	}
	return nil
}

// GetTransactionsByAccount fetches transaction history for a given account
func (r *TransactionRepo) GetTransactionsByAccount(accountID string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := config.DB.Where("account_id = ?", accountID).Find(&transactions).Error
	return transactions, err
}
