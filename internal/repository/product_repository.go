package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	Get() ([]domain.Product, error)
	GetByID(id uuid.UUID) (*domain.Product, error)
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

func (r productRepository) Get() ([]domain.Product, error) {
	products := []domain.Product{}

	if err := r.db.Preload("Images").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r productRepository) GetByID(id uuid.UUID) (*domain.Product, error) {
	product := domain.Product{ID: id}

	if err := r.db.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
