package main

import (
	"context"
	"fmt"
	"log"

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
		log.Fatal(err)
	}

	downloader := s3manager.NewDownloader(sess)
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Printf("Baixando arquivo %s do bucket %s", key, bucket)

		buff := &aws.WriteAtBuffer{}

		numBytes, err := downloader.Download(buff,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(key),
			})
		
		if err != nil {
			log.Fatalf("Falha ao baixar o arquivo %s, %v", key, err)
		}

		content := buff.Bytes()[:numBytes]

		fmt.Print("Conteudo do arquivo baixado:\n")
		fmt.Printf("%s\n", content)

		// Implementar logica de validar se é uma lista de JSONs ou um único JSON
		
		// Registrar erro em bucket de logs caso haja
		// enviar dados para fila SQS Transformer/register caso haja sucesso
		// Estudar implementação de testes e estudar GO 
	}
}
func main() {
	lambda.Start(handler)
}