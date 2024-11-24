package service

import (
	"errors"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
)

type CartProductService interface {
	CreateCartProduct(newCartProductDTO *contract.NewCartProductDTO) error
	UpdateCartProductQuantity(
		updateCartProductQuantityDTO *contract.UpdateCartProductQuantityDTO,
	) error
	DeleteCartProduct(id string) error
}

type cartProductService struct {
	cartProductRepo repository.CartProductRepository
	productRepo     repository.ProductRepository
}

func NewCartProductService(
	cartProductRepo repository.CartProductRepository,
	productRepo repository.ProductRepository,
) CartProductService {
	return &cartProductService{
		cartProductRepo: cartProductRepo,
		productRepo:     productRepo,
	}
}

func (s *cartProductService) CreateCartProduct(newCartProductDTO *contract.NewCartProductDTO) error {
	if err := pkg.ValidateStruct(newCartProductDTO); err != nil {
		return err
	}

	cartUUID, err := uuid.Parse(newCartProductDTO.CartID)
	if err != nil {
		return errors.New("invalid cart ID")
	}

	productUUID, err := uuid.Parse(newCartProductDTO.ProductID)
	if err != nil {
		return errors.New("invalid product ID")
	}

	product, err := s.productRepo.GetByID(productUUID)
	if err != nil {
		return err
	}

	cartProduct := &domain.CartProduct{
		ID:       uuid.New(),
		CartID:   cartUUID,
		Product:  *product,
		Quantity: 1,
	}

	if err := s.cartProductRepo.Create(cartProduct); err != nil {
		return err
	}

	return nil
}

func (s *cartProductService) UpdateCartProductQuantity(
	updateCartProductQuantityDTO *contract.UpdateCartProductQuantityDTO,
) error {
	if err := pkg.ValidateStruct(updateCartProductQuantityDTO); err != nil {
		return err
	}

	cartProductUUID, err := uuid.Parse(updateCartProductQuantityDTO.CartProductID)
	if err != nil {
		return errors.New("invalid product ID")
	}

	err = s.cartProductRepo.UpdateQuantity(
		cartProductUUID,
		updateCartProductQuantityDTO.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *cartProductService) DeleteCartProduct(id string) error {
	cartProductUUID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid product ID")
	}

	if err := s.cartProductRepo.Delete(cartProductUUID); err != nil {
		return err
	}

	return nil
}
