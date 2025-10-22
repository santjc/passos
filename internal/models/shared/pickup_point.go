package shared

import (
	"time"

	"github.com/google/uuid"
)

type PickupPoint struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Directions      *string   `json:"directions,omitempty" db:"directions"`
	Address         string    `json:"address" db:"address"`
	Latitude        *float64  `json:"latitude,omitempty" db:"latitude"`
	Longitude       *float64  `json:"longitude,omitempty" db:"longitude"`
	GooglePlaceId   *string   `json:"googlePlaceId,omitempty" db:"google_place_id"`
	Street          *string   `json:"street,omitempty" db:"street"`
	PostalCode      *string   `json:"postalCode,omitempty" db:"postal_code"`
	Locality        *string   `json:"locality,omitempty" db:"locality"`
	Region          *string   `json:"region,omitempty" db:"region"`
	State           *string   `json:"state,omitempty" db:"state"`
	Country         *string   `json:"country,omitempty" db:"country"`
	LocalDateTime   string    `json:"localDateTime" db:"local_date_time"`
	LocalDateTimeTo *string   `json:"localDateTimeTo,omitempty" db:"local_date_time_to"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
