package service

import (
	"errors"
	"testing"

	"banking_ledger/gateway-service/repo/test/mocks"
	kafkaMocks "banking_ledger/gateway-service/config/mocks"
	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPerformTransaction(t *testing.T) {
	newID  := "71a93c64-5a61-4046-85aa-ad6a85d06b9d"
	tests := []struct {
		name          string
		transaction   *model.Transaction
		account       *model.Account
		toAccount     *model.Account
		repoResponse  error
		expectedError string
		shouldCallRepo bool
	}{
		{
			name: "Successful Withdrawal",
			transaction: &model.Transaction{
				AccountID:       "71a93c64-5a61-4046-85aa-ad6a85d06b9d",
				Amount:          500.0,
				TransactionType: "withdraw",
			},
			account: &model.Account{
				Amount: 1000.0,
			},
			repoResponse:  nil,
			expectedError: "",
			shouldCallRepo: true,
		},
		{
			name: "Insufficient Funds Withdrawal",
			transaction: &model.Transaction{
				AccountID:       "71a93c64-5a61-4046-85aa-ad6a85d06b9d",
				Amount:          2000.0,
				TransactionType: "withdraw",
			},
			account: &model.Account{
				Amount: 1000.0,
			},
			repoResponse:  nil,
			expectedError: "insufficient funds",
			shouldCallRepo: false,
		},
		{
			name: "Successful Transfer",
			transaction: &model.Transaction{
				AccountID:       "71a93c64-5a61-4046-85aa-ad6a85d06b9d",
				ToAccountID:     &newID,
				Amount:          300.0,
				TransactionType: "transfer",
			},
			account: &model.Account{
				Amount: 1000.0,
			},
			toAccount: &model.Account{
				Amount: 500.0,
			},
			repoResponse:  nil,
			expectedError: "",
			shouldCallRepo: true,
		},
		{
			name: "Destination Account Not Found",
			transaction: &model.Transaction{
				AccountID:       "71a93c64-5a61-4046-85aa-ad6a85d06b9d",
				ToAccountID:     &newID,
				Amount:          300.0,
				TransactionType: "transfer",
			},
			account: &model.Account{
				Amount: 1000.0,
			},
			repoResponse:  errors.New("destination account not found"),
			expectedError: "destination account not found",
			shouldCallRepo: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTransactionRepo := new(mocks.TransactionRepository)
			mockAccountRepo := new(mocks.AccountRepository)
			mockProducer := new(kafkaMocks.KafkaProducer)
			transactionService := service.NewTransactionService(mockTransactionRepo, mockAccountRepo, mockProducer)

			mockAccountRepo.On("GetAccountByID", mock.Anything).Return(tt.account, nil)
			if tt.transaction.TransactionType == "transfer" {
				mockAccountRepo.On("GetAccountByID", mock.Anything).Return(tt.toAccount, tt.repoResponse)
			}
			mockAccountRepo.On("UpdateAccount", mock.Anything, mock.Anything).Return(nil)
			mockTransactionRepo.On("SaveTransaction", mock.Anything, mock.Anything).Return(tt.repoResponse)
			mockProducer.On("ProduceTransaction", mock.Anything).Return(nil)
			

			err := transactionService.PerformTransaction(tt.transaction)

			if tt.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			}
		})
	}
}
