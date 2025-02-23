package model

import (
	"time"
)

type Account struct {
	ID         uint      `json:"id"`
	AccountID  string    `json:"account_id"`
	NickName   string    `json:"nick_name"`
	Amount     float64   `json:"amount"`
	CustomerID string    `json:"customer_id"`
	Customer   Customer  `json:"customer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
