package domain

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID
	User        User `gorm:"foreignKey:UserID"`
	CPF         string
	UF          string
	CEP         string
	City        string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
