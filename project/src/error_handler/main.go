package main

import (
	"context"
	"encoding/json"
	"fmt"
	logItem "github.com/Giovani-RodriguesS/Serverless-Aws/project/src/error_handler/internal/models"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/database"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"strings"
	"time"
)

var AccountTypes = []string{"SAVINGS", "CHECKING", "CURRENT"}
var TransactionTypes = []string{"DEBIT", "CREDIT"}

func detectErrorCause(message string) string {
	var data models.Data
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		fmt.Println(err)
		return "InvalidJson"
	}

	if data.Type == "account" {
		var acc models.Account
		if err := json.Unmarshal(data.Data, &acc); err == nil {
			fmt.Println(err)
			return "InvalidJson"
		}

		// Valida campos obrigatórios
		if acc.AccountID == "" || acc.Name == "" || acc.Type == "" {
			return "InvalidAccountEvent:MissingFields"
		}

		var absent bool = true
		for _, a := range AccountTypes {
			if strings.ToUpper(acc.Type) == a {
				absent = false
			}
		}
		if absent {
			return "InvalidAccountEvent:InvalidType"
		}
	}

	if data.Type == "transaction" {
		var tsc models.Transaction

		if err := json.Unmarshal(data.Data, &tsc); err == nil {
			fmt.Println(err)
			return "InvalidJson"
		}

		// Valida campos obrigatórios
		if tsc.TransactionID == "" || tsc.AccountID == "" || tsc.Timestamp == "" || tsc.Type == "" {
			return "InvalidTransactionEvent:MissingFields"
		}
		if tsc.Amount <= 0 {
			return "InvalidTransactionEvent:InvalidAmount"
		}

		var absent bool = true
		for _, t := range TransactionTypes {
			if strings.ToUpper(tsc.Type) == t {
				absent = false
			}
			if absent {
				return "InvalidTransactionEvent:InvalidType"
			}
		}
	}
	return "Unknown"
}

func generateLog(message events.SQSMessage) logItem.Log {

	logItem := logItem.Log{
		ID:        message.MessageId,
		Message:   message.Body,
		Cause:     detectErrorCause(message.Body),
		Level:     "Error",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	return logItem
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.SQSEventResponse, error) {

	// Conexão com DynamoDB
	dynamo, err := database.ConnDB(ctx)
	if err != nil {
		return events.SQSEventResponse{}, fmt.Errorf("erro ao obter conexão ao Banco de dados: %v", err)
	}

	var batchItemFailures []events.SQSBatchItemFailure

	// Processa cada mensagem
	for _, message := range sqsEvent.Records {
		fmt.Printf("processando mensagem ID: %s, Body: %s", message.MessageId, message.Body)
		fmt.Print(message.MessageId)

		logItem := generateLog(message)

		// Formato item DynamoDB
		av, err := attributevalue.MarshalMap(logItem)
		if err != nil {
			fmt.Printf("Erro ao serializar item: %v", err)
			batchItemFailures = append(batchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: message.MessageId})
			continue
		}

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
