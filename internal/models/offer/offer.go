package offer

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Offer struct {
	ID          uuid.UUID          `json:"id" db:"id"`
	ProductID   uuid.UUID          `json:"productId" db:"product_id"`
	OptionID    *uuid.UUID         `json:"optionId,omitempty" db:"option_id"`
	Title       string             `json:"title" db:"title"`
	Code        string             `json:"code" db:"code"`
	Description *string            `json:"description,omitempty" db:"description"`
	NetDiscount *types.NetDiscount `json:"netDiscount,omitempty" db:"net_discount"`
	Usable      bool               `json:"usable" db:"usable"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OfferRestrictions struct {
	ID       uuid.UUID `json:"id" db:"id"`
	OfferID  uuid.UUID `json:"offerId" db:"offer_id"`
	MinUnits *int      `json:"minUnits,omitempty" db:"min_units"`
	MaxUnits *int      `json:"maxUnits,omitempty" db:"max_units"`
	MinTotal *int      `json:"minTotal,omitempty" db:"min_total"`
	MaxTotal *int      `json:"maxTotal,omitempty" db:"max_total"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OfferRestrictionsUnit struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	OfferRestrictionsID uuid.UUID `json:"offerRestrictionsId" db:"offer_restrictions_id"`
	UnitID              uuid.UUID `json:"unitId" db:"unit_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type OfferDiscount struct {
	ID      uuid.UUID `json:"id" db:"id"`
	OfferID uuid.UUID `json:"offerId" db:"offer_id"`
	Net     *int      `json:"net,omitempty" db:"net"`
	Retail  int       `json:"retail" db:"retail"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OfferDiscountTax struct {
	ID              uuid.UUID `json:"id" db:"id"`
	OfferDiscountID uuid.UUID `json:"offerDiscountId" db:"offer_discount_id"`
	TaxID           uuid.UUID `json:"taxId" db:"tax_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
