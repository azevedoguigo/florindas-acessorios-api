package service

import (
	"errors"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
)

type AddressService interface {
	CreateAddress(newAddress *contract.NewAddress) error
}

type addressService struct {
	addressRepo repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) AddressService {
	return addressService{addressRepo: addressRepo}
}

func (s addressService) CreateAddress(newAddress *contract.NewAddress) error {
	if err := pkg.ValidateStruct(newAddress); err != nil {
		return err
	}

	userUUID, err := uuid.Parse(newAddress.UserID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	address := &domain.Address{
		ID:         uuid.New(),
		UserID:     userUUID,
		CEP:        newAddress.CEP,
		UF:         newAddress.UF,
		City:       newAddress.City,
		Street:     newAddress.Street,
		Number:     newAddress.Number,
		Complement: newAddress.Complement,
	}

	if err := s.addressRepo.Create(address); err != nil {
		return err
	}

	return nil
}
