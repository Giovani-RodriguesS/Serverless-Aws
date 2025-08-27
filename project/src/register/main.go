package main

import (
	//"context"
	"encoding/json"
	"fmt"

	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/models"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/sqs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func parseJsonToItem(body string) (any, error) {
	var data models.Data
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, err
	}

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
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Deserialização do JSON
	item, err := parseJsonToItem(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "JSON inválido ou vazio",
		}, nil
	}

	// Serialização do JSON
	jsonBytes, err := json.Marshal(item)

	if err != nil {
		fmt.Printf("Erro ao converter produto para JSON: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Erro ao converter produto para JSON",
		}, nil
	}

	msg := string(jsonBytes)

	// Mensageria
	err = sqs.PostMessages(msg)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Erro ao tentar enviar mensagem",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       msg,
	}, nil
}

// Função principal
func main() {
	lambda.Start(handler)
}
