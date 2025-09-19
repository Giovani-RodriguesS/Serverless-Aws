package main

import (
	"context"
	"fmt"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/database"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/writer/internal"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.SQSEventResponse, error) {

	// Cenexão com DynamoDB
	dynamo, err := database.ConnDB(ctx)
	if err != nil {
		return events.SQSEventResponse{}, fmt.Errorf("erro ao obter conexão ao Banco de dados: %v", err)
	}

	var batchItemFailures []events.SQSBatchItemFailure

	for _, message := range sqsEvent.Records {
		fmt.Printf("processando mensagem ID: %s", message.MessageId)

		// Deserialização do JSON
		data, err := internal.ParseJsonToItem(message.Body)

		if err != nil {
			fmt.Printf("Erro ao deserializar dados: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

		item, err := internal.WrapUpItem(data)

		if err != nil {
			fmt.Printf("Erro ao empacotar item: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}
		// Formato item DynamoDB
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			fmt.Printf("Erro ao serializar item: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

		// Inserir o item no DynamoDB
		err = database.PutOnTable(ctx, dynamo, av)

		if err != nil {
			fmt.Printf("Erro ao inserir mensagem no banco de dados: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}
	}

	return events.SQSEventResponse{BatchItemFailures: batchItemFailures}, nil
}

func main() {
	lambda.Start(handler)
}
