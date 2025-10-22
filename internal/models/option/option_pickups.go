package option

import (
	"time"

	"github.com/google/uuid"
)

type OptionPickups struct {
	ID              uuid.UUID `json:"id" db:"id"`
	PickupRequired  *bool     `json:"pickupRequired,omitempty" db:"pickup_required"`
	PickupAvailable *bool     `json:"pickupAvailable,omitempty" db:"pickup_available"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type OptionPickupPoint struct {
	ID              uuid.UUID `json:"id" db:"id"`
	OptionPickupsID uuid.UUID `json:"optionPickupsId" db:"option_pickups_id"`
	PickupPointID   uuid.UUID `json:"pickupPointId" db:"pickup_point_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
