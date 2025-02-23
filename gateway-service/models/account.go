package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID uuid.UUID `gorm:"type:char(36);unique;not null" json:"account_id"`
	NickName  string    `json:"nick_name"`
	Amount    float64   `gorm:"default:0.0;not null" json:"amount"`

	CustomerID uuid.UUID `gorm:"not null" json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate hook to generate UUID before inserting into DB
func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.AccountID = uuid.NewV4()
	return nil
}

func (Account) TableName() string {
	return "account"
}
