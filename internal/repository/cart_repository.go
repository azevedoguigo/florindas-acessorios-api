package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *domain.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return cartRepository{db: db}
}

func (r cartRepository) Create(cart *domain.Cart) error {
	return r.db.Create(cart).Error
}
