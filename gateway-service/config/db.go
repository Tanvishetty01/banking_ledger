package config

import (
	"fmt"
	"log"
	"os"

	model "banking_ledger/gateway-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Ensures table names remain plural (e.g., "accounts")
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("📌 Connected to Database!")

	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&model.Account{})
	DB.AutoMigrate(&model.Transaction{})

}
