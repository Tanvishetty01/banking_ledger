package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamoDB Client
var DynamoDB *dynamodb.Client

func InitDynamoDB() {
	// Load AWS SDK Config with Environment Variables
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"), // Any region (doesn't matter for local)
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(
			func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     "local",
					SecretAccessKey: "local",
					SessionToken:    "",
					Source:          "HardcodedCredentials",
				}, nil
			},
		)),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL: "http://localhost:8000",
						// URL: "http://dynamodb:8000", // Local DynamoDB URL // localhost:
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)),
	)
	if err != nil {
		log.Fatalf("❌ Unable to load AWS SDK config: %v", err)
	}

	DynamoDB = dynamodb.NewFromConfig(cfg)
	log.Println("✅ DynamoDB Initialized (Local Mode)")
}
