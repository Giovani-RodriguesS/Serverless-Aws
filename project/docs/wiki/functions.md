##  Funções Lambda do Projeto

### TransformerFunction

* **Responsabilidade:**
  Recebe eventos do **S3 (upload de objetos)**, aplica transformações/normalizações nos dados e envia mensagens para a fila **WriterQueue (SQS)**.
* **Eventos:**

  * **S3Event:** Disparado em `s3:ObjectCreated:*` no bucket configurado.
* **Recursos acessados:**

  * **SQS (WriterQueue):** Envio de mensagens.
  * **CloudWatch Logs:** Registro de logs de execução.
* **Deployment:**

  * `AutoPublishAlias: live`
  * `DeploymentPreference: Canary10Percent5Minutes`
* **Alarmes integrados ao canary:**

    * `AliasErrorMetricGreaterThanZeroAlarm`
    * `LatestVersionErrorMetricGreaterThanZeroAlarm` 

---

### WriterFunction

* **Responsabilidade:**
  Consome mensagens da fila **WriterQueue (SQS)**, processa os dados e grava no **DynamoDB (AccountTable)**.
* **Eventos:**

  * **SQSEvent:** Mensagens recebidas da WriterQueue.
* **Recursos acessados:**

  * **DynamoDB (AccountTable):** Escrita de registros processados.
  * **CloudWatch Logs:** Logs de processamento.
  * **SQS (WriterQueue):** Consumo e remoção de mensagens.
* **Deployment:**

  * `AutoPublishAlias: live`
  * `DeploymentPreference: Canary10Percent5Minutes`
  * **Alarmes integrados ao canary:**

    * `AliasErrorMetricGreaterThanZeroAlarm`
    * `LatestVersionErrorMetricGreaterThanZeroAlarm`

---

### ErrorHandlerFunction

* **Responsabilidade:**
  Processa mensagens da **DeadLetterQueue (DLQ)**, registra informações de falhas no **DynamoDB (LogTable)** e gera relatórios de erro.
* **Eventos:**

  * **DQLEvent:** Mensagens recebidas da DLQ (falhas do WriterFunction).
* **Recursos acessados:**

  * **DynamoDB (LogTable):** Escrita de registros de erro.
  * **CloudWatch Logs:** Logs do tratamento de falhas.
  * **SQS (DeadLetterQueue):** Consumo e exclusão de mensagens.
* **Deployment:**

  * `AutoPublishAlias: live`
  * `DeploymentPreference: Canary10Percent5Minutes`
    
* **Alarmes integrados ao canary:**

    * `AliasErrorMetricGreaterThanZeroAlarm`
    * `LatestVersionErrorMetricGreaterThanZeroAlarm`

---

### Observações Gerais

* Todas as funções seguem o mesmo padrão de **deploy seguro com canary**:

  * **10% de tráfego** direcionado para a nova versão
  * **5 minutos de observação** antes da substituição completa
* **Alias `live`** sempre aponta para a versão estável em produção.
* **CloudWatch Alarms** são críticos para rollback automático em caso de falhas
