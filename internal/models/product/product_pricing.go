package product

import (
	"passos/internal/types"
	"time"

	"github.com/google/uuid"
)

type ProductPricing struct {
	ID              uuid.UUID         `json:"id" db:"id"`
	DefaultCurrency *string           `json:"defaultCurrency,omitempty" db:"default_currency"`
	PricingPer      *types.PricingPer `json:"pricingPer,omitempty" db:"pricing_per"`
	IncludeTax      *bool             `json:"includeTax,omitempty" db:"include_tax"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type ProductPricingCurrency struct {
	ID               uuid.UUID `json:"id" db:"id"`
	ProductPricingID uuid.UUID `json:"productPricingId" db:"product_pricing_id"`
	Currency         string    `json:"currency" db:"currency"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
