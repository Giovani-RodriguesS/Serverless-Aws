## AWS SQS

O AWS SQS é um serviço de filas de mensagens da AWS que permite **desacoplar** e **escalar** aplicações.

**Conceitos:**
* **Produtores** enviam mensagens para uma fila SQS.
* O SQS **armazena** essas mensagens.
* **Consumidores** recebem, processam e, após sucesso, **excluem** as mensagens.
* Um **Visibility Timeout** impede o processamento duplicado da mesma mensagem.

**Funcionalidades Principais:**

* **Filas Standard:**
    * Alta taxa de transferência.
    * Garantia "at-least-once delivery" (pelo menos uma entrega).
    * **Não garante a ordem** das mensagens.
    * Ideal para logs e tarefas em background.

* **Filas FIFO (First-In-First-Out):**
    * **Garante a ordem exata** das mensagens.
    * Garantia "exactly-once processing" (processamento exatamente uma vez).
    * **Deduplicação** de mensagens.
    * Ideal para transações financeiras e gerenciamento de pedidos.

* **Filas de Mensagens Mortas (Dead-Letter Queues - DLQ):**
    * Armazenam mensagens que **não puderam ser processadas** com sucesso após um número configurável de tentativas.
    * Facilitam a **depuração** de falhas no processamento.


Referencias: 
- https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html