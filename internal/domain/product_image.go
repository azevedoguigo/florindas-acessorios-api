package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	URL       string
	ProductID uuid.UUID
}
