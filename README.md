### Banking Ledger System - README

### Overview

This Banking Ledger System consists of two microservices:

Gateway Service - Handles account creation, transactions, and retrieving transaction history.

Transaction Service - Stores transaction logs in DynamoDB received via Kafka

### How It Works

The Gateway Service receives API requests.

Transactions are published to Kafka.

The Transaction Service stores logs in DynamoDB.
