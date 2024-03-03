# Go Bank Api

Este repositório é uma API de um banco (fictício) que conta com contas de usuário, saldo fiduciário, crypto, transferências, cartões, depósitos e futuras ideias.

## Aprendizado


A principal finalidade é aprender e praticar Go, utilizando (e aprendendo) boas práticas e seguindo o padrão RESTful.

## Tecnologias


| Linguagem                                                                                                    | Banco de Dados                                                                                                                                 | Cache                                                                                                                           | Container                                           |
|--------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| [![go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/) | [![postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/) | [![redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/) | [![docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/) |


## Características

- Contas de usuário
- Saldo fiduciário
- Crypto (BTC) (apenas estético, nada envolvendo blockchain)
- Transferências entre contas
- Depósitos 
- Cartões de crédito com limites definidos
- Futuras ideias

## Bibliotecas


- Esta API usa [Gin](https://gin-gonic.com/) como web framework
- [JWT](https://github.com/golang-jwt/jwt) para autenticação
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) para hash de senhas e dados sensívels (como cvv) no banco de dados
- [golang-migrate](https://github.com/golang-migrate/migrate) para controlar migrações no banco de dados

## Notas

Mais uma vez, a ideia do projeto é apredizado, não pretendo nem de perto simular a complexidade de um banco real ou envolver de fato crypto e blockchains

## Rodando a api


- Clone o repositório
```
  git clone https://github.com/gaia-j/go-bank-api 
```

- Entre no diretório
```
  cd go-bank-api
```
<br>

- Configure os arquivos `.env`* e `configs/compose-db.yaml`** com credenciais do banco de dados, jwt e hashId <br>


- Rode o container*** do banco de dados
```
  docker compose -f configs/compose-db.yaml up
```
atenção: Caso não use ou não deseja usar docker, precisará de um banco de dados postgres rodando localmente ou em nuvem.

<br>

- Baixe e instale uma versão do [golang-migrate]( ###) compatível com seu sistema operacional

- Execute a migração do banco de dados com o golang-migrate, substituindo suas credenciais postgres

```
  migrate -database postgres://user:password@localhost:5432/db?sslmode=disable -path scripts/db/migrations up
```
<br>

- Finalmente, execute a aplicação:
```
  go run cmd/goApi/main.go
```

<br>



#### * utilize o arquivo `.env-mock` como referência
#### ** note que esse container usa armazenamento persistente
#### *** necessário instalar o [docker](https://www.docker.com/products/docker-desktop/) ou [docker engine](https://docs.docker.com/engine/install/) (apenas linux)


## Build e Docs

🚧🚧🚧🚧🚧🚧🚧🚧🚧

## Estrutura principal do projeto

Inicialmente, essa é a estrutura base que usarei de referência:
```bash
project-root/ 
├── cmd/
│   ├── your-app-name/
│   │   ├── main.go         # Application entry point
│   │   └── ...             # Other application-specific files
│   └── another-app/
│       ├── main.go         # Another application entry point
│       └── ...
├── internal/                # Private application and package code
│   ├── config/
│   │   ├── config.go       # Configuration logic
│   │   └── ...
│   ├── database/
│   │   ├── database.go     # Database setup and access
│   │   └── ...
│   └── ...
├── pkg/                     # Public, reusable packages
│   ├── mypackage/
│   │   ├── mypackage.go    # Public package code
│   │   └── ...
│   └── ...
├── api/                     # API-related code (e.g., REST or gRPC)
│   ├── handler/
│   │   ├── handler.go      # HTTP request handlers
│   │   └── ...
│   ├── middleware/
│   │   ├── middleware.go  # Middleware for HTTP requests
│   │   └── ...
│   └── ...
├── scripts/                 # Build, deployment, and maintenance scripts
│   ├── build.sh
│   ├── deploy.sh
│   └── ...
├── configs/                 # Configuration files for different environments
│   ├── development.yaml
│   ├── production.yaml
│   └── ...
├── tests/                   # Unit and integration tests
│   ├── unit/
│   │   ├── ...
│   └── integration/
│       ├── ...
├── docs/                    # Project documentation
├── .gitignore               # Gitignore file
├── go.mod                   # Go module file
├── go.sum                   # Go module dependencies file
└── README.md                # Project README
```

## Authors
- [@gaia-j](https://www.github.com/gaia-j)
