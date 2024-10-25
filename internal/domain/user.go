package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name     string
	Email    string
	Password string
	Role     string
}
