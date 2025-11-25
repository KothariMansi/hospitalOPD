# hospitalOPD

A lightweight OPD (Outpatient Department) management API written in Go. This repository implements the backend server, database queries (via sqlc), and utilities for configuration and migrations.

This tailored README reflects the repository layout and build/run steps for this project as found in the repository (Go 1.24, Gin, sqlc + MySQL).

## Table of contents

- [What this is](#what-this-is)
- [Tech stack & key dependencies](#tech-stack--key-dependencies)
- [Repository layout](#repository-layout)
- [Prerequisites](#prerequisites)
- [Configuration / environment variables](#configuration--environment-variables)
- [Database: schema, queries, and sqlc](#database-schema-queries-and-sqlc)
- [Build & run (local)](#build--run-local)
- [Testing](#testing)
- [Development notes](#development-notes)
- [Contributing](#contributing)
- [Contact](#contact)

## What this is

hospitalOPD is a Go-based HTTP API server for managing outpatient department workflows (patients, appointments/specialities, staff). The server is implemented with Gin and uses generated database code from sqlc to access a MySQL database.

The application entrypoint is `main.go`, which loads configuration, connects to the database, initializes the store and API server, then starts the HTTP server.

## Tech stack & key dependencies

- Language: Go (go 1.24.1 — declared in go.mod)
- HTTP framework: github.com/gin-gonic/gin
- Database: MySQL (github.com/go-sql-driver/mysql)
- DB query generator: sqlc (configured in `sqlc.yaml`)
- Config: github.com/spf13/viper
- UUIDs: github.com/google/uuid

See `go.mod` for a full dependency list.

## Repository layout (relevant files/folders)

- main.go — application entrypoint
- api/ — HTTP handlers and server initialization (server.go, user.go, hospital.go, speciality.go, check_up_time.go, client.go, ...)
- db/
  - query/ — SQL query files used by sqlc
  - migration/ — SQL migration files for schema
  - sqlc/ — auto-generated code produced by sqlc
  - util/ — configuration and helpers (config loader used in main.go)
- sqlc.yaml — sqlc configuration targeting MySQL
- go.mod / go.sum

## Prerequisites

- Go 1.24 (or compatible)
- MySQL server (for local development you can use a local mysql instance or Docker)
- sqlc (if you need to re-generate db/sqlc)
- git

Optional:
- Docker / docker-compose if you add containerization later

## Configuration / environment variables

The app loads configuration via a util loader (viper). The repository expects environment variables / configuration for the DB and server address. Typical variables used by this project include (verify exact keys in `db/util`):

Example .env (adapt values as needed)
```
# Database
DB_DRIVER=mysql
DB_SOURCE=user:password@tcp(localhost:3306)/hospitalopd?parseTime=true

# Server
SERVER_ADDRESS=0.0.0.0:8080

# Optional general secret (if used by the application)
# SECRET_KEY=replace-with-a-secret
```

Note: The code references fields like `config.DBDriver`, `config.DBSource`, and `config.ServerAddress` in `main.go`. Use those exact keys when checking the config struct in `db/util`.

## Database: schema, queries, and sqlc

- sqlc is configured in `sqlc.yaml` to use MySQL, with:
  - queries in `db/query/`
  - schema files in `db/migration/`
  - generated Go package output at `db/sqlc/`

Typical workflow:
1. Apply migrations to create the database schema. If migration SQL files are in `db/migration/`, apply them with MySQL CLI or a migration tool (e.g., goose, migrate):
   - Example (mysql client):  
     mysql -u user -p hospitalopd < db/migration/000001_init.sql
2. Generate typed DB access code if you change SQL or schema:
   - Install sqlc (https://sqlc.dev) and run:
     sqlc generate
   - This will update files under `db/sqlc/`.
3. Ensure `DB_SOURCE`/`DB_DRIVER` point to the running MySQL instance.

## Build & run (local)

From the repository root:

1. Set environment variables (or use a .env loader). Example:
   - export DB_DRIVER=mysql
   - export DB_SOURCE='user:password@tcp(localhost:3306)/hospitalopd?parseTime=true'
   - export SERVER_ADDRESS='0.0.0.0:8080'

2. (Optional) Generate sqlc code if you made SQL changes:
   - sqlc generate

3. Run the server:
   - go run main.go
   or build a binary:
   - go build -o hospitalopd .
   - ./hospitalopd

The server binds to `config.ServerAddress` (e.g., `:8080`) as loaded by the config loader.

## Testing

Run unit tests and package tests:
- go test ./...

For behavior/integration tests you may need a running MySQL instance or a test database.

## Development notes

- Config is loaded via a util package (viper). Inspect `db/util` to see exact keys and file-based config options.
- API routes, handlers, and request models are in the `api/` package.
- SQL defined in `db/query/*.sql` is compiled by sqlc — prefer editing/adding .sql files there and re-running `sqlc generate`.
- Consider adding:
  - a Makefile with common targets (build, run, test, sqlc)
  - Dockerfile / docker-compose for a reproducible dev environment
  - CI (GitHub Actions) to run `go test` and vet/lint checks

## Contributing

- Fork the repo.
- Create a topic/feature branch: `git checkout -b feat/your-feature`.
- Add tests for new functionality and run the test suite.
- Open a pull request with a clear description of the changes.

Security note: Do not commit real patient data or secrets into the repository. Use environment variables or a secrets manager for production deployments.

## Contact / Maintainer

Maintainer: KothariMansi — https://github.com/KothariMansi
