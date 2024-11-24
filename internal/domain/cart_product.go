package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	CartID   uuid.UUID
	Product  Product
	Quantity uint64
}
