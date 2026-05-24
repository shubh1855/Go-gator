### Repo description

```text
CLI RSS aggregator built in Go as a learning project using PostgreSQL, SQLC, and Goose
```

---

### README.md

````markdown
# Go-Gator

A CLI-based RSS feed aggregator written in Go.

This project is being built as part of learning backend development in Go, focusing on database-backed CLI application design.

## Current Features

- Command-based CLI architecture
- User registration
- User login
- PostgreSQL integration
- Database migrations with Goose
- Type-safe query generation with SQLC
- Local config management

## Tech Stack

- Go
- PostgreSQL
- SQLC
- Goose
- lib/pq
- UUID

## Project Structure

```text
.
├── internal
│   ├── cli
│   │   ├── commands.go
│   │   ├── login.go
│   │   ├── register.go
│   │   └── state.go
│   ├── config
│   │   └── config.go
│   └── database
│       ├── db.go
│       ├── models.go
│       └── users.sql.go
├── sql
│   ├── queries
│   │   └── users.sql
│   └── schema
│       └── 001_users.sql
├── main.go
├── sqlc.yaml
├── go.mod
└── go.sum
```
````

## Setup

Install dependencies:

```bash
go mod tidy
```

Install tools:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Database Setup

Create PostgreSQL database:

```sql
CREATE DATABASE gator;
```

Run migrations:

```bash
cd sql/schema
goose postgres "postgres://postgres:<password>@localhost:5432/gator" up
```

## Config

Create:

```bash
~/.gatorconfig.json
```

Example:

```json
{
  "db_url": "postgres://postgres:<password>@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Usage

Register:

```bash
go run . register alice
```

Login:

```bash
go run . login alice
```

## Notes

This project is still under active development. More RSS aggregation functionality will be added incrementally.
