package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{db: db}
}

func (r productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}
