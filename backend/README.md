# Backend

Go REST API for the link-sharing app. Built with Chi router using a layered Handler → Service → Repository architecture.

## Stack

- **Go** (Chi router, sqlx, JWT)
- **MySQL 8** via Docker
- **dbmate** for migrations
- **Air** for hot reload

## Prerequisites

- Go 1.21+
- [Task](https://taskfile.dev) (`brew install go-task`)
- [Air](https://github.com/air-verse/air) (`go install github.com/air-verse/air@latest`)
- [dbmate](https://github.com/amacneil/dbmate) (`brew install dbmate`)
- Docker (for MySQL)

## Setup

1. Start the database:
   ```bash
   docker compose up -d
   ```

2. Copy and configure environment variables:
   ```bash
   cp .env.example .env
   ```

3. Run migrations:
   ```bash
   task migrate_up
   ```

4. Start the server with hot reload:
   ```bash
   task run_server
   ```

The API will be available at `http://localhost:8080`.

## Environment Variables

See [`.env.example`](.env.example).

## API Reference

See [`docs/API.md`](docs/API.md).

## Database Migrations

```bash
task migrate_new   # Create a new migration file
task migrate_up    # Apply all pending migrations
task migrate_down  # Roll back the last migration
task reset_db      # Drop, recreate, and migrate the database
```

## Testing

```bash
go test ./...
```

Tests use `testify` for assertions and `go-sqlmock` for database mocking.

## Project Structure

```
backend/
├── cmd/app/            # Entry point and server setup
├── db/
│   ├── db.go           # Database connection
│   └── migrations/     # SQL migration files
├── internal/
│   ├── config/         # Environment variable loading
│   ├── handlers/       # HTTP handlers and router
│   ├── services/       # Business logic
│   ├── repositories/   # Data access layer
│   ├── models/         # Domain types
│   └── pkg/
│       ├── jwt/        # Token generation and middleware
│       └── password/   # Password hashing
└── uploads/            # Uploaded avatar files
```
