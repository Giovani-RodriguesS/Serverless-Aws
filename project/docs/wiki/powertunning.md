# Lambda Power Tuning

## Visão Geral

>A`WS Lambda Power Tuning` é uma ferramenta open-source que ajuda a visualizar e aprimorar a configuração do processamento/memória das funções Lambda

>Essa ferramenta consiste numa `máquina de estado do Step Functions` que executa teste de carga na função escolhida e retorna uma saída que contém as informações obtidas através dos testes: uma avalição de qual melhor e pior configuração para a função testada.


## Como Funciona
- Você fornece o ARN (*Amazon Resource Name*) de uma função Lambda.

- A ferramenta invoca essa função com diferentes configurações de memória (de 128 MB a 10 GB), você define.

- Ela analisa os logs de execução e sugere a melhor configuração de memória para `minimizar custos` ou `maximizar a performance`.

- A execução ocorre na sua própria conta AWS, simulando cenários reais como chamadas de API e cold starts.

- Os resultados são apresentados em um gráfico, mostrando o `custo` e a `velocidade médios` para cada configuração.


### Parâmetros de Execução (no Momento da Execução)

Esses parâmetros são usados para controlar o teste da função Lambda.

* `lambdaARN` (obrigatório): ARN da função a ser testada.
* `powerValues` (opcional): Lista de valores de memória (em MB) para testar.
* `num` (obrigatório): Número de vezes que cada configuração de memória será invocada.
* `payload` (opcional): Dados de entrada para as invocações.
* `parallelInvocation` (opcional): Se `true`, as invocações ocorrem em paralelo.
* `strategy` (opcional): Estratégia de otimização (`cost`, `speed` ou `balanced`).
* `autoOptimize` (opcional): Se `true`, aplica a configuração ideal após o teste.
* `dryRun` (opcional): Se `true`, executa a função apenas uma vez para teste.
* `discardTopBottom` (opcional): Porcentagem de resultados mais rápidos/lentos a serem descartados.

---

### Parâmetros de Configuração (no Momento da Implantação)

Esses parâmetros são definidos ao implantar a _state machine_ e afetam o comportamento padrão da ferramenta.

* `PowerValues`: Valores de memória padrão para o teste.
* `totalExecutionTimeout`: Tempo máximo de execução da _state machine_ (padrão de 300 segundos).
* `logGroupRetentionInDays`: Quantos dias os logs de eventos serão retidos.
* `securityGroupIds` / `subnetIds`: Configurações de rede para as funções Lambda.
* `stateMachineNamePrefix`: Permite customizar o nome da _state machine_.


Para mais informações sobre a ferramenta, consulte a documentação abaixo
- [Aws Lambda Power Tunning - Repo Github](https://github.com/alexcasalboni/aws-lambda-power-tuning?tab=readme-ov-file)

- [Aws Lambda Power Tunning - Página Web AWS](https://serverlessrepo.aws.amazon.com/applications/arn:aws:serverlessrepo:us-east-1:451282441545:applications~aws-lambda-power-tuning)

