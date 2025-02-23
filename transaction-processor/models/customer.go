package model

import (
	"time"
)

type Customer struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   string  `gorm:"type:char(36);unique;not null" json:"customer_id"`
	FirstName    string     `gorm:"not null" json:"first_name"`
	LastName     string     `gorm:"not null" json:"last_name"`
	Email        string     `gorm:"unique;not null" json:"email"`
	Password     string     `gorm:"size:100;not null;" json:"password"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
