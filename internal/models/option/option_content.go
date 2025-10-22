package option

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type OptionContent struct {
	ID               uuid.UUID           `json:"id" db:"id"`
	Title            *string             `json:"title,omitempty" db:"title"`
	Subtitle         *string             `json:"subtitle,omitempty" db:"subtitle"`
	Language         *string             `json:"language,omitempty" db:"language"`
	ShortDescription *string             `json:"shortDescription,omitempty" db:"short_description"`
	Duration         *string             `json:"duration,omitempty" db:"duration"`
	DurationAmount   *string             `json:"durationAmount,omitempty" db:"duration_amount"`
	DurationUnit     *types.DurationUnit `json:"durationUnit,omitempty" db:"duration_unit"`
	CoverImageUrl    *string             `json:"coverImageUrl,omitempty" db:"cover_image_url"`

	FromPointID *uuid.UUID `json:"fromPointId,omitempty" db:"from_point_id"`
	ToPointID   *uuid.UUID `json:"toPointId,omitempty" db:"to_point_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type Itinerary struct {
	ID               uuid.UUID `json:"id" db:"id"`
	OptionContentID  uuid.UUID `json:"optionContentId" db:"option_content_id"`
	Name             string    `json:"name" db:"name"`
	Type             string    `json:"type" db:"type"`
	Description      string    `json:"description" db:"description"`
	Address          string    `json:"address" db:"address"`
	GooglePlaceId    string    `json:"googlePlaceId" db:"google_place_id"`
	Latitude         float64   `json:"latitude" db:"latitude"`
	Longitude        float64   `json:"longitude" db:"longitude"`
	TravelTime       string    `json:"travelTime" db:"travel_time"`
	TravelTimeAmount int       `json:"travelTimeAmount" db:"travel_time_amount"`
	TravelTimeUnit   string    `json:"travelTimeUnit" db:"travel_time_unit"`
	Duration         string    `json:"duration" db:"duration"`
	DurationAmount   int       `json:"durationAmount" db:"duration_amount"`
	DurationUnit     string    `json:"durationUnit" db:"duration_unit"`
	SortOrder        int       `json:"sortOrder" db:"sort_order"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Point struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	InternalName     string     `json:"internalName" db:"internal_name"`
	ShortDescription string     `json:"shortDescription" db:"short_description"`
	Title            string     `json:"title" db:"title"`
	PointGroupID     *uuid.UUID `json:"pointGroupId,omitempty" db:"point_group_id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type PointGroup struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Title            string    `json:"title" db:"title"`
	ShortDescription string    `json:"shortDescription" db:"short_description"`
	InternalName     string    `json:"internalName" db:"internal_name"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
