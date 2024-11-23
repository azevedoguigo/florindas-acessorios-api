package service

import (
	"errors"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/google/uuid"
)

type CartService interface {
	GetCartByUserID(userID string) (*domain.Cart, error)
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return cartService{cartRepo: cartRepo}
}

func (s cartService) GetCartByUserID(userID string) (*domain.Cart, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("ivalid user ID")
	}

	cart, err := s.cartRepo.FindByUserID(userUUID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
