package pricing

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Pricing struct {
	ID                uuid.UUID       `json:"id" db:"id"`
	OptionID          *uuid.UUID      `json:"optionId,omitempty" db:"option_id"`
	UnitID            *uuid.UUID      `json:"unitId,omitempty" db:"unit_id"`
	AvailabilityID    *uuid.UUID      `json:"availabilityId,omitempty" db:"availability_id"`
	BookingID         *uuid.UUID      `json:"bookingId,omitempty" db:"booking_id"`
	UnitItemID        *uuid.UUID      `json:"unitItemId,omitempty" db:"unit_item_id"`
	OrderID           *uuid.UUID      `json:"orderId,omitempty" db:"order_id"`
	PricingType       string          `json:"pricingType" db:"pricing_type"`
	UnitType          *types.UnitType `json:"unitType,omitempty" db:"unit_type"`
	Original          int             `json:"original" db:"original"`
	Retail            int             `json:"retail" db:"retail"`
	Net               *int            `json:"net,omitempty" db:"net"`
	Currency          string          `json:"currency" db:"currency"`
	CurrencyPrecision int             `json:"currencyPrecision" db:"currency_precision"`

	OfferDiscountID *uuid.UUID `json:"offerDiscountId,omitempty" db:"offer_discount_id"`
	ExtraID         *uuid.UUID `json:"extraId,omitempty" db:"extra_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type PricingTax struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PricingID uuid.UUID `json:"pricingId" db:"pricing_id"`
	TaxID     uuid.UUID `json:"taxId" db:"tax_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Tax struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Type        string    `json:"type" db:"type"`
	Description string    `json:"description" db:"description"`
	Amount      int       `json:"amount" db:"amount"`
	Percentage  *float64  `json:"percentage,omitempty" db:"percentage"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type ExtraPricing struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UnitID    uuid.UUID `json:"unitId" db:"unit_id"`
	PricingID uuid.UUID `json:"pricingId" db:"pricing_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
