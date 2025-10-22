package availability

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Availability struct {
	ID                 uuid.UUID                `json:"id" db:"id"`
	OptionID           uuid.UUID                `json:"optionId" db:"option_id"`
	LocalDateTimeStart time.Time                `json:"localDateTimeStart" db:"local_date_time_start"`
	LocalDateTimeEnd   time.Time                `json:"localDateTimeEnd" db:"local_date_time_end"`
	AllDay             bool                     `json:"allDay" db:"all_day"`
	Available          bool                     `json:"available" db:"available"`
	Status             types.AvailabilityStatus `json:"status" db:"status"`
	Vacancies          *int                     `json:"vacancies,omitempty" db:"vacancies"`
	Capacity           *int                     `json:"capacity,omitempty" db:"capacity"`
	MaxUnits           *int                     `json:"maxUnits,omitempty" db:"max_units"`
	UtcCutoffAt        time.Time                `json:"utcCutoffAt" db:"utc_cutoff_at"`

	MeetingPointID       *uuid.UUID `json:"meetingPointId,omitempty" db:"meeting_point_id"`
	MeetingLocalDateTime *time.Time `json:"meetingLocalDateTime,omitempty" db:"meeting_local_date_time"`
	TourGroupID          *uuid.UUID `json:"tourGroupId,omitempty" db:"tour_group_id"`

	PickupAvailable *bool `json:"pickupAvailable,omitempty" db:"pickup_available"`
	PickupRequired  *bool `json:"pickupRequired,omitempty" db:"pickup_required"`

	OfferCode  *string    `json:"offerCode,omitempty" db:"offer_code"`
	OfferTitle *string    `json:"offerTitle,omitempty" db:"offer_title"`
	OfferID    *uuid.UUID `json:"offerId,omitempty" db:"offer_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OpeningHours struct {
	ID              uuid.UUID `json:"id" db:"id"`
	AvailabilityID  uuid.UUID `json:"availabilityId" db:"availability_id"`
	From            string    `json:"from" db:"from"`
	To              string    `json:"to" db:"to"`
	Frequency       *string   `json:"frequency,omitempty" db:"frequency"`
	FrequencyAmount *int      `json:"frequencyAmount,omitempty" db:"frequency_amount"`
	FrequencyUnit   *string   `json:"frequencyUnit,omitempty" db:"frequency_unit"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type AvailabilityNotice struct {
	ID             uuid.UUID `json:"id" db:"id"`
	AvailabilityID uuid.UUID `json:"availabilityId" db:"availability_id"`
	NoticeID       uuid.UUID `json:"noticeId" db:"notice_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type AvailabilityPickupPoint struct {
	ID              uuid.UUID `json:"id" db:"id"`
	AvailabilityID  uuid.UUID `json:"availabilityId" db:"availability_id"`
	PickupPointID   uuid.UUID `json:"pickupPointId" db:"pickup_point_id"`
	LocalDateTime   string    `json:"localDateTime" db:"local_date_time"`
	LocalDateTimeTo *string   `json:"localDateTimeTo,omitempty" db:"local_date_time_to"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type AvailabilityOffer struct {
	ID             uuid.UUID `json:"id" db:"id"`
	AvailabilityID uuid.UUID `json:"availabilityId" db:"availability_id"`
	OfferID        uuid.UUID `json:"offerId" db:"offer_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type AvailabilityResource struct {
	ID             uuid.UUID `json:"id" db:"id"`
	AvailabilityID uuid.UUID `json:"availabilityId" db:"availability_id"`
	ResourceID     uuid.UUID `json:"resourceId" db:"resource_id"`
	Quantity       int       `json:"quantity" db:"quantity"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type TourGroup struct {
	ID           uuid.UUID `json:"id" db:"id"`
	InternalName string    `json:"internalName" db:"internal_name"`
	Title        string    `json:"title" db:"title"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
