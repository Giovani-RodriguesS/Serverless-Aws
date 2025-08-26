package main

import (
	"context"
	"fmt"
	"time"

	"error_handler/internal/database"
	"error_handler/internal/models"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.SQSEventResponse, error) {
	
	// Conexão com DynamoDB
	dynamo, err := database.ConnDB(ctx)
	if err != nil {
		return events.SQSEventResponse{}, fmt.Errorf("erro ao obter conexão ao Banco de dados: %v", err)
	}

	var batchItemFailures []events.SQSBatchItemFailure

	// Processa cada mensagem
	for _, message := range sqsEvent.Records {
		log.Printf("processando mensagem ID: %s, Body: %s", message.MessageId, message.Body)
		fmt.Print(message.MessageId)
		
		// FUNÇÂO PARA VALIDAR CAUSA DO ERRO
		
		logItem := models.Log{
			ID:   message.MessageId,
			Message: message.Body,
			Cause:   "Unknown",
			Level:   "Error",
			Timestamp: time.Now().Format(time.RFC3339),
		}

		// Formato item DynamoDB
		av, err := attributevalue.MarshalMap(logItem)
		if err != nil {
			log.Printf("Erro ao serializar item: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

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