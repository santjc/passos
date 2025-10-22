package unit

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Unit struct {
	ID           uuid.UUID      `json:"id" db:"id"`
	OptionID     uuid.UUID      `json:"optionId" db:"option_id"`
	InternalName string         `json:"internalName" db:"internal_name"`
	Reference    string         `json:"reference" db:"reference"`
	Type         types.UnitType `json:"type" db:"type"`

	MinAge      int  `json:"minAge" db:"min_age"`
	MaxAge      int  `json:"maxAge" db:"max_age"`
	IdRequired  bool `json:"idRequired" db:"id_required"`
	MinQuantity *int `json:"minQuantity,omitempty" db:"min_quantity"`
	MaxQuantity *int `json:"maxQuantity,omitempty" db:"max_quantity"`
	PaxCount    int  `json:"paxCount" db:"pax_count"`

	UnitContentID *uuid.UUID `json:"unitContentId,omitempty" db:"unit_content_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type UnitAccompaniedBy struct {
	ID                    uuid.UUID      `json:"id" db:"id"`
	UnitID                uuid.UUID      `json:"unitId" db:"unit_id"`
	AccompaniedByUnitType types.UnitType `json:"accompaniedByUnitType" db:"accompanied_by_unit_type"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UnitContactField struct {
	ID         uuid.UUID          `json:"id" db:"id"`
	UnitID     uuid.UUID          `json:"unitId" db:"unit_id"`
	Field      types.ContactField `json:"field" db:"field"`
	IsRequired bool               `json:"isRequired" db:"is_required"`
	IsVisible  bool               `json:"isVisible" db:"is_visible"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
