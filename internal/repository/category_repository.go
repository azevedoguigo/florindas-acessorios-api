package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	Get() ([]domain.Category, error)
	GetByID(id uuid.UUID) (*domain.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepository{db: db}
}

func (r categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r categoryRepository) Get() ([]domain.Category, error) {
	var categories []domain.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r categoryRepository) GetByID(id uuid.UUID) (*domain.Category, error) {
	var category domain.Category

	if err := r.db.Preload("Products.Images").Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
