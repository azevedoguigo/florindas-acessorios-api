package repository

import (
	"fmt"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(client *domain.Client) error
	FindByID(id uuid.UUID) (*domain.Client, error)
	FindByUserID(id uuid.UUID) (*domain.Client, error)
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

func (r clientRepository) FindByID(id uuid.UUID) (*domain.Client, error) {
	var client domain.Client

	if err := r.db.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r clientRepository) FindByUserID(userID uuid.UUID) (*domain.Client, error) {
	var client domain.Client

	fmt.Println(userID)

	if err := r.db.Where("user_id = ?", userID).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}
