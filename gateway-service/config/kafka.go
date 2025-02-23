package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"banking_ledger/gateway-service/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka configuration details
const (
	ConfirmationTopic = "transaction-confirmations"
	GroupID           = "transaction-processor"
)

var TransactionTopic = "transaction-events"

// KafkaProducer interface for producing Kafka messages
type KafkaProducer interface {
	ProduceTransaction(transaction model.Transaction) error
}

// Producer struct for Kafka
type Producer struct {
	producer *kafka.Producer
}

// NewKafkaProducer creates a new Kafka producer instance
func NewKafkaProducer() (KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOT_STRAP_SERVER"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return &Producer{producer: p}, nil
}

// ProduceTransaction sends a transaction message to Kafka
func (p *Producer) ProduceTransaction(transaction model.Transaction) error {
	message, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to serialize transaction: %w", err)
	}

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &TransactionTopic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	p.producer.Flush(5000)
	log.Println("✅ Transaction message sent successfully!")
	return nil
}
