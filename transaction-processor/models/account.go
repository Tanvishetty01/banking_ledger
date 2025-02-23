package model

import (

	"time"
)

type Account struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   string `gorm:"type:char(36);unique;not null" json:"account_id"`
	NickName    string    `json:"nick_name"`
	Amount      float64   `gorm:"default:0.0;not null" json:"amount"`

	CustomerID string     `gorm:"not null" json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
