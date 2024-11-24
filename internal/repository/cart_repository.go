package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository interface {
	Create(cart *domain.Cart) error
	FindByUserID(userID uuid.UUID) (*domain.Cart, error)
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

func (r cartRepository) FindByUserID(userID uuid.UUID) (*domain.Cart, error) {
	cart := domain.Cart{UserID: userID}

	err := r.db.
		Preload("CartProducts.Product.Images").
		Preload(clause.Associations).
		First(&cart).
		Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
