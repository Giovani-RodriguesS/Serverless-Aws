##azip -r function.zip src/transformer
#
#awslocal lambda create-function \
#    --function-name minha-funcao-s3-trigger \
#    --runtime provided.al2023 \
#    --role arn:aws:iam::000000000000:role/irrelevant \
#    --handler bootstrap \
#    --zip-file fileb://function.zip \
#    --endpoint-url=http://localhost:4566
#
## 3. Adicione permissão para o S3 invocar a função Lambda
#awslocal lambda add-permission \
#    --function-name minha-funcao-s3-trigger \
#    --statement-id s3-invoke \
#    --action "lambda:InvokeFunction" \
#    --principal s3.amazonaws.com \
#    --source-arn arn:aws:s3:::meu-bucket-teste \
#    --endpoint-url=http://localhost:4566

awslocal s3api put-bucket-notification-configuration \
    --bucket meu-bucket-teste \
    --notification-configuration file://notification.json \
    --endpoint-url=http://localhost:4566