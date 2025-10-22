package product

import (
	"passos/internal/types"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                     uuid.UUID              `json:"id" db:"id"`
	InternalName           string                 `json:"internalName" db:"internal_name"`
	Reference              *string                `json:"reference,omitempty" db:"reference"`
	Locale                 string                 `json:"locale" db:"locale"`
	TimeZone               string                 `json:"timeZone" db:"time_zone"`
	AllowFreesale          bool                   `json:"allowFreesale" db:"allow_freesale"`
	InstantConfirmation    bool                   `json:"instantConfirmation" db:"instant_confirmation"`
	InstantDelivery        bool                   `json:"instantDelivery" db:"instant_delivery"`
	AvailabilityRequired   bool                   `json:"availabilityRequired" db:"availability_required"`
	AvailabilityType       types.AvailabilityType `json:"availabilityType" db:"availability_type"`
	RedemptionMethod       types.RedemptionMethod `json:"redemptionMethod" db:"redemption_method"`
	FreesaleDurationAmount int                    `json:"freesaleDurationAmount" db:"freesale_duration_amount"`
	FreesaleDurationUnit   string                 `json:"freesaleDurationUnit" db:"freesale_duration_unit"`

	ProductContentID   uuid.UUID  `json:"productContentId" db:"product_content_id"`
	ProductPricingID   uuid.UUID  `json:"productPricingId" db:"product_pricing_id"`
	ProductPackageID   uuid.UUID  `json:"productPackageId" db:"product_package_id"`
	ProductQuestionsID *uuid.UUID `json:"productQuestionsId,omitempty" db:"product_questions_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
