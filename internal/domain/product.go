package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Images      []ProductImage
	Name        string
	Description string
	Price       float64
	Quantity    uint64
	CategoryID  uuid.UUID
}
