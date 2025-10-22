package unit

import (
	"time"

	"github.com/google/uuid"
)

type UnitContent struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       *string   `json:"title,omitempty" db:"title"`
	TitlePlural *string   `json:"titlePlural,omitempty" db:"title_plural"`
	Subtitle    *string   `json:"subtitle,omitempty" db:"subtitle"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
