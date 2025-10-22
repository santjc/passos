package shared

import (
	"time"

	"github.com/google/uuid"
)

type ResourceGroup struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Split bool      `json:"split" db:"split"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type Resource struct {
	ID              uuid.UUID `json:"id" db:"id"`
	ResourceGroupID uuid.UUID `json:"resourceGroupId" db:"resource_group_id"`
	Title           string    `json:"title" db:"title"`
	Seating         bool      `json:"seating" db:"seating"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type Seat struct {
	ID         uuid.UUID `json:"id" db:"id"`
	ResourceID uuid.UUID `json:"resourceId" db:"resource_id"`
	Title      string    `json:"title" db:"title"`
	Column     int       `json:"column" db:"column"`
	Row        int       `json:"row" db:"row"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type ResourceAllocation struct {
	ID              uuid.UUID `json:"id" db:"id"`
	BookingID       uuid.UUID `json:"bookingId" db:"booking_id"`
	ResourceGroupID uuid.UUID `json:"resourceGroupId" db:"resource_group_id"`
	ResourceID      uuid.UUID `json:"resourceId" db:"resource_id"`
	PaxCount        int       `json:"paxCount" db:"pax_count"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type ResourceAllocationSeat struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	ResourceAllocationID uuid.UUID `json:"resourceAllocationId" db:"resource_allocation_id"`
	SeatID               uuid.UUID `json:"seatId" db:"seat_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
