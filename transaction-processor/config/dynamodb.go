package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamoDB Client
var DynamoDB *dynamodb.Client

func InitDynamoDB() {
	// Read from environment variables
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsDynamoDBEndpoint := os.Getenv("AWS_DYNAMODB_ENDPOINT")

	// Set default values if environment variables are not set
	if awsRegion == "" {
		awsRegion = "us-east-1" // Default region
	}
	if awsDynamoDBEndpoint == "" {
		awsDynamoDBEndpoint = "http://localhost:8000" // Default local DynamoDB
	}

	// Load AWS SDK Config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(
			func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     awsAccessKey,
					SecretAccessKey: awsSecretKey,
					SessionToken:    "",
					Source:          "EnvironmentVariables",
				}, nil
			},
		)),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{URL: awsDynamoDBEndpoint}, nil
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
