package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(addres *domain.Address) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return addressRepository{db: db}
}

func (r addressRepository) Create(addres *domain.Address) error {
	return r.db.Create(addres).Error
}
