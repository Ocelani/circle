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

A execução possibilita o uso da flag `-sql`, do qual recebe o caminho do arquivo SQL para criar a tabela TB01. É indicado que seja executado com a flag caso seja a primeira vez que execute o programa para que seja criada a tabela.

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
  "user": {
    "id": "1",
    "col_texto": "test",
    "col_dt": "",
  },
}
```

### Testes

Para a execução de testes manuais, execute o script na pasta `./scripts`, ou o comando abaixo:
```sh
make tests-script
```

Para a execução de testes unitários, execute o comando abaixo:
```sh
make tests
```

Em relação ao volume de testes unitários, gostaria de ter agregado mais, porém, não foi possível devido ao tempo disponível.

### Problemas conhecidos

A execução da imagem do container correspondente à aplicação não ocorre devidamente.
Dessa forma, é indicado que a execução do programa ocorra manualmente, como indicado previamente:
```sh
docker compose up -d
go mod tidy
go run ./cmd/userapi
```