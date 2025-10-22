package product

import (
	"passos/internal/types"
	"time"

	"github.com/google/uuid"
)

type ProductQuestions struct {
	ID uuid.UUID `json:"id" db:"id"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type Question struct {
	ID                 uuid.UUID       `json:"id" db:"id"`
	ProductQuestionsID uuid.UUID       `json:"productQuestionsId" db:"product_questions_id"`
	Title              string          `json:"title" db:"title"`
	ShortDescription   string          `json:"shortDescription" db:"short_description"`
	Required           bool            `json:"required" db:"required"`
	InputType          types.InputType `json:"inputType" db:"input_type"`
	CoverImageUrl      string          `json:"coverImageUrl" db:"cover_image_url"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type QuestionSelectOption struct {
	ID         uuid.UUID `json:"id" db:"id"`
	QuestionID uuid.UUID `json:"questionId" db:"question_id"`
	Label      string    `json:"label" db:"label"`
	Value      string    `json:"value" db:"value"`
	SortOrder  int       `json:"sortOrder" db:"sort_order"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type QuestionAnswer struct {
	ID         uuid.UUID `json:"id" db:"id"`
	QuestionID uuid.UUID `json:"questionId" db:"question_id"`
	BookingID  uuid.UUID `json:"bookingId" db:"booking_id"`
	Value      string    `json:"value" db:"value"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
