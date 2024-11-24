package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) error
	FindByUserID(userID uuid.UUID) (*domain.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepository{db: db}
}

func (r adminRepository) Create(admin *domain.Admin) error {
	return r.db.Create(admin).Error
}

func (r adminRepository) FindByUserID(userID uuid.UUID) (*domain.Admin, error) {
	admin := domain.Admin{UserID: userID}

	if err := r.db.First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}
