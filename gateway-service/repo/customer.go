package repo

import (
	"banking_ledger/gateway-service/config"
	"banking_ledger/gateway-service/models"
)

// CustomerRepository defines the interface for customer data operations
type CustomerRepository interface {
	CreateCustomer(customer *model.Customer) error
	GetCustomerByID(customerID string) (*model.Customer, error)
}

// CustomerRepo implements CustomerRepository
type CustomerRepo struct{}

// NewCustomerRepo creates a new instance of CustomerRepo
func NewCustomerRepo() CustomerRepository {
	return &CustomerRepo{}
}

// CreateCustomer saves a new customer in the database
func (r *CustomerRepo) CreateCustomer(customer *model.Customer) error {
	return config.DB.Create(customer).Error
}

// GetCustomerByID fetches a customer by ID
func (r *CustomerRepo) GetCustomerByID(customerID string) (*model.Customer, error) {
	var customer model.Customer
	err := config.DB.First(&customer, "id = ?", customerID).Error
	return &customer, err
}
