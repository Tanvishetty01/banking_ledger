package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID              uint      `json:"id"`
	TransactionID   uuid.UUID `json:"transaction_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	Notes           string    `json:"notes"`

	AccountID   string  `json:"account_id"`
	ToAccountID *string `json:"to_account_id"`
	CustomerID  string  `json:"customer_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
