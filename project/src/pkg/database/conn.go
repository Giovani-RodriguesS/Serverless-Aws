package database

import (
	"context"
	"log"
	"os"

	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ConnDB(ctx context.Context) (*dynamodb.Client, error) {

	var cfg aws.Config
	var err error
	env := os.Getenv("ENV")

	if env == "local" {
		region := os.Getenv("AWS_DEFAULT_REGION")
		URL := os.Getenv("DB_URL")
		cfg, err = utils.GetConfig(ctx, region, URL)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}

	if err != nil {
		log.Fatalf("erro ao carregar configurações no DynamoDB: %v", err)
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
