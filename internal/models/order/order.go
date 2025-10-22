package order

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Order struct {
	ID                uuid.UUID         `json:"id" db:"id"`
	TestMode          bool              `json:"testMode" db:"test_mode"`
	SupplierReference string            `json:"supplierReference" db:"supplier_reference"`
	SettlementMethod  string            `json:"settlementMethod" db:"settlement_method"`
	Status            types.OrderStatus `json:"status" db:"status"`
	UtcExpiresAt      *time.Time        `json:"utcExpiresAt,omitempty" db:"utc_expires_at"`
	UtcConfirmedAt    *time.Time        `json:"utcConfirmedAt,omitempty" db:"utc_confirmed_at"`
	Cancellable       bool              `json:"cancellable" db:"cancellable"`
	ContactID         uuid.UUID         `json:"contactId" db:"contact_id"`
	TermsAccepted     *bool             `json:"termsAccepted,omitempty" db:"terms_accepted"`
	ReturnUrl         *string           `json:"returnUrl,omitempty" db:"return_url"`
	Confirmable       *bool             `json:"confirmable,omitempty" db:"confirmable"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OrderBooking struct {
	ID        uuid.UUID `json:"id" db:"id"`
	OrderID   uuid.UUID `json:"orderId" db:"order_id"`
	BookingID uuid.UUID `json:"bookingId" db:"booking_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
