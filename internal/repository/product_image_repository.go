package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type ProductImageRepository interface {
	Create(productImage *domain.ProductImage) error
}

type productImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return productImageRepository{db: db}
}

func (r productImageRepository) Create(productImage *domain.ProductImage) error {
	return r.db.Create(productImage).Error
}
