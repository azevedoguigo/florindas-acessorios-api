package service

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(newCategoryDTO *contract.NewCategoryDTO) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return categoryService{repo: repo}
}

func (s categoryService) CreateCategory(newCategoryDTO *contract.NewCategoryDTO) error {
	if err := pkg.ValidateStruct(newCategoryDTO); err != nil {
		return err
	}

	category := &domain.Category{
		ID:   uuid.New(),
		Name: newCategoryDTO.Name,
	}

	err := s.repo.Create(category)
	if err != nil {
		return err
	}

	return nil
}
