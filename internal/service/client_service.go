package service

import (
	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ClientService interface {
	CreateClient(newClientDTO *contract.NewClientDTO) error
}

type clientService struct {
	userRepo   repository.UserRepository
	clientRepo repository.ClientRepository
}

func NewClientService(userRepo repository.UserRepository, clientRepo repository.ClientRepository) ClientService {
	return clientService{
		userRepo:   userRepo,
		clientRepo: clientRepo,
	}
}

func (s clientService) CreateClient(newClientDTO *contract.NewClientDTO) error {
	if err := pkg.ValidateStruct(newClientDTO); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newClientDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID:       uuid.New(),
		Name:     newClientDTO.Name,
		Email:    newClientDTO.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	client := &domain.Client{
		ID:          uuid.New(),
		UserID:      user.ID,
		CPF:         newClientDTO.CPF,
		Address:     newClientDTO.Address,
		PhoneNumber: newClientDTO.PhoneNumber,
		Role:        "client",
	}

	err = s.clientRepo.Create(client)
	if err != nil {
		return err
	}

	return nil
}
