package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"banking_ledger/transaction-processor/config"
	"banking_ledger/transaction-processor/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka Consumer Function
func ConsumeTransactions() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "transaction-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("❌ Failed to create Kafka consumer: %s", err)
	}

	topic := "transaction-events"
	consumer.SubscribeTopics([]string{topic}, nil)

	log.Println("🎧 Kafka Consumer started... Listening for transactions...")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var event model.Transaction
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("❌ Error parsing Kafka message: %v", err)
				continue
			}

			// Process transaction and save to DynamoDB
			log.Printf("🔹 Received Transaction: %+v", event)
			saveToDynamoDB(event)
		} else {
			log.Printf("❌ Error reading Kafka message: %v", err)
		}
	}
}

// Function to store transaction in DynamoDB
func saveToDynamoDB(event model.Transaction) {
	var ai *types.AttributeValueMemberS
	if event.ToAccountID != nil {
		ai = &types.AttributeValueMemberS{Value: *event.ToAccountID}
	} else {
		ai = &types.AttributeValueMemberS{Value: "N/A"} // ✅ Store NULL explicitly
	}

	item := map[string]types.AttributeValue{
		"transaction_id":   &types.AttributeValueMemberS{Value: string(event.TransactionID.String())},
		"amount":           &types.AttributeValueMemberN{Value: fmt.Sprintf("%.2f", event.Amount)},
		"transaction_type": &types.AttributeValueMemberS{Value: event.TransactionType},
		"notes":            &types.AttributeValueMemberS{Value: event.Notes},
		"from_account_id": &types.AttributeValueMemberS{Value: event.AccountID},
		"to_account_id": ai,
		"customer_id":   &types.AttributeValueMemberS{Value: event.CustomerID},
		"timestamp": &types.AttributeValueMemberS{Value: event.UpdatedAt.String()},
	}
	_, err := config.DynamoDB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Transactions"),
		Item:      item,
	})

	if err != nil {
		log.Printf("❌ Failed to insert into DynamoDB: %v", err)
	} else {
		log.Println("✅ Transaction saved to DynamoDB")
	}
}

// Helper function to handle nil strings
// func getString(value string) string {
// 	if value == "" {
// 		return "NULL"
// 	}
// 	return value
// }
