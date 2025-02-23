package service

import (
	"errors"
	"time"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/repo"
	"banking_ledger/gateway-service/config"

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
	config config.KafkaProducer
}

// NewTransactionService creates a new instance of ITransactionService
func NewTransactionService(tr repo.TransactionRepository, ar repo.AccountRepository, c config.KafkaProducer) ITransactionService {
	return &TransactionService{
		transactionRepo: tr,
		accountRepo:     ar,
		config: c,
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

// func (s *TransactionService) PerformTransaction(transaction *model.Transaction) error {
// 	// 1️⃣ Validate Transaction
// 	if err := s.validateTransaction(transaction); err != nil {
// 		return err
// 	}

// 	// 2️⃣ Start a Single Database Transaction
// 	tx := config.DB.Begin()
// 	defer tx.Commit() // ✅ Rollback if commit is not called

// 	// 3️⃣ Fetch Account (Inside Transaction)
// 	account, err := tx.GetAccountByID(transaction.AccountID)
// 	if err != nil {
// 		return errors.New("account not found")
// 	}

// 	// 4️⃣ Process Transaction Logic
// 	if err := s.processTransaction(transaction, account, tx); err != nil {
// 		return err
// 	}

// 	// 5️⃣ Save Transaction Record
// 	if err := tx.SaveTransaction(transaction); err != nil {
// 		return err
// 	}

// 	// 6️⃣ Commit Database Transaction ✅ (Ensures Atomicity)
// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}

// 	// 7️⃣ Publish to Kafka (Decoupled from DB Transaction)
// 	go s.config.ProduceTransaction(transaction)

// 	return nil
// }

// GetTransactionHistory fetches transaction history
func (s *TransactionService) GetTransactionHistory(accountID string) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByAccount(accountID)
}


// func (s *TransactionService) processTransaction(transaction *model.Transaction, account *model.Account, tx repository.Transaction) error {
// 	switch transaction.TransactionType {
// 	case "withdraw":
// 		if account.Amount < transaction.Amount {
// 			return errors.New("insufficient funds")
// 		}
// 		account.Amount -= transaction.Amount

// 	case "deposit":
// 		account.Amount += transaction.Amount

// 	case "transfer":
// 		toAccount, err := tx.GetAccountByID(*transaction.ToAccountID)
// 		if err != nil {
// 			return errors.New("destination account not found")
// 		}
// 		if account.Amount < transaction.Amount {
// 			return errors.New("insufficient funds")
// 		}
// 		account.Amount -= transaction.Amount
// 		toAccount.Amount += transaction.Amount
// 		if err := tx.UpdateAccount(toAccount); err != nil {
// 			return err
// 		}
// 	}

// 	return tx.UpdateAccount(account) // ✅ Uses the same transaction
// }
