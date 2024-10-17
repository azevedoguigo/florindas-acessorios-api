package domain

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	User        User      `json:"-"`
	CPF         string    `json:"cpf"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
