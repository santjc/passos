package mapping

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type Mapping struct {
	ID                uuid.UUID            `json:"id" db:"id"`
	ResellerReference string               `json:"resellerReference" db:"reseller_reference"`
	ResellerStatus    types.ResellerStatus `json:"resellerStatus" db:"reseller_status"`
	Title             string               `json:"title" db:"title"`
	URL               string               `json:"url" db:"url"`
	WebhookURL        *string              `json:"webhookUrl,omitempty" db:"webhook_url"`
	OptionRequired    bool                 `json:"optionRequired" db:"option_required"`
	UnitRequired      bool                 `json:"unitRequired" db:"unit_required"`
	ProductID         *uuid.UUID           `json:"productId,omitempty" db:"product_id"`
	OptionID          *uuid.UUID           `json:"optionId,omitempty" db:"option_id"`
	UnitID            *uuid.UUID           `json:"unitId,omitempty" db:"unit_id"`
	Connected         bool                 `json:"connected" db:"connected"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
