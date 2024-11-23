package domain

import "github.com/google/uuid"

type Cart struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID uuid.UUID
}
