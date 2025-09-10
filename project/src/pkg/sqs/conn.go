package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	pkg "github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func PostMessages(message []pkg.Data) error {

	var cfg aws.Config
	var err error

	cfg, err = config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Printf("erro ao carregar configuração do SQS")
	}

	client := sqs.NewFromConfig(cfg)
	queueURL := os.Getenv("QUEUE_URL")

	if queueURL == "" {
		fmt.Printf("QUEUE_URL não está definido")
	}

	var item string
	for _, i := range message {
		bytes, marshalErr := json.Marshal(i)
		if marshalErr != nil {
			return marshalErr
		}
		item = string(bytes)
		input := &sqs.SendMessageInput{
			QueueUrl:    &queueURL,
			MessageBody: &item,
		}
		_, err = client.SendMessage(context.TODO(), input)
	}

	return err
}
