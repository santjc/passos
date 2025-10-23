# PassOS Architecture Documentation

## Overview

PassOS is an open-source ticketing and booking management system designed to be OCTO-compliant. The architecture follows a pragmatic approach combining Repository Pattern, Dependency Injection, and SQLc code generation to achieve type-safety, maintainability, and performance.

## Core Principles

1. **Single Source of Truth**: Database schema serves as the canonical data model
2. **Code Generation**: SQLc generates type-safe Go code from SQL queries
3. **Minimal Duplication**: Avoid creating redundant DTOs when SQLc can generate them
4. **Type Safety**: Maintain type safety from database to HTTP response
5. **Testability**: Interfaces enable dependency injection and mocking
6. **Separation of Concerns**: Clear boundaries between layers

## Architecture Layers

```
┌─────────────────────────────────────────────────────────┐
│                     HTTP Layer                          │
│  (Handlers, Routes, Middleware)                         │
└─────────────────────┬───────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────┐
│                  Service Layer                          │
│  (Business Logic, Orchestration)                        │
└─────────────────────┬───────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────┐
│                Repository Layer                         │
│  (Interfaces + SQLc Generated Implementations)          │
└─────────────────────┬───────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────┐
│                   Database Layer                        │
│  (PostgreSQL, Migrations)                               │
└─────────────────────────────────────────────────────────┘
```

## Project Structure

```
passos/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── database/
│   │   ├── migrations/                # Database migrations
│   │   │   ├── 000001_initial_schema.up.sql
│   │   │   └── 000001_initial_schema.down.sql
│   │   ├── queries/                   # SQL queries by domain
│   │   │   ├── product.sql
│   │   │   ├── booking.sql
│   │   │   ├── availability.sql
│   │   │   └── offer.sql
│   │   └── connection.go              # Database connection setup
│   ├── repository/
│   │   ├── models.go                  # SQLc generated models
│   │   ├── db.go                      # SQLc generated DB interface
│   │   ├── queries.sql.go             # SQLc generated query methods
│   │   └── interfaces/                # Repository interfaces for DI
│   │       ├── product.go
│   │       ├── booking.go
│   │       └── availability.go
│   ├── service/                       # Business logic layer
│   │   ├── product_service.go
│   │   ├── booking_service.go
│   │   └── availability_service.go
│   ├── http/                          # HTTP transport layer
│   │   ├── handler/
│   │   │   ├── product_handler.go
│   │   │   ├── booking_handler.go
│   │   │   └── availability_handler.go
│   │   ├── dto/                       # Request/Response DTOs (when needed)
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── logger.go
│   │   │   └── validator.go
│   │   └── routes.go                  # Route definitions
│   ├── container/
│   │   └── container.go               # Dependency injection container
│   └── pkg/                           # Shared utilities
│       ├── errors/
│       ├── validator/
│       └── logger/
├── sqlc.yaml                          # SQLc configuration
└── go.mod
```

## Layer Responsibilities

### 1. Database Layer

**Purpose**: Schema definition and data persistence

**Components**:
- Migrations: Version-controlled schema changes using golang-migrate
- Queries: SQL query definitions organized by domain

**Key Files**:
- `internal/database/migrations/*.sql`: Migration files
- `internal/database/queries/*.sql`: SQL query definitions

### 2. Repository Layer

**Purpose**: Data access abstraction with type-safe queries

**Components**:
- SQLc Generated Code: Type-safe Go structs and query methods
- Repository Interfaces: Contracts for dependency injection

**Key Concepts**:
- SQLc generates models directly from database schema
- Models include JSON tags with camelCase naming
- Interfaces wrap SQLc queries for testability and DI
- No manual entity duplication

**Example Interface**:
```go
// internal/repository/interfaces/product.go
type ProductRepository interface {
    GetByID(ctx context.Context, id uuid.UUID) (repository.Product, error)
    List(ctx context.Context, params repository.ListProductsParams) ([]repository.Product, error)
    Create(ctx context.Context, params repository.CreateProductParams) error
    Update(ctx context.Context, params repository.UpdateProductParams) error
    Delete(ctx context.Context, id uuid.UUID) error
}
```

**Example Implementation**:
```go
type productRepository struct {
    queries *repository.Queries
}

func NewProductRepository(queries *repository.Queries) ProductRepository {
    return &productRepository{queries: queries}
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (repository.Product, error) {
    return r.queries.GetProductByID(ctx, id)
}
```

### 3. Service Layer

**Purpose**: Business logic and orchestration

**Responsibilities**:
- Business rule validation
- Transaction management
- Orchestration of multiple repositories
- Domain-specific logic

**Key Points**:
- Services receive repository interfaces via dependency injection
- Use SQLc-generated models directly (no entity duplication)
- Return SQLc models or custom DTOs when transformation is needed

**Example**:
```go
// internal/service/product_service.go
type ProductService struct {
    productRepo interfaces.ProductRepository
    optionRepo  interfaces.OptionRepository
}

func NewProductService(
    productRepo interfaces.ProductRepository,
    optionRepo interfaces.OptionRepository,
) *ProductService {
    return &ProductService{
        productRepo: productRepo,
        optionRepo:  optionRepo,
    }
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (*repository.Product, error) {
    product, err := s.productRepo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    return &product, nil
}
```

### 4. HTTP Layer

**Purpose**: HTTP transport and request/response handling

**Components**:
- Handlers: HTTP request handlers
- DTOs: Request/Response data transfer objects (when needed)
- Middleware: Cross-cutting concerns (auth, logging, validation)
- Routes: HTTP route definitions using Chi router

**When to Use DTOs**:

**Critical Design Decision**: While SQLc models can technically be serialized directly to JSON, this approach couples the public API contract to the internal database schema, making it difficult to evolve the database without breaking the API.

**Recommended Approach**:

Always use DTOs as a contract layer between the API and internal models:

**Response DTOs** (Always):
- Provide stable API contracts independent of DB schema
- Allow DB schema evolution without breaking changes
- Hide internal implementation details (`pgtype.Text`, nullable wrappers)
- Enable API versioning (v1, v2) with different DTOs pointing to same DB
- Support OCTO-specific formats without modifying database

**Request DTOs** (Always):
- Validate and sanitize user input
- Provide clear API documentation
- Prevent injection of unexpected fields
- Allow business-level validation separate from DB constraints

**Direct SQLc Model Usage** (Internal Only):
- Between Repository and Service layers
- Internal service-to-service communication
- Background jobs and internal processes
- Testing and development tools

**Example Handler with DTOs**:
```go
// internal/http/dto/product.go
type ProductResponse struct {
    ID                  string `json:"id"`
    InternalName        string `json:"internalName"`
    Locale              string `json:"locale"`
    TimeZone            string `json:"timeZone"`
    AllowFreesale       bool   `json:"allowFreesale"`
    InstantConfirmation bool   `json:"instantConfirmation"`
    InstantDelivery     bool   `json:"instantDelivery"`
    CreatedAt           string `json:"createdAt"`
    UpdatedAt           string `json:"updatedAt"`
}

func ToProductResponse(p *repository.Product) *ProductResponse {
    return &ProductResponse{
        ID:                  p.ID.String(),
        InternalName:        p.InternalName,
        Locale:              p.Locale,
        TimeZone:            p.TimeZone,
        AllowFreesale:       p.AllowFreesale,
        InstantConfirmation: p.InstantConfirmation,
        InstantDelivery:     p.InstantDelivery,
        CreatedAt:           p.CreatedAt.Format(time.RFC3339),
        UpdatedAt:           p.UpdatedAt.Format(time.RFC3339),
    }
}

// internal/http/handler/product_handler.go
type ProductHandler struct {
    productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
    return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    id, err := uuid.Parse(chi.URLParam(r, "id"))
    if err != nil {
        respondError(w, http.StatusBadRequest, "Invalid ID")
        return
    }
    
    product, err := h.productService.GetProduct(r.Context(), id)
    if err != nil {
        respondError(w, http.StatusNotFound, "Product not found")
        return
    }
    
    // Convert internal model to API response DTO
    response := dto.ToProductResponse(product)
    respondJSON(w, http.StatusOK, response)
}
```

### 5. Container (Dependency Injection)

**Purpose**: Centralized dependency management and lifecycle

**Responsibilities**:
- Initialize all dependencies in correct order
- Inject dependencies into constructors
- Manage resource lifecycle

**Example**:
```go
// internal/container/container.go
type Container struct {
    DB               *sql.DB
    Queries          *repository.Queries
    ProductRepo      interfaces.ProductRepository
    BookingRepo      interfaces.BookingRepository
    ProductService   *service.ProductService
    BookingService   *service.BookingService
    ProductHandler   *handler.ProductHandler
    BookingHandler   *handler.BookingHandler
}

func NewContainer(db *sql.DB) *Container {
    c := &Container{DB: db}
    
    // Initialize SQLc queries
    c.Queries = repository.New(db)
    
    // Initialize repositories
    c.ProductRepo = interfaces.NewProductRepository(c.Queries)
    c.BookingRepo = interfaces.NewBookingRepository(c.Queries)
    
    // Initialize services
    c.ProductService = service.NewProductService(c.ProductRepo)
    c.BookingService = service.NewBookingService(c.BookingRepo, c.ProductRepo)
    
    // Initialize handlers
    c.ProductHandler = handler.NewProductHandler(c.ProductService)
    c.BookingHandler = handler.NewBookingHandler(c.BookingService)
    
    return c
}
```

## SQLc Integration

### Configuration

SQLc is configured to generate optimal code for our architecture:

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "internal/database/queries"
    schema: "internal/database/migrations"
    gen:
      go:
        package: "repository"
        out: "internal/repository"
        sql_package: "pgx/v5"
        emit_json_tags: true              # Generate JSON tags
        emit_prepared_queries: false
        emit_interface: false
        emit_exact_table_names: false
        emit_empty_slices: true           # Return [] instead of nil
        json_tags_case_style: "camel"     # camelCase for REST APIs
        overrides:
          - db_type: "uuid"
            go_type:
              import: "github.com/google/uuid"
              type: "UUID"
          - db_type: "timestamptz"
            go_type:
              import: "time"
              type: "Time"
```

### Query Organization

Queries are organized by domain in separate SQL files:

```sql
-- internal/database/queries/product.sql

-- name: GetProductByID :one
SELECT * FROM product 
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListProducts :many
SELECT * FROM product 
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateProduct :exec
INSERT INTO product (
    id, internal_name, locale, time_zone,
    allow_freesale, instant_confirmation, instant_delivery,
    availability_required, availability_type, redemption_method,
    freesale_duration_amount, freesale_duration_unit,
    product_content_id, product_pricing_id, product_package_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
);
```

### Complex Queries with JOINs

SQLc handles complex queries and generates appropriate structs:

```sql
-- name: GetBookingWithDetails :one
SELECT 
    b.id, b.uuid, b.status, b.utc_created_at,
    b.cancellable, b.freesale, b.notes,
    p.id as product_id,
    p.internal_name as product_name,
    p.locale as product_locale,
    c.id as contact_id,
    c.first_name, c.last_name,
    c.email_address, c.phone_number,
    o.id as option_id,
    o.internal_name as option_name
FROM booking b
JOIN product p ON b.product_id = p.id
JOIN contact c ON b.contact_id = c.id
JOIN option o ON b.option_id = o.id
WHERE b.uuid = $1 AND b.deleted_at IS NULL;
```

SQLc generates:
```go
type GetBookingWithDetailsRow struct {
    ID            uuid.UUID     `json:"id"`
    Uuid          uuid.UUID     `json:"uuid"`
    Status        BookingStatus `json:"status"`
    UtcCreatedAt  time.Time     `json:"utcCreatedAt"`
    Cancellable   bool          `json:"cancellable"`
    Freesale      bool          `json:"freesale"`
    Notes         pgtype.Text   `json:"notes"`
    ProductID     uuid.UUID     `json:"productId"`
    ProductName   string        `json:"productName"`
    ProductLocale string        `json:"productLocale"`
    ContactID     uuid.UUID     `json:"contactId"`
    FirstName     pgtype.Text   `json:"firstName"`
    LastName      pgtype.Text   `json:"lastName"`
    EmailAddress  pgtype.Text   `json:"emailAddress"`
    PhoneNumber   pgtype.Text   `json:"phoneNumber"`
    OptionID      uuid.UUID     `json:"optionId"`
    OptionName    string        `json:"optionName"`
}
```

## Data Flow

### Read Operation Flow

```
HTTP Request
    ↓
Handler (validate input)
    ↓
Service (business logic)
    ↓
Repository Interface
    ↓
SQLc Generated Query
    ↓
PostgreSQL
    ↓
SQLc Generated Model
    ↓
Service (transform if needed)
    ↓
Handler (serialize to JSON)
    ↓
HTTP Response
```

### Write Operation Flow

```
HTTP Request (with JSON body)
    ↓
Handler (decode & validate)
    ↓
Service (business logic & validation)
    ↓
Repository Interface
    ↓
SQLc Generated Mutation
    ↓
PostgreSQL Transaction
    ↓
Commit/Rollback
    ↓
Handler (respond with status)
    ↓
HTTP Response
```

## Testing Strategy

### Unit Tests

Each layer is tested independently using mocks:

```go
// Repository tests use test database
func TestProductRepository_GetByID(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    repo := interfaces.NewProductRepository(repository.New(db))
    
    product, err := repo.GetByID(context.Background(), testProductID)
    assert.NoError(t, err)
    assert.Equal(t, "test-product", product.InternalName)
}

// Service tests use mock repositories
func TestProductService_GetProduct(t *testing.T) {
    mockRepo := new(MockProductRepository)
    service := service.NewProductService(mockRepo)
    
    mockRepo.On("GetByID", mock.Anything, testID).Return(testProduct, nil)
    
    product, err := service.GetProduct(context.Background(), testID)
    assert.NoError(t, err)
    assert.Equal(t, testProduct.ID, product.ID)
}

// Handler tests use mock services
func TestProductHandler_GetProduct(t *testing.T) {
    mockService := new(MockProductService)
    handler := handler.NewProductHandler(mockService)
    
    mockService.On("GetProduct", mock.Anything, testID).Return(&testProduct, nil)
    
    req := httptest.NewRequest("GET", "/products/"+testID.String(), nil)
    w := httptest.NewRecorder()
    
    handler.GetProduct(w, req)
    
    assert.Equal(t, http.StatusOK, w.Code)
}
```

### Integration Tests

Integration tests verify the full stack using a test database:

```go
func TestCreateBookingIntegration(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    container := container.NewContainer(db)
    router := http.NewRouter(container)
    
    body := `{
        "productId": "...",
        "optionId": "...",
        "contact": {...}
    }`
    
    req := httptest.NewRequest("POST", "/api/v1/bookings", strings.NewReader(body))
    w := httptest.NewRecorder()
    
    router.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusCreated, w.Code)
}
```

## Database Migrations

Database schema changes are managed through versioned migrations using golang-migrate.

### Migration Files

```
internal/database/migrations/
├── 000001_initial_schema.up.sql
├── 000001_initial_schema.down.sql
├── 000002_add_payment_tables.up.sql
└── 000002_add_payment_tables.down.sql
```

### Migration Commands

```bash
make migrate-up           # Apply all pending migrations
make migrate-down         # Rollback last migration
make migrate-version      # Check current version
make migrate-create NAME=add_feature  # Create new migration
```

See [MIGRATIONS.md](./MIGRATIONS.md) for detailed migration documentation.

## Code Generation Workflow

```bash
# 1. Write/modify SQL schema
vim internal/database/migrations/000002_add_feature.up.sql

# 2. Apply migration
make migrate-up

# 3. Write SQL queries
vim internal/database/queries/feature.sql

# 4. Generate Go code
make sqlc-generate

# 5. Implement business logic
vim internal/service/feature_service.go

# 6. Implement handlers
vim internal/http/handler/feature_handler.go

# 7. Wire dependencies
vim internal/container/container.go

# 8. Add routes
vim internal/http/routes.go
```

## Error Handling

### Error Types

```go
// internal/pkg/errors/errors.go
var (
    ErrNotFound         = errors.New("resource not found")
    ErrInvalidInput     = errors.New("invalid input")
    ErrUnauthorized     = errors.New("unauthorized")
    ErrForbidden        = errors.New("forbidden")
    ErrConflict         = errors.New("resource conflict")
    ErrInternalServer   = errors.New("internal server error")
)

type AppError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details any    `json:"details,omitempty"`
}
```

### Error Response Format

```json
{
    "error": {
        "code": "BOOKING_NOT_FOUND",
        "message": "Booking with UUID abc123 not found",
        "details": {
            "uuid": "abc123"
        }
    }
}
```

## Configuration Management

Configuration is managed through environment variables:

```bash
# Database
DATABASE_URL=postgres://user:pass@localhost:5432/passos?sslmode=disable

# Server
PORT=8080
ENV=development

# Feature flags
ENABLE_OCTO_VALIDATION=true
```

## Performance Considerations

### Database Connection Pool

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxLifetime(5 * time.Minute)
```

### Query Optimization

- Use appropriate indexes (defined in migrations)
- Leverage PostgreSQL prepared statements via SQLc
- Implement pagination for list endpoints
- Use JOINs efficiently to minimize round trips

### Caching Strategy

```go
// Future: Implement caching layer
// - Redis for frequently accessed data
// - Cache invalidation on mutations
// - TTL-based expiration
```

## Security Considerations

### Input Validation

All input is validated at the handler layer before processing:

```go
type CreateBookingRequest struct {
    ProductID uuid.UUID `json:"productId" validate:"required,uuid"`
    OptionID  uuid.UUID `json:"optionId" validate:"required,uuid"`
    Contact   Contact   `json:"contact" validate:"required"`
}
```

### Authentication & Authorization

```go
// internal/http/middleware/auth.go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        // Validate token
        // Add user context
        next.ServeHTTP(w, r)
    })
}
```

## API Contract vs Internal Models

### The DTO Layer Rationale

One of the most critical architectural decisions is maintaining a clear separation between the public API contract (DTOs) and internal data models (SQLc-generated structs).

### Why DTOs are Essential

**1. API Stability**
```go
// Database schema can change
ALTER TABLE product ADD COLUMN new_internal_field TEXT;
ALTER TABLE product RENAME COLUMN time_zone TO timezone;

// But API contract remains stable
type ProductResponse struct {
    ID       string `json:"id"`
    TimeZone string `json:"timeZone"`  // Still camelCase, same field name
    // new_internal_field is NOT exposed
}
```

**2. Database Evolution**
```go
// Version 1: Simple structure
type Product struct {
    Name string
}

// Version 2: Normalized structure (breaking change in DB)
type Product struct {
    ProductContentID uuid.UUID
}
type ProductContent struct {
    Name string
}

// But API stays the same
type ProductResponse struct {
    Name string `json:"name"`  // Populated from ProductContent now
}
```

**3. Type Safety for Clients**
```go
// Bad: Exposing internal types
type BadResponse struct {
    ID        uuid.UUID          `json:"id"`           // Clients must handle UUID
    UpdatedAt pgtype.Timestamptz `json:"updatedAt"`    // Clients must handle pgtype
    Notes     pgtype.Text        `json:"notes"`        // Nullable wrapper exposed
}

// Good: Clean API types
type GoodResponse struct {
    ID        string  `json:"id"`           // Simple string
    UpdatedAt string  `json:"updatedAt"`    // ISO 8601 string
    Notes     *string `json:"notes"`        // Native Go nullable
}
```

**4. API Versioning**
```go
// v1 API
type ProductResponseV1 struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// v2 API - different format, same database
type ProductResponseV2 struct {
    ProductID   string          `json:"productId"`    // Field renamed
    ProductInfo ProductInfoV2   `json:"productInfo"`  // Nested structure
}

// Both map from the same repository.Product
```

**5. OCTO Compliance**
```go
// OCTO requires specific formats
type OCTOProductResponse struct {
    ID               string              `json:"id"`
    InternalName     string              `json:"internalName"`
    Reference        string              `json:"reference"`
    Locale           string              `json:"locale"`
    TimeZone         string              `json:"timeZone"`
    AllowFreesale    bool                `json:"allowFreesale"`
    AvailabilityType string              `json:"availabilityType"`
    Options          []OCTOOptionSummary `json:"options"`
    // OCTO-specific fields that don't exist in DB
    SupplierID       string              `json:"supplierId"`
    Currency         string              `json:"currency"`
}

// Map from internal models with additional logic
func ToOCTOProductResponse(p *repository.Product, opts []repository.Option) *OCTOProductResponse {
    return &OCTOProductResponse{
        ID:               p.ID.String(),
        InternalName:     p.InternalName,
        // ... map fields
        SupplierID:       "PASSOS", // Business logic
        Currency:         determineCurrency(p.Locale),
        Options:          mapOptions(opts),
    }
}
```

### DTO Best Practices

**1. Keep DTOs Simple**
```go
// Good: Simple, flat structure
type BookingResponse struct {
    ID            string `json:"id"`
    Status        string `json:"status"`
    ProductName   string `json:"productName"`
    ContactEmail  string `json:"contactEmail"`
    TotalAmount   int    `json:"totalAmount"`
    Currency      string `json:"currency"`
}

// Avoid: Exposing internal complexity
type BadBookingResponse struct {
    Booking repository.Booking           // Entire DB model
    Product repository.Product           // Entire DB model
    Contact repository.Contact           // Entire DB model
}
```

**2. Use Consistent Naming**
```go
// Consistent API naming conventions
type ProductResponse struct {
    ID        string `json:"id"`           // Not "productId"
    Name      string `json:"name"`         // Not "productName"
    CreatedAt string `json:"createdAt"`    // ISO 8601
}
```

**3. Handle Nulls Appropriately**
```go
// Convert pgtype nullables to Go pointers
func ToProductResponse(p *repository.Product) *ProductResponse {
    var reference *string
    if p.Reference.Valid {
        reference = &p.Reference.String
    }
    
    return &ProductResponse{
        ID:        p.ID.String(),
        Reference: reference,  // nil if not present
    }
}
```

**4. Validate Input DTOs**
```go
type CreateProductRequest struct {
    InternalName string `json:"internalName" validate:"required,min=3,max=100"`
    Locale       string `json:"locale" validate:"required,iso639_1"`
    TimeZone     string `json:"timeZone" validate:"required,timezone"`
}

func (r *CreateProductRequest) Validate() error {
    // Business-level validation
    if !isValidLocale(r.Locale) {
        return errors.New("unsupported locale")
    }
    return nil
}
```

### Mapping Strategies

**1. Simple Mappers**
```go
func ToProductResponse(p *repository.Product) *ProductResponse {
    return &ProductResponse{
        ID:           p.ID.String(),
        InternalName: p.InternalName,
        Locale:       p.Locale,
        CreatedAt:    p.CreatedAt.Format(time.RFC3339),
    }
}
```

**2. Complex Mappers with Aggregation**
```go
func ToBookingDetailResponse(
    b *repository.Booking,
    p *repository.Product,
    c *repository.Contact,
    items []repository.UnitItem,
) *BookingDetailResponse {
    return &BookingDetailResponse{
        ID:          b.UUID.String(),
        Status:      string(b.Status),
        Product:     ToProductSummary(p),
        Contact:     ToContactInfo(c),
        Items:       mapUnitItems(items),
        TotalAmount: calculateTotal(items),
    }
}
```

**3. Batch Mappers**
```go
func ToProductListResponse(products []repository.Product) []ProductSummaryResponse {
    result := make([]ProductSummaryResponse, len(products))
    for i, p := range products {
        result[i] = *ToProductSummary(&p)
    }
    return result
}
```

### Performance Considerations

**1. Avoid N+1 Queries**
```go
// Bad: N+1 query problem
func (s *ProductService) GetProductsWithOptions(ctx context.Context) ([]*ProductWithOptionsDTO, error) {
    products, _ := s.productRepo.List(ctx)
    
    result := make([]*ProductWithOptionsDTO, len(products))
    for i, p := range products {
        options, _ := s.optionRepo.GetByProductID(ctx, p.ID) // N queries!
        result[i] = ToProductWithOptionsDTO(&p, options)
    }
    return result, nil
}

// Good: Single query with JOIN
func (s *ProductService) GetProductsWithOptions(ctx context.Context) ([]*ProductWithOptionsDTO, error) {
    // Use SQLc query with JOIN
    rows, _ := s.queries.GetProductsWithOptions(ctx)
    return mapToProductWithOptionsDTO(rows), nil
}
```

**2. Lazy Loading in DTOs**
```go
type ProductDetailResponse struct {
    ID      string                `json:"id"`
    Name    string                `json:"name"`
    Options []OptionSummary       `json:"options"`
    Stats   *ProductStats         `json:"stats,omitempty"` // Only if requested
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    includeStats := r.URL.Query().Get("includeStats") == "true"
    
    product, _ := h.service.GetProduct(ctx, id)
    response := dto.ToProductResponse(product)
    
    if includeStats {
        stats, _ := h.service.GetProductStats(ctx, id)
        response.Stats = dto.ToProductStatsDTO(stats)
    }
    
    respondJSON(w, http.StatusOK, response)
}
```

## OCTO Compliance

The architecture supports OCTO (Open Connectivity for Tourism Operations) compliance:


## Monitoring and Observability

### Logging

Structured logging using standard library or third-party logger:

```go
log.Info("booking created",
    "booking_id", booking.ID,
    "product_id", booking.ProductID,
    "status", booking.Status,
)
```

### Metrics

```go
// Track key metrics
metricsHandler.IncrementBookingCount()
metricsHandler.RecordRequestDuration(duration)
```

### Health Checks

```go
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
    health := map[string]string{
        "status": "healthy",
        "database": h.checkDatabase(),
    }
    respondJSON(w, http.StatusOK, health)
}
```

## Deployment

### Docker Support

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:latest
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]
```

### Environment-Specific Configuration

```bash
# Development
ENV=development
LOG_LEVEL=debug

# Production
ENV=production
LOG_LEVEL=info
```

## Best Practices

1. **Always Use DTOs for Public APIs**: Never expose SQLc models directly in HTTP responses to maintain API stability
2. **Keep Queries Simple**: Write clear, readable SQL in separate files organized by domain
3. **Use Transactions**: Wrap multi-step operations in database transactions
4. **Validate Early**: Validate input at the handler layer using request DTOs
5. **Handle Nulls Properly**: Use pgtype internally, convert to Go pointers in DTOs
6. **Separate Concerns**: SQLc models for data layer, DTOs for API layer
7. **Write Tests**: Maintain high test coverage across all layers
8. **Document APIs**: Use OpenAPI/Swagger for API documentation with DTO definitions
9. **Version APIs**: Use versioned endpoints (/api/v1/) with version-specific DTOs
10. **Monitor Performance**: Track query performance and optimize as needed
11. **Use Mapper Functions**: Create explicit mapper functions (ToXXXResponse) for clarity
12. **Plan for Evolution**: Design DTOs with future changes in mind

## Future Enhancements

- [ ] Implement caching layer (Redis)
- [ ] Add API rate limiting
- [ ] Implement webhook support for external integrations
- [ ] Add comprehensive API documentation (OpenAPI)
- [ ] Implement event sourcing for audit trails
- [ ] Add GraphQL endpoint as alternative to REST
- [ ] Implement distributed tracing (OpenTelemetry)
- [ ] Add support for multiple database replicas

## Contributing

When contributing to the codebase:

1. Follow the established architecture patterns
2. Write tests for new features
3. Update documentation as needed
4. Use SQLc for database interactions
5. Maintain separation of concerns across layers
6. Follow Go best practices and idioms

## Key Architectural Decisions

### Decision 1: DTO Layer for All Public APIs

**Decision**: Always use DTOs as the contract layer between public APIs and internal models, even when SQLc models could technically be serialized directly.

**Rationale**:
- Prevents coupling of public API to internal database schema
- Enables independent evolution of database and API
- Provides stable contracts for API consumers
- Facilitates API versioning (v1, v2, etc.)
- Hides internal implementation details
- Supports OCTO compliance with format-specific responses

**Trade-offs**:
- Additional code for DTOs and mappers
- Slight performance overhead for mapping
- More initial development time

**Benefits**:
- Long-term maintainability and flexibility
- Breaking changes in database don't break API
- Clean separation of concerns
- Professional API design
- Easier to document and version

### Decision 2: SQLc for Type-Safe Data Access

**Decision**: Use SQLc to generate type-safe Go code from SQL queries instead of ORM or manual SQL.

**Rationale**:
- Maintains SQL as first-class citizen
- Zero runtime overhead (no reflection)
- Compile-time type safety
- Database schema as single source of truth
- Excellent performance

**Trade-offs**:
- Requires code generation step
- Less abstraction than ORMs
- Need to write SQL manually

**Benefits**:
- Performance and type safety
- Clear and explicit queries
- Easy to optimize
- Great for complex queries

### Decision 3: Repository Pattern with Interfaces

**Decision**: Use repository interfaces for dependency injection and testing, even though SQLc generates the implementations.

**Rationale**:
- Enables dependency injection
- Facilitates unit testing with mocks
- Clear contracts between layers
- Flexibility to change implementations

**Trade-offs**:
- Additional interface definitions
- One more layer of abstraction

**Benefits**:
- Testable code
- Loose coupling
- Clean architecture principles

## References

- [SQLc Documentation](https://docs.sqlc.dev/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Chi Router](https://github.com/go-chi/chi)
- [Repository Pattern](https://martinfowler.com/eaaCatalog/repository.html)
- [OCTO Specification](https://octospec.com/)
- [API Design Best Practices](https://restfulapi.net/)

---

**Last Updated**: 2025-01-23
**Architecture Version**: 1.1

**Changelog**:
- v1.1: Added critical section on DTO layer rationale and API contract separation
- v1.0: Initial architecture documentation

