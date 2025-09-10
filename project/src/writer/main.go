package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/database"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func parseJsonToItem(body string) (models.Data, error) {
	// Converte o dado bruto para ser possivel identificar o objeto
	var data models.Data
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return models.Data{}, err
	}
	return data, nil
}

func wrapUpItem(data models.Data) (any, error) {
	// Identifica qual item será gravado no Dynamo
	switch data.Type {
	case "account":
		var acc models.Account
		if err := json.Unmarshal(data.Data, &acc); err != nil {
			return nil, err
		}

		return &acc, nil

	case "transaction":
		var tsc models.Transaction
		if err := json.Unmarshal(data.Data, &tsc); err != nil {
			return nil, err
		}
		return &tsc, nil

	default:
		return nil, fmt.Errorf("tipo não reconhecido: %s", data.Type)
	}
}

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
		data, err := parseJsonToItem(message.Body)

		if err != nil {
			fmt.Printf("Erro ao deserializar dados: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

		item, err := wrapUpItem(data)

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
