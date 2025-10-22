# PassOS

The project aims to develop an open-source ticketing and booking management system. The system will provide a modular and extensible foundation for managing attractions, tours, and events, focusing on simplicity, performance, and scalability.

The initial goal is to build a lightweight backend written in **Go (Golang)** with **PostgreSQL** as the main database. The system will expose a RESTful API to handle attractions, availabilities, bookings, and ticket validations.

## Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 15+
- golang-migrate CLI
- SQLc (optional, for code generation)

### Setup Database

```bash
# Start PostgreSQL with Docker
make docker-run

# Set database URL
export DATABASE_URL="postgres://user:password@localhost:5432/passos?sslmode=disable"

# Run migrations
make migrate-up

# Check migration status
make migrate-version
```

### Generate Code (SQLc)

```bash
make sqlc-generate
```

### Run Application

```bash
# Development mode with hot reload
make watch

# Or build and run
make build
./main
```

## Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations. See [MIGRATIONS.md](./MIGRATIONS.md) for detailed documentation.

### Common Commands

```bash
make migrate-up          # Apply all pending migrations
make migrate-down        # Rollback last migration
make migrate-version     # Check current version
make migrate-create NAME=add_feature  # Create new migration
```

## Documentation

- [MVP PassOS: Open Source OCTO Ticketing API v1](https://www.notion.so/MVP-PassOS-Open-Source-OCTO-Ticketing-API-v1-28d65c634a888059960cf39480785b17?pvs=21)
- [PassOS: Data Model](https://www.notion.so/PassOS-Data-Model-28e65c634a888031a9b9d1f7a48ea25c?pvs=21)
- [PassOS Architecture](https://www.notion.so/PassOS-Architecture-29365c634a8880929a38ed9025a1b273?pvs=21)
- [Database Migrations Guide](./MIGRATIONS.md)