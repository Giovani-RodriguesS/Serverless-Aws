# Serverless-AWS

Repositório foi criado visando implementação prática dos estudos referentes ao serviço AWS Lambda e Arquitetura Orientada a Eventos. 

## Objetivo
O objetivo foi criar um sistema com funções Lambda como microsserviços responsáveis por:
- Converter `.json` em dados estruturados
- Gravar no DynamoDB, depois de validar os dados em fila
- Processar dados em DQL (Dead Queue Letter)

Foram usados padrões como:
- Desacoplamento por meio de Filas
- Resiliência
- Observabilidade

Antipadrões observados e evitados:
- Função Lambda monolito
- função que invoca outra função
- Invocação recursiva

## Requisitos

* [AWS CLI](https://docs.aws.amazon.com/pt_br/cli/latest/userguide/getting-started-install.html)
* [Docker](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* [SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Estrutura do Projeto

* **`src/`**: Código-fonte em Go. Cada microsserviço é uma função AWS Lambda independente.
* **`docs/`**: Documentação, diagramas de arquitetura e guias sobre DynamoDB, SQS e Lambda.
* **`events/`**: Payloads de eventos (JSON) para testes locais das funções.
* **`templates/`**: Arquivos `template.yaml` que definem a infraestrutura como código (IaC).
* **`scripts/`**: Pasta para scripts auxiliares.

```bash
├── Makefile
├── README.md
├── docs
│   ├── arquitetura-aws.drawio
│   ├── img
│   ├── overview.md
│   └── wiki
├── events
│   ├── event.json
│   └── sqs_event.json
├── samconfig.toml
├── scripts
│   └── go-fmt.sh
├── src
│   ├── error_handler
│   ├── pkg
│   ├── register
│   ├── transformer
│   └── writer
├── template.yaml
└── templates
    ├── resources
    ├── samconfig.toml
    └── template.yaml
```

