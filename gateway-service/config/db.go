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
	// Read DB credentials from environment variables
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// dbCharset := os.Getenv("DB_CHARSET")
	// // dbParseTime := os.Getenv("DB_PARSE_TIME")
	// dbLoc := os.Getenv("DB_LOC")

	// Construct DSN
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%s",
	// 	dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbLoc)
		// dsn := "tanvi:root@tcp(localhost:3306)/ledger_account?charset=utf8mb4&parseTime=True&loc=Local"
    dsn := "tanvi:root@tcp(localhost:3306)/ledger_account?charset=utf8mb4&parseTime=True&loc=Local" // MySQL DSN

	// dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable" // PostgreSQL DSN

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("📌 Connected to Database!")

	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&model.Account{})
	DB.AutoMigrate(&model.Transaction{})
	 
}

// Convert `string` boolean to actual `bool`
func parseBool(str string) bool {
	return str == "true" || str == "1"
}
