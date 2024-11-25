package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	Get() ([]domain.Product, error)
	GetByID(id uuid.UUID) (*domain.Product, error)
	GetMostRecent() ([]domain.Product, error)
	Update(id uuid.UUID, product *contract.UpdateProductDTO) error
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

	if err := r.db.Preload("Images").First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r productRepository) GetMostRecent() ([]domain.Product, error) {
	products := []domain.Product{}

	err := r.db.Order("created_at DESC").Limit(10).Preload("Images").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r productRepository) Update(id uuid.UUID, dto *contract.UpdateProductDTO) error {
	var product domain.Product

	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return err
	}

	if dto.Name != "" {
		product.Name = dto.Name
	}

	if dto.Description != "" {
		product.Description = dto.Description
	}

	if dto.Price != nil {
		product.Price = dto.Price
	}

	if dto.Quantity != nil {
		product.Quantity = dto.Quantity
	}

	if err := r.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
