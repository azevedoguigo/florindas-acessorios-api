package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(client *domain.Client) error
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return clientRepository{db: db}
}

func (r clientRepository) Create(client *domain.Client) error {
	return r.db.Create(client).Error
}
