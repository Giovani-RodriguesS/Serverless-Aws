# S3 Local - LocalStack

Para simular o `S3` localmente, usamos `localstack` no Docker. Primeiramente, como na documentação oficial mostra, criamos um `docker-compose.yaml`. Nele, definimos as configurações do `localstack`, como serviços que seriam usados nele, volumes e portas.

- [Documentação](https://docs.localstack.cloud/getting-started/installation/)


Instalação:

```bash
services:
  localstack:
    image: localstack/localstack:latest
    container_name: localstack
    ports:
      - "4566:4566"
    environment:
      - SERVICES=s3, lambda
      - AWS_DEFAULT_REGION=us-east-1
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock"
      - ".:/etc/localstack/resources"
```
```bash
docker compose up -d
```

No localstack, configuramos o bucket S3 e a função que executará:
```bash
# 1. Cria o bucket S3
awslocal s3 mb s3://bucket-da-funcao

# zip -r function.zip arquivo.go

# 2. Cria a função Lambda
awslocal lambda create-function \
    --function-name funcao-golang \
    --runtime provided.al2023 \
    --role arn:aws:iam::000000000000:role/irrelevant \
    --handler bootstrap \
    --zip-file fileb://function.zip \
    --endpoint-url=http://localhost:4566

# 3. Adiciona permissão para o S3 invocar a função Lambda
awslocal lambda add-permission \
    --function-name funcao-golang \
    --statement-id s3-invoke \
    --action "lambda:InvokeFunction" \
    --principal s3.amazonaws.com \
    --source-arn arn:aws:s3:::meu-bucket-teste \
    --endpoint-url=http://localhost:4566

```
Configura a notificação no bucket S3 para acionar a Lambda

```json
{
  "LambdaFunctionConfigurations": [
    {
      "LambdaFunctionArn": "arn:aws:lambda:us-east-1:000000000000:function:minha-funcao-s3-trigger",
      "Events": ["s3:ObjectCreated:*"]
    }
  ]
}
```
Aplicamos essa notificação ao bucket criado:
```bash
awslocal s3api put-bucket-notification-configuration \
    --bucket meu-bucket-teste \
    --notification-configuration file://notification.json \
    --endpoint-url=http://localhost:4566
```

Agora, ao subir arquivos no bucket, ele deverá acionar a função lambda conectada a ele.

---

