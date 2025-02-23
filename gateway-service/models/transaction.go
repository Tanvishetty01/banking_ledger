package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

// Transaction model
type Transaction struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID   uuid.UUID `gorm:"type:char(36);unique;not null" json:"transaction_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	TransactionType string    `gorm:"not null" json:"transaction_type"`
	Notes           string    `json:"notes"`

	AccountID   string  `gorm:"not null" json:"account_id"`
	ToAccountID *string `gorm:"null" json:"to_account_id"`
	CustomerID  string  `gorm:"not null" json:"customer_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) error {
	transaction.TransactionID = uuid.NewV4()
	return nil
}

func (Transaction) TableName() string {
	return "transaction"
}
