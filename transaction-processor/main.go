package main

import (
	"banking_ledger/transaction-processor/config"
	"banking_ledger/transaction-processor/handler"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("🚀 Starting Transaction Consumer Service...")

	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Println("⚠️ Warning: No .env file found, using system environment variables")
	}
	// Initialize DynamoDB
	config.InitDynamoDB()

	// Start Kafka Consumer
	handler.ConsumeTransactions()
}
