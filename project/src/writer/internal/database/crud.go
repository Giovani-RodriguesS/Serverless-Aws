package database

import (
	"context"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTable(ctx context.Context, dynamo *dynamodb.Client) (error) {
	_, err := dynamo.CreateTable(ctx, &dynamodb.CreateTableInput{
			TableName: aws.String(os.Getenv("TABLE_NAME")),
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
		    		AttributeType: types.ScalarAttributeTypeS,
				},
			},
			KeySchema: []types.KeySchemaElement{
				{
		    		AttributeName: aws.String("id"),
		    		KeyType:     types.KeyTypeHash,
				},
			},
			ProvisionedThroughput: &types.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		})
	
	return err
}

func PutOnTable(ctx context.Context, dynamo *dynamodb.Client, item map[string]types.AttributeValue) (error) {
	table := os.Getenv("TABLE_NAME")
	_, err := dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	})

	return err

}