package product

import (
	"time"

	"github.com/google/uuid"
)

type ProductPackage struct {
	ID        uuid.UUID `json:"id" db:"id"`
	IsPackage bool      `json:"isPackage" db:"is_package"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
