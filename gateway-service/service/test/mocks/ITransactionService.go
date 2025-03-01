// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	model "banking_ledger/gateway-service/models"

	mock "github.com/stretchr/testify/mock"
)

// ITransactionService is an autogenerated mock type for the ITransactionService type
type ITransactionService struct {
	mock.Mock
}

// GetTransactionHistory provides a mock function with given fields: accountID
func (_m *ITransactionService) GetTransactionHistory(accountID string) ([]model.Transaction, error) {
	ret := _m.Called(accountID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionHistory")
	}

	var r0 []model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]model.Transaction, error)); ok {
		return rf(accountID)
	}
	if rf, ok := ret.Get(0).(func(string) []model.Transaction); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerformTransaction provides a mock function with given fields: transaction
func (_m *ITransactionService) PerformTransaction(transaction *model.Transaction) error {
	ret := _m.Called(transaction)

	if len(ret) == 0 {
		panic("no return value specified for PerformTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Transaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewITransactionService creates a new instance of ITransactionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITransactionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITransactionService {
	mock := &ITransactionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
