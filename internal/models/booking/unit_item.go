package booking

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type UnitItem struct {
	ID                uuid.UUID           `json:"id" db:"id"`
	UUID              uuid.UUID           `json:"uuid" db:"uuid"`
	BookingID         uuid.UUID           `json:"bookingId" db:"booking_id"`
	UnitID            uuid.UUID           `json:"unitId" db:"unit_id"`
	ResellerReference *string             `json:"resellerReference,omitempty" db:"reseller_reference"`
	SupplierReference *string             `json:"supplierReference,omitempty" db:"supplier_reference"`
	Status            types.BookingStatus `json:"status" db:"status"`
	UtcRedeemedAt     *time.Time          `json:"utcRedeemedAt,omitempty" db:"utc_redeemed_at"`

	ContactID *uuid.UUID `json:"contactId,omitempty" db:"contact_id"`
	TicketID  *uuid.UUID `json:"ticketId,omitempty" db:"ticket_id"`
	PricingID *uuid.UUID `json:"pricingId,omitempty" db:"pricing_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
