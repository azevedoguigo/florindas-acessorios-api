package service

import (
	"errors"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(newProductDTO *contract.NewProductDTO) error
}

type productService struct {
	productRepo      repository.ProductRepository
	productImageRepo repository.ProductImageRepository
}

func NewProductService(productRepo repository.ProductRepository, productImageRepo repository.ProductImageRepository) ProductService {
	return productService{
		productRepo:      productRepo,
		productImageRepo: productImageRepo,
	}
}

func (s productService) CreateProduct(newProductDTO *contract.NewProductDTO) error {
	if err := pkg.ValidateStruct(newProductDTO); err != nil {
		return err
	}

	categoryUUID, err := uuid.Parse(newProductDTO.CategoryID)
	if err != nil {
		return errors.New("invalid category ID")
	}

	product := &domain.Product{
		ID:          uuid.New(),
		Name:        newProductDTO.Name,
		Description: newProductDTO.Description,
		Price:       newProductDTO.Price,
		Quantity:    newProductDTO.Quantity,
		CategoryID:  categoryUUID,
	}

	if err := s.productRepo.Create(product); err != nil {
		return err
	}

	for i := 0; i < len(newProductDTO.Images); i++ {
		productImage := &domain.ProductImage{
			ID:        uuid.New(),
			URL:       newProductDTO.Images[i],
			ProductID: product.ID,
		}

		if err := s.productImageRepo.Create(productImage); err != nil {
			return err
		}
	}

	return nil
}
