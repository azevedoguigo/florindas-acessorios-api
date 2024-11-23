package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	if err := r.db.First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}
