package product

import (
	"time"

	"github.com/google/uuid"
)

type ProductContent struct {
	ID                     uuid.UUID `json:"id" db:"id"`
	Title                  *string   `json:"title,omitempty" db:"title"`
	Country                *string   `json:"country,omitempty" db:"country"`
	Location               *string   `json:"location,omitempty" db:"location"`
	Subtitle               *string   `json:"subtitle,omitempty" db:"subtitle"`
	ShortDescription       *string   `json:"shortDescription,omitempty" db:"short_description"`
	Description            *string   `json:"description,omitempty" db:"description"`
	Highlights             *string   `json:"highlights,omitempty" db:"highlights"`
	Inclusions             *string   `json:"inclusions,omitempty" db:"inclusions"`
	Exclusions             *string   `json:"exclusions,omitempty" db:"exclusions"`
	BookingTerms           *string   `json:"bookingTerms,omitempty" db:"booking_terms"`
	RedemptionInstructions *string   `json:"redemptionInstructions,omitempty" db:"redemption_instructions"`
	CancellationPolicy     *string   `json:"cancellationPolicy,omitempty" db:"cancellation_policy"`
	Destination            *string   `json:"destination,omitempty" db:"destination"`
	Categories             *string   `json:"categories,omitempty" db:"categories"`
	Faqs                   *string   `json:"faqs,omitempty" db:"faqs"`
	CoverImageUrl          *string   `json:"coverImageUrl,omitempty" db:"cover_image_url"`
	BannerImageUrl         *string   `json:"bannerImageUrl,omitempty" db:"banner_image_url"`
	VideoUrl               *string   `json:"videoUrl,omitempty" db:"video_url"`
	GalleryImages          *string   `json:"galleryImages,omitempty" db:"gallery_images"`
	BannerImages           *string   `json:"bannerImages,omitempty" db:"banner_images"`
	PointToPoint           *bool     `json:"pointToPoint,omitempty" db:"point_to_point"`
	PrivacyTerms           *string   `json:"privacyTerms,omitempty" db:"privacy_terms"`
	Alert                  *string   `json:"alert,omitempty" db:"alert"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
