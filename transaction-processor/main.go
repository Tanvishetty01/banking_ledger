package main

import (
	"banking_ledger/transaction-processor/config"
	"banking_ledger/transaction-processor/handler"
	"log"
)

func main() {
	log.Println("🚀 Starting Transaction Consumer Service...")

	// Initialize DynamoDB
	config.InitDynamoDB()

	// Start Kafka Consumer
	handler.ConsumeTransactions()
}
