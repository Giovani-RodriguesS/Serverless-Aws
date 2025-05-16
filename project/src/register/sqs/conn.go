package sqs

import (
	"context"
	"log"
	"os"
	"register/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func PostMessages(message string) (error) {
	
	var cfg aws.Config
	var err error
	env := os.Getenv("ENV")

	if env == "local"{
		region := os.Getenv("AWS_DEFAULT_REGION")
		URL := os.Getenv("DB_URL")
		cfg, err =  utils.GetConfig(context.TODO(), region, URL)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO())
	}

	if err != nil {
		log.Fatal("erro ao carregar configuração do SQS")
	} 

	client := sqs.NewFromConfig(cfg)
	queueURL := os.Getenv("QUEUE_URL")

	if queueURL == "" {
		log.Fatal("QUEUE_URL não está definido")
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    &queueURL,
		MessageBody: &message,
	}

	_, err = client.SendMessage(context.TODO(), input)

	return err
}
