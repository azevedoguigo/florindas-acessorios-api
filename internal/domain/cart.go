package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID uuid.UUID
}
