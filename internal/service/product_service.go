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
	GetProducts() ([]domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	UpdateProduct(id string, updateProductDTO *contract.UpdateProductDTO) error
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
		Price:       &newProductDTO.Price,
		Quantity:    &newProductDTO.Quantity,
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

func (s productService) GetProducts() ([]domain.Product, error) {
	products, err := s.productRepo.Get()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s productService) GetProductByID(id string) (*domain.Product, error) {
	productUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	product, err := s.productRepo.GetByID(productUUID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s productService) UpdateProduct(id string, updateProductDTO *contract.UpdateProductDTO) error {
	productUUID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid product ID")
	}

	if err := pkg.ValidateStruct(updateProductDTO); err != nil {
		return err
	}

	if err := s.productRepo.Update(productUUID, updateProductDTO); err != nil {
		return err
	}

	return nil
}
