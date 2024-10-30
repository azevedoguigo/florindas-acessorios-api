package repository

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) error
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
