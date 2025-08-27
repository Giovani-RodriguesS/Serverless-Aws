package main

import (
	"context"
	"fmt"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/sqs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Printf("Erro ao criar sessão: %v\n", err)
		return
	}

	downloader := s3manager.NewDownloader(sess)

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		fmt.Printf("Baixando arquivo %s do bucket %s", key, bucket)

		buff := &aws.WriteAtBuffer{}

		numBytes, err := downloader.Download(buff,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(key),
			})

		if err != nil {
			fmt.Printf("Falha ao baixar o arquivo %s, %v", key, err)
		}

		payload := buff.Bytes()[:numBytes]

		// Mensageria
		err = sqs.PostMessages(string(payload))

		if err != nil {
			fmt.Printf("Erro ao enviar mensagem para SQS: %v", err)
		}
	}
}
func main() {
	lambda.Start(handler)
}
