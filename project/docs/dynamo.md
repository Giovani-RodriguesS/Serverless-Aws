# Documentação do AWS DynamoDB em Ambiente Serverless Local

## Conexão entre AWS Lambda e DynamoDB Local

A conexão entre AWS Lambda e DynamoDB em ambiente local foi implementada em várias etapas:

1. **Configuração do DynamoDB Local**:
   - Imagem Docker: `amazon/dynamodb-local`
   - Porta de exposição: 8000
   - Rede Docker: configurada para usar a rede `sam` (bridge)

2. **Invocação da função Lambda na mesma rede**:
   - Utilizei o parâmetro `--docker-network sam` para garantir que a função Lambda fosse executada na mesma rede Docker do DynamoDB
   ```bash
   sam local invoke WriterFunction --docker-network sam -e events/sqs_event.json
   ```

3. **Configuração de credenciais**:
   - Mesmo em ambiente local, o DynamoDB requer credenciais (ainda que fictícias)
   - Implementei a seguinte configuração em `config.LoadDefaultConfig`, em **conn.go**:
   ```go
   config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("dummy", "dummy", ""))
   ```

4. **Compartilhamento do banco de dados**:
   - **Problema**: inconsistência entre tabelas criadas via função Lambda e via AWS CLI
   - **Solução**: uso do parâmetro `-sharedDb` na inicialização do DynamoDB local. Isso garante que todas as operações usem o mesmo banco de dados independente da **região ou credenciais**

## Benefícios do DynamoDB em Contexto Serverless

1. **Escalabilidade automática**: adapta-se facilmente à carga de trabalho sem necessidade de provisionamento manual

2. **Modelo pay-per-use**: alinhado com o conceito serverless de pagar apenas pelo que é utilizado

3. **Baixa latência**: acesso rápido aos dados, essencial para funções Lambda que precisam ser eficientes

4. **Integração nativa**: fácil conexão com outros serviços AWS como Lambda e API Gateway

5. **Consistência eventual ou forte**: flexibilidade para escolher o modelo de consistência adequado à aplicação

6. **Operação sem estado**: complementa perfeitamente o paradigma serverless de funções sem estado

## Resolução de Problemas de Conexão Local

| Problema | Causa | Solução |
|----------|-------|---------|
| Lambda não encontrava o DynamoDB | Redes Docker diferentes | Usar `--docker-network sam` para colocar ambos na mesma rede |
| Erros de autenticação | Ausência de credenciais | Configurar credenciais fictícias usando `NewStaticCredentialsProvider` |
| Inconsistência entre tabelas | Isolamento padrão do DynamoDB local | Utilizar o parâmetro `-sharedDb` ao iniciar o DynamoDB local |
| Endereço de endpoint incorreto | URL do DynamoDB local mal configurada | Configurar endpoint explicitamente para `http://dynamodb-local:8000` |

Esta configuração garante um ambiente de desenvolvimento local confiável para testar a integração entre serviços serverless, simulando adequadamente o comportamento em produção na AWS.