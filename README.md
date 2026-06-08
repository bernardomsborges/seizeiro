# Seizeiro

Sistema de análise e triagem automática de processos SEI usando IA.

## Requerimentos

1. [Go 1.26](https://go.dev)
2. [PostgreSQL](https://www.postgresql.com)
3. [Docker](https://www.docker.com)
4. [goose](https://pressly.github.io/goose/)
5. [sqlc](https://sqlc.dev/)

## Banco de Dados

### Schema

A gestão do schema do banco de dados é feita através de migrações
usando a ferramente [goose](https://pressly.github.io/goose/). A ferramenta
está disponível no projeto como uma [go tool](https://pkg.go.dev/cmd/go#hdr-Run_specified_go_tool).

Para criar uma nova migração, execute o comando abaixo:

```bash
go tool goose -dir internal/postgres/migrations create <nome> sql
```

Para aplicar as migrações ao banco de dados, execute o comando abaixo:

```bash
go tool goose -dir internal/postgres/migrations postgres <database_url> up
```

### Queries

As queries do banco de dados são compiladas de arquivos `.sql` pela ferramenta
[sqlc](https://sqlc.dev/). Para compilar novas queries, execute o comando abaixo:

```bash
go tool sqlc generate # ou make sql/build
```

Consulte a [documentação](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html) do
sqlc se precisar adicionar novas queries.
