package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ConnDB(ctx context.Context) (*dynamodb.Client, error) {
	
	cfg, err := config.LoadDefaultConfig(ctx)
	
	if err != nil {
		log.Fatalf("erro ao carregar configurações no DynamoDB: %v", err)
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}