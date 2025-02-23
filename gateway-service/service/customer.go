package service

import (
	"time"

	"banking_ledger/gateway-service/models"
	"banking_ledger/gateway-service/repo"

	uuid "github.com/satori/go.uuid"
)

// ICustomerService defines the interface for customer service
type ICustomerService interface {
	CreateCustomer(customer *model.Customer) error
	GetCustomerByID(customerID string) (*model.Customer, error)
}

// CustomerService implements ICustomerService
type CustomerService struct {
	customerRepo repo.CustomerRepository
}

// NewCustomerService creates a new instance of ICustomerService
func NewCustomerService(cr repo.CustomerRepository) ICustomerService {
	return &CustomerService{
		customerRepo: cr,
	}
}

// CreateCustomer handles creating a new customer
func (s *CustomerService) CreateCustomer(customer *model.Customer) error {
	customer.CustomerID = uuid.NewV4()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	return s.customerRepo.CreateCustomer(customer)
}

// GetCustomerByID fetches a customer by ID
func (s *CustomerService) GetCustomerByID(customerID string) (*model.Customer, error) {
	return s.customerRepo.GetCustomerByID(customerID)
}
