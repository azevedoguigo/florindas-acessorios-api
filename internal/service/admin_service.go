package service

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateAdmin(newAdminDTO *contract.NewAdminDTO) error
}

type adminService struct {
	userRepo  repository.UserRepository
	adminRepo repository.AdminRepository
}

func NewAdminService(userRepo repository.UserRepository, adminRepo repository.AdminRepository) AdminService {
	return adminService{
		userRepo:  userRepo,
		adminRepo: adminRepo,
	}
}

func (s adminService) CreateAdmin(newAdminDTO *contract.NewAdminDTO) error {
	if err := pkg.ValidateStruct(newAdminDTO); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdminDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID:       uuid.New(),
		Name:     newAdminDTO.Name,
		Email:    newAdminDTO.Email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	admin := &domain.Admin{
		ID:     uuid.New(),
		UserID: user.ID,
	}

	err = s.adminRepo.Create(admin)
	if err != nil {
		return err
	}

	return nil
}
