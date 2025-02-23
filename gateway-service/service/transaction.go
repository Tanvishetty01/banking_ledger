package service

import (
	"errors"
	"time"

	"banking_ledger/gateway-service/config"
	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/repo"

	uuid "github.com/satori/go.uuid"
)

// ITransactionService defines the interface for transaction service
type ITransactionService interface {
	PerformTransaction(transaction *model.Transaction) error
	GetTransactionHistory(accountID string) ([]model.Transaction, error)
}

// TransactionService implements ITransactionService
type TransactionService struct {
	transactionRepo repo.TransactionRepository
	accountRepo     repo.AccountRepository
	config          config.KafkaProducer
}

// NewTransactionService creates a new instance of ITransactionService
func NewTransactionService(tr repo.TransactionRepository, ar repo.AccountRepository, c config.KafkaProducer) ITransactionService {
	return &TransactionService{
		transactionRepo: tr,
		accountRepo:     ar,
		config:          c,
	}
}

// PerformTransaction processes a transaction
func (s *TransactionService) PerformTransaction(transaction *model.Transaction) error {
	account, err := s.accountRepo.GetAccountByID(transaction.AccountID)
	if err != nil {
		return errors.New("account not found")
	}
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	transaction.TransactionID = uuid.NewV4()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	switch transaction.TransactionType {
	case "withdraw":
		if account.Amount < transaction.Amount {
			return errors.New("insufficient funds")
		}
		account.Amount -= transaction.Amount

	case "deposit":
		account.Amount += transaction.Amount

	case "transfer":
		toAccount, err := s.accountRepo.GetAccountByID(*transaction.ToAccountID)
		if err != nil {
			return errors.New("destination account not found")
		}
		if account.Amount < transaction.Amount {
			return errors.New("insufficient funds")
		}
		account.Amount -= transaction.Amount
		toAccount.Amount += transaction.Amount
		s.accountRepo.UpdateAccount(toAccount, tx)
	}

	s.accountRepo.UpdateAccount(account, tx)
	err = s.transactionRepo.SaveTransaction(transaction, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	// Send to Kafka
	s.config.ProduceTransaction(*transaction)
	return nil
}

// GetTransactionHistory fetches transaction history
func (s *TransactionService) GetTransactionHistory(accountID string) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByAccount(accountID)
}
