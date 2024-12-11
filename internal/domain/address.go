package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model

	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID     uuid.UUID
	CEP        string
	UF         string
	City       string
	Street     string
	Number     string
	Complement string
}
