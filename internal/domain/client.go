package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID      uuid.UUID
	CPF         string
	UF          string
	CEP         string
	City        string
	Address     string
	PhoneNumber string
}
