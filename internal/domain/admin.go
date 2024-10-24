package domain

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
