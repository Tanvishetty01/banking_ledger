package main

import (
	"log"
	"banking_ledger/transaction-processor/config"
	"banking_ledger/transaction-processor/handler"
)

func main() {
	log.Println("🚀 Starting Transaction Consumer Service...")

	// Initialize DynamoDB
	config.InitDynamoDB()

	// Start Kafka Consumer
	handler.ConsumeTransactions()
}
