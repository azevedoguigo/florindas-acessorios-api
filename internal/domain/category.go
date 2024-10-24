package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
