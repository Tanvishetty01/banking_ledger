package service

import (
	"errors"
	"testing"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/service"
	"banking_ledger/gateway-service/repo/test/mocks"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name          string
		account       *model.Account
		repoResponse  error
		expectedError string
		shouldCallRepo bool
	}{
		{
			name: "Success",
			account: &model.Account{
				CustomerID: uuid.NewV4(),
				Amount:     1000.0,
			},
			repoResponse:  nil,
			expectedError: "",
			shouldCallRepo: true,
		},
		{
			name: "Repository Error",
			account: &model.Account{
				CustomerID: uuid.NewV4(),
				Amount:     1000.0,
			},
			repoResponse:  errors.New("database error"),
			expectedError: "database error",
			shouldCallRepo: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.AccountRepository)
			accountService := service.NewAccountService(mockRepo)

			if tt.shouldCallRepo {
				mockRepo.On("CreateAccount", mock.Anything).Return(tt.repoResponse)
			}

			err := accountService.CreateAccount(tt.account)

			if tt.expectedError == "" {
				assert.Nil(t, err)
				assert.NotEqual(t, uuid.Nil, tt.account.ID)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			}

			// mockRepo.AssertExpectations(t)
		})
	}
}
