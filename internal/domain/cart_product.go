package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CartID    uuid.UUID `gorm:"not null;constraint:OnDelete:CASCADE;"`
	ProductID uuid.UUID `gorm:"not null;constraint:OnDelete:CASCADE;"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	Quantity  uint64
}
