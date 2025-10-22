package option

import (
	"passos/internal/types"
	"time"

	"github.com/google/uuid"
)

type Option struct {
	ID                       uuid.UUID `json:"id" db:"id"`
	ProductID                uuid.UUID `json:"productId" db:"product_id"`
	IsDefault                bool      `json:"isDefault" db:"is_default"`
	InternalName             string    `json:"internalName" db:"internal_name"`
	Reference                *string   `json:"reference,omitempty" db:"reference"`
	CancellationCutoff       string    `json:"cancellationCutoff" db:"cancellation_cutoff"`
	CancellationCutoffAmount int       `json:"cancellationCutoffAmount" db:"cancellation_cutoff_amount"`
	CancellationCutoffUnit   string    `json:"cancellationCutoffUnit" db:"cancellation_cutoff_unit"`
	AvailabilityCutoff       string    `json:"availabilityCutoff" db:"availability_cutoff"`
	AvailabilityCutoffAmount int       `json:"availabilityCutoffAmount" db:"availability_cutoff_amount"`
	AvailabilityCutoffUnit   string    `json:"availabilityCutoffUnit" db:"availability_cutoff_unit"`

	MinUnits    int `json:"minUnits" db:"min_units"`
	MaxUnits    int `json:"maxUnits" db:"max_units"`
	MinPaxCount int `json:"minPaxCount" db:"min_pax_count"`
	MaxPaxCount int `json:"maxPaxCount" db:"max_pax_count"`

	OptionContentID uuid.UUID  `json:"optionContentId" db:"option_content_id"`
	OptionPickupsID *uuid.UUID `json:"optionPickupsId,omitempty" db:"option_pickups_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OptionAvailabilityTime struct {
	ID        uuid.UUID `json:"id" db:"id"`
	OptionID  uuid.UUID `json:"optionId" db:"option_id"`
	LocalTime string    `json:"localTime" db:"local_time"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type OptionContactField struct {
	ID         uuid.UUID          `json:"id" db:"id"`
	OptionID   uuid.UUID          `json:"optionId" db:"option_id"`
	Field      types.ContactField `json:"field" db:"field"`
	IsRequired bool               `json:"isRequired" db:"is_required"`
	IsVisible  bool               `json:"isVisible" db:"is_visible"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
