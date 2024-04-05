# Digital Circle

### Pré requisitos

- Docker
- Golang

Primeiramente, copie o arquivo `example.env` e nomeie `.env`.

Os programas dependem de um banco de dados PostgreSQL e um broker Kafka.

É disponibilizado um arquivo makefile com alguns scripts facilitadores de desenvolvimento.
Para levantar os containers via docker compose, execute o comando:
```sh
make docker-up
# ou
docker compose up -d
```

### API

A execução possibilita o uso da flag `-sql`, do qual recebe o caminho do arquivo SQL para criar a tabela TB01. 
É indicado que seja executado com a flag caso seja a primeira vez que execute o programa para que seja criada a tabela.

Para executar a API:
```sh
make tidy
make run-api
# ou
go mod tidy
go run ./cmd/api -sql ./db/create_table_TB01.sql
```

#### Endpoints

São disponibilizado os seguintes endpoints:
- `POST /tb01`: Criar um registro de dados correspondentes a tb01

O envio dos dados no corpo da requisição segue o seguinte exemplo de modelo:
```json
{
  "col_texto":"test",
}
```

O formato padrão de resposta da API segue o modelo abaixo:
```json
{
  "id": "1",
  "col_texto": "test",
  "col_dt": "2024-04-05T18:41:13.51496-03:00",
}
```

### Kafka App

A execução possibilita o uso das flags:
- `c`: nome do tópico que deseja consumir as mensagens (consome o tópico 1 por padrão)
- `k`: chave da mensagem a ser enviada
- `m`: mensagem a ser enviada

É indicado que execute o comando para criar os tópicos após a inicialização do container do Kafka:
```sh
make kafka-create-topic topico1
make kafka-create-topic topico2
```

É importante que o arquivo `.env` contenha informações consistentes para a execução correta do programa, como por exemplo, o nome referente aos tópicos.

Execute o consumidor do topico 2 primeiro:
```sh
go run ./cmd/kafka -c topico2
```

Em outro terminal, execute o comando que envia a mensagem para o tópico 1, faz a leitura do tópico 1 e encaminha para o tópico 2:
```sh
go run ./cmd/kafka -m mensagem -k chave
```

### Testes

É disponibilizado um shell script para fazer requisição nos endpoints da API.

Para a execução dos testes manuais via script, execute o script na pasta `./scripts`, ou o comando abaixo:
```sh
make test-script
```

Para a execução de testes unitários com checagem de condição de corrida, execute o comando abaixo:
```sh
make tests
```

### Obrigado! 