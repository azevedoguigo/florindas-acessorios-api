package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type CartProductRepository interface {
	Create(cartProduct *domain.CartProduct) error
}

type cartProductRepository struct {
	db *gorm.DB
}

func NewCartProductRepository(db *gorm.DB) CartProductRepository {
	return cartProductRepository{db: db}
}

func (r cartProductRepository) Create(cartProduct *domain.CartProduct) error {
	return r.db.Create(cartProduct).Error
}
