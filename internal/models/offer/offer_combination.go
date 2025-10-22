package offer

import (
	"time"

	"github.com/google/uuid"
)

type OfferCombination struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	ProductID        uuid.UUID  `json:"productId" db:"product_id"`
	OptionID         uuid.UUID  `json:"optionId" db:"option_id"`
	OfferCode        string     `json:"offerCode" db:"offer_code"`
	OfferTitle       string     `json:"offerTitle" db:"offer_title"`
	ShortDescription *string    `json:"shortDescription,omitempty" db:"short_description"`
	PricingID        uuid.UUID  `json:"pricingId" db:"pricing_id"`
	BookingID        *uuid.UUID `json:"bookingId,omitempty" db:"booking_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OfferCombinationUnit struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	OfferCombinationID uuid.UUID `json:"offerCombinationId" db:"offer_combination_id"`
	UnitID             uuid.UUID `json:"unitId" db:"unit_id"`
	Quantity           int       `json:"quantity" db:"quantity"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
