# Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for PostgreSQL database migrations.

## Setup

### Install golang-migrate

```bash
# macOS
brew install golang-migrate

# Or via Go
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Configure Database URL

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/passos?sslmode=disable"
```

## Commands

| Command | Description |
|---------|-------------|
| `make migrate-up` | Apply all pending migrations |
| `make migrate-down` | Rollback last migration |
| `make migrate-down-all` | Rollback all migrations |
| `make migrate-version` | Check current migration version |
| `make migrate-create NAME=feature` | Create new migration |
| `make migrate-goto VERSION=1` | Go to specific version |
| `make migrate-force VERSION=1` | Force version (error recovery) |
| `make sqlc-generate` | Generate Go code from SQL |

## Quick Start

```bash
# Start PostgreSQL
make docker-run

# Set database URL
export DATABASE_URL="postgres://user:password@localhost:5432/passos?sslmode=disable"

# Apply migrations
make migrate-up

# Check status
make migrate-version
```

## Creating Migrations

```bash
# Create new migration
make migrate-create NAME=add_users_table

# Edit generated files:
# - 000002_add_users_table.up.sql
# - 000002_add_users_table.down.sql

# Apply migration
make migrate-up
```

## Best Practices

### ✅ Do
- Always create both UP and DOWN files
- Test migrations in development first
- Use descriptive migration names
- Commit code and migrations together
- Backup before production migrations

### ❌ Don't
- Modify applied migrations (create new ones instead)
- Use DROP CASCADE without consideration
- Skip migration versions
- Run migrations manually in production

## Example Migration

**UP: `000002_add_email_index.up.sql`**
```sql
CREATE INDEX idx_contact_email ON contact(email_address) 
WHERE deleted_at IS NULL;

ALTER TABLE contact 
ADD CONSTRAINT uq_contact_email UNIQUE (email_address);
```

**DOWN: `000002_add_email_index.down.sql`**
```sql
ALTER TABLE contact DROP CONSTRAINT IF EXISTS uq_contact_email;
DROP INDEX IF EXISTS idx_contact_email;
```

## Troubleshooting

### Dirty Database Version
```bash
make migrate-version
make migrate-force VERSION=1
make migrate-up
```

### Connection Issues
```bash
# Check PostgreSQL is running
docker ps

# Test connection
psql $DATABASE_URL -c "SELECT 1"
```

## File Structure

```
internal/database/
├── migrations/
│   ├── 000001_initial_schema.up.sql
│   ├── 000001_initial_schema.down.sql
│   └── ...
├── queries.sql
└── schema.sql
```

## References

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [SQLc Documentation](https://docs.sqlc.dev/)
- [PostgreSQL Best Practices](https://www.postgresql.org/docs/current/ddl-alter.html)

