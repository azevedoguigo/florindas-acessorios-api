package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProductRepository interface {
	Create(cartProduct *domain.CartProduct) error
	UpdateQuantity(id uuid.UUID, quantity uint64) error
	Delete(id uuid.UUID) error
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

func (r cartProductRepository) UpdateQuantity(id uuid.UUID, quantity uint64) error {
	cartProduct := &domain.CartProduct{ID: id}

	return r.db.Model(&cartProduct).Update("quantity", quantity).Error
}

func (r cartProductRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.CartProduct{}, id).Error
}
