package shared

import (
	"time"

	"passos/internal/types"

	"github.com/google/uuid"
)

type ProductDeliveryFormat struct {
	ID             uuid.UUID            `json:"id" db:"id"`
	ProductID      uuid.UUID            `json:"productId" db:"product_id"`
	DeliveryFormat types.DeliveryFormat `json:"deliveryFormat" db:"delivery_format"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type ProductDeliveryMethod struct {
	ID             uuid.UUID            `json:"id" db:"id"`
	ProductID      uuid.UUID            `json:"productId" db:"product_id"`
	DeliveryMethod types.DeliveryMethod `json:"deliveryMethod" db:"delivery_method"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
