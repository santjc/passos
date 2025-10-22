# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Database Migration Commands
# Requires: DATABASE_URL environment variable
# Example: export DATABASE_URL="postgres://user:password@localhost:5432/passos?sslmode=disable"

migrate-up:
	@echo "Running migrations..."
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" up

migrate-down:
	@echo "Rolling back last migration..."
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" down 1

migrate-down-all:
	@echo "Rolling back all migrations..."
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" down

migrate-force:
	@echo "Forcing migration version to $(VERSION)..."
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" force $(VERSION)

migrate-version:
	@echo "Current migration version:"
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" version

migrate-create:
	@echo "Creating migration $(NAME)..."
	@migrate create -ext sql -dir internal/database/migrations -seq $(NAME)

migrate-goto:
	@echo "Migrating to version $(VERSION)..."
	@migrate -path internal/database/migrations -database "$(DATABASE_URL)" goto $(VERSION)

# SQLc commands
sqlc-generate:
	@echo "Generating sqlc code..."
	@sqlc generate

.PHONY: all build run test clean watch docker-run docker-down itest \
	migrate-up migrate-down migrate-down-all migrate-force migrate-version \
	migrate-create migrate-goto sqlc-generate
