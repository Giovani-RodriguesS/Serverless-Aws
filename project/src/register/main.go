package main

import (
	//"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"register/sqs"
)

type Product struct {
	Nome string `json:"nome"`
	Cargo string `json:"cargo"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var product Product
	// Deserialização do JSON
	err := json.Unmarshal([]byte(request.Body), &product)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "JSON inválido",
		}, nil
	}

	// Serialização do JSON
	json, err := json.Marshal(product)

	if err != nil {
		log.Fatalf("Erro ao converter produto para JSON: %v", err)
	}

	msg := string(json)

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
