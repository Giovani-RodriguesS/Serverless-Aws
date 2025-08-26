package database

import (
	"context"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutOnTable(ctx context.Context, dynamo *dynamodb.Client, item map[string]types.AttributeValue) (error) {
	table := os.Getenv("TABLE_NAME")
	_, err := dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	})

	return err

}