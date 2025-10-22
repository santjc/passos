package booking

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Booking struct {
	ID                uuid.UUID           `json:"id" db:"id"`
	UUID              uuid.UUID           `json:"uuid" db:"uuid"`
	TestMode          bool                `json:"testMode" db:"test_mode"`
	ResellerReference *string             `json:"resellerReference,omitempty" db:"reseller_reference"`
	SupplierReference *string             `json:"supplierReference,omitempty" db:"supplier_reference"`
	Status            types.BookingStatus `json:"status" db:"status"`
	UtcCreatedAt      time.Time           `json:"utcCreatedAt" db:"utc_created_at"`
	UtcUpdatedAt      *time.Time          `json:"utcUpdatedAt,omitempty" db:"utc_updated_at"`
	UtcExpiresAt      *time.Time          `json:"utcExpiresAt,omitempty" db:"utc_expires_at"`
	UtcRedeemedAt     *time.Time          `json:"utcRedeemedAt,omitempty" db:"utc_redeemed_at"`
	UtcConfirmedAt    *time.Time          `json:"utcConfirmedAt,omitempty" db:"utc_confirmed_at"`
	ProductID         uuid.UUID           `json:"productId" db:"product_id"`
	OptionID          uuid.UUID           `json:"optionId" db:"option_id"`
	Cancellable       bool                `json:"cancellable" db:"cancellable"`
	Freesale          bool                `json:"freesale" db:"freesale"`
	AvailabilityID    *uuid.UUID          `json:"availabilityId,omitempty" db:"availability_id"`
	ContactID         uuid.UUID           `json:"contactId" db:"contact_id"`
	Notes             *string             `json:"notes,omitempty" db:"notes"`

	MeetingPointID       *uuid.UUID `json:"meetingPointId,omitempty" db:"meeting_point_id"`
	MeetingLocalDateTime *time.Time `json:"meetingLocalDateTime,omitempty" db:"meeting_local_date_time"`
	Duration             *string    `json:"duration,omitempty" db:"duration"`
	DurationAmount       *string    `json:"durationAmount,omitempty" db:"duration_amount"`
	DurationUnit         *string    `json:"durationUnit,omitempty" db:"duration_unit"`
	TermsAccepted        *bool      `json:"termsAccepted,omitempty" db:"terms_accepted"`

	PickupRequested *bool      `json:"pickupRequested,omitempty" db:"pickup_requested"`
	PickupPointID   *uuid.UUID `json:"pickupPointId,omitempty" db:"pickup_point_id"`
	PickupHotel     *string    `json:"pickupHotel,omitempty" db:"pickup_hotel"`
	PickupHotelRoom *string    `json:"pickupHotelRoom,omitempty" db:"pickup_hotel_room"`

	OrderID        *string `json:"orderId,omitempty" db:"order_id"`
	OrderReference *string `json:"orderReference,omitempty" db:"order_reference"`
	Primary        *bool   `json:"primary,omitempty" db:"primary"`

	OfferCode          *string    `json:"offerCode,omitempty" db:"offer_code"`
	OfferTitle         *string    `json:"offerTitle,omitempty" db:"offer_title"`
	OfferIsCombination *bool      `json:"offerIsCombination,omitempty" db:"offer_is_combination"`
	OfferID            *uuid.UUID `json:"offerId,omitempty" db:"offer_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type Cancellation struct {
	ID             uuid.UUID `json:"id" db:"id"`
	BookingID      uuid.UUID `json:"bookingId" db:"booking_id"`
	Refund         string    `json:"refund" db:"refund"`
	Reason         *string   `json:"reason,omitempty" db:"reason"`
	UtcCancelledAt time.Time `json:"utcCancelledAt" db:"utc_cancelled_at"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Contact struct {
	ID             uuid.UUID `json:"id" db:"id"`
	FullName       *string   `json:"fullName,omitempty" db:"full_name"`
	FirstName      *string   `json:"firstName,omitempty" db:"first_name"`
	LastName       *string   `json:"lastName,omitempty" db:"last_name"`
	EmailAddress   *string   `json:"emailAddress,omitempty" db:"email_address"`
	PhoneNumber    *string   `json:"phoneNumber,omitempty" db:"phone_number"`
	PostalCode     *string   `json:"postalCode,omitempty" db:"postal_code"`
	Country        *string   `json:"country,omitempty" db:"country"`
	Notes          *string   `json:"notes,omitempty" db:"notes"`
	AllowMarketing *bool     `json:"allowMarketing,omitempty" db:"allow_marketing"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type BookingContact struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BookingID uuid.UUID `json:"bookingId" db:"booking_id"`
	ContactID uuid.UUID `json:"contactId" db:"contact_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type ContactLocale struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ContactID uuid.UUID `json:"contactId" db:"contact_id"`
	Locale    string    `json:"locale" db:"locale"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Ticket struct {
	ID               uuid.UUID              `json:"id" db:"id"`
	BookingID        uuid.UUID              `json:"bookingId" db:"booking_id"`
	RedemptionMethod types.RedemptionMethod `json:"redemptionMethod" db:"redemption_method"`
	UtcRedeemedAt    *time.Time             `json:"utcRedeemedAt,omitempty" db:"utc_redeemed_at"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type DeliveryOption struct {
	ID             uuid.UUID            `json:"id" db:"id"`
	TicketID       uuid.UUID            `json:"ticketId" db:"ticket_id"`
	DeliveryFormat types.DeliveryFormat `json:"deliveryFormat" db:"delivery_format"`
	DeliveryValue  string               `json:"deliveryValue" db:"delivery_value"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type BookingDeliveryMethod struct {
	ID             uuid.UUID            `json:"id" db:"id"`
	BookingID      uuid.UUID            `json:"bookingId" db:"booking_id"`
	DeliveryMethod types.DeliveryMethod `json:"deliveryMethod" db:"delivery_method"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type BookingQuestionAnswer struct {
	ID               uuid.UUID `json:"id" db:"id"`
	BookingID        uuid.UUID `json:"bookingId" db:"booking_id"`
	QuestionAnswerID uuid.UUID `json:"questionAnswerId" db:"question_answer_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type BookingNotice struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BookingID uuid.UUID `json:"bookingId" db:"booking_id"`
	NoticeID  uuid.UUID `json:"noticeId" db:"notice_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Notice struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Title            *string   `json:"title,omitempty" db:"title"`
	ShortDescription string    `json:"shortDescription" db:"short_description"`
	CoverImageUrl    *string   `json:"coverImageUrl,omitempty" db:"cover_image_url"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
