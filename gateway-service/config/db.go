package config

import (
	"fmt"
	"log"
	// "os"

	"banking_ledger/gateway-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "tanvi:root@tcp(localhost:3306)/ledger_account?charset=utf8mb4&parseTime=True&loc=Local" // MySQL DSN

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("📌 Connected to Database!")

	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&model.Account{})
	DB.AutoMigrate(&model.Transaction{})

}
