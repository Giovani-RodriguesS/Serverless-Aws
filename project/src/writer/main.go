package main

import (
	"context"
	"fmt"
	"log"
	"writer/internal/database"
	"writer/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	_ "github.com/lib/pq"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.SQSEventResponse, error) {

	// Cenexão com DynamoDB
	dynamo, err := database.ConnDB(ctx)
	if err != nil {
		return events.SQSEventResponse{}, fmt.Errorf("erro ao obter conexão ao Banco de dados: %v", err)
	}

	var batchItemFailures []events.SQSBatchItemFailure

	for _, message := range sqsEvent.Records {
		log.Printf("processando mensagem ID: %s, Body: %s", message.MessageId, message.Body)
		fmt.Print(message.MessageId)
		item := models.Item{
			ID:   message.MessageId,
			Name: message.Body,
		}

		// Formato item DynamoDB
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			log.Printf("Erro ao serializar item: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

		// CRIAR TABELA NO DYNAMODB
		// err = database.CreateTable(ctx, dynamo)
		// if err != nil {
		//  	log.Printf("Erro ao criar banco %v", err)
		// }

		// Inserir o item no DynamoDB
		err = database.PutOnTable(ctx, dynamo, av)

		if err != nil {
			log.Printf("Erro ao inserir mensagem no banco de dados: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}
	}

	return events.SQSEventResponse{BatchItemFailures: batchItemFailures}, nil
}

func main() {
	lambda.Start(handler)
}
