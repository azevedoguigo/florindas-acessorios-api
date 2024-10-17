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
	UF          string    `json:"uf"`
	CEP         string    `json:"cep"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
