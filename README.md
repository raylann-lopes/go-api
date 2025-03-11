# Go API

Este é um projeto de uma API simples construída com a linguagem Go. O objetivo é fornecer um ponto de partida para a criação de APIs eficientes e escaláveis.

## Estrutura do Projeto

O projeto é organizado da seguinte forma:

- `cmd/` - Contém o ponto de entrada da aplicação.
- `controller/` - Gerencia as lógicas de controle da API.
- `db/` - Configurações relacionadas ao banco de dados.
- `model/` - Define os modelos de dados.
- `repository/` - Interage com o banco de dados.
- `usecase/` - Contém as regras de negócios da aplicação.
- `.env` - Arquivo de variáveis de ambiente.
- `Dockerfile` - Configuração para a criação de um container Docker.
- `docker-compose.yml` - Arquivo para orquestração de containers.
- `go.mod` e `go.sum` - Gerenciamento de dependências Go.

## Como Executar

### 1. Instale as dependências:

```bash
go mod tidy
```

### 2. Configure o arquivo `.env` com suas variáveis de ambiente.

### 3. Execute o projeto:

```bash
go run cmd/main.go
```

Ou, se preferir usar o Docker:

```bash
docker-compose up
```

### 3. Tecnologias Usadas:

- Go 
- Docker 
- Banco de Dados PostgreSQL


### 3. Tecnologias Usadas:

Sinta-se à vontade para contribuir com melhorias ou correções. Para isso, faça um fork deste repositório, crie uma branch com a sua modificação e envie um pull request.
