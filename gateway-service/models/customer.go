package model

import (
	"gorm.io/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Customer model 
type Customer struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   uuid.UUID  `gorm:"type:char(36);unique;not null" json:"customer_id"`
	FirstName    string     `gorm:"not null" json:"first_name"`
	LastName     string     `gorm:"not null" json:"last_name"`
	Email        string     `gorm:"unique;not null" json:"email"`
	Password     string     `gorm:"size:100;not null;" json:"password"`

	CreatedAt time.Time
	UpdatedAt time.Time
}


func (customer *Customer) BeforeCreate(tx *gorm.DB) error {
	customer.CustomerID = uuid.NewV4()
	return nil
}

// func Hash(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func (customer *Customer) BeforeSave() error {
// 	hashedPassword, err := Hash(customer.Password)
// 	if err != nil {
// 		return err
// 	}
// 	customer.Password = string(hashedPassword)
// 	return nil
// }
