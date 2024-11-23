package service

import (
	"errors"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/domain"
	"github.com/azevedoguigo/florindas-acessorios-api/internal/repository"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ClientService interface {
	CreateClient(newClientDTO *contract.NewClientDTO) error
	GetClientByID(id string) (*contract.GetClientResponseDTO, error)
	GetClientByUserID(id string) (*contract.GetClientResponseDTO, error)
}

type clientService struct {
	userRepo   repository.UserRepository
	clientRepo repository.ClientRepository
	cartRepo   repository.CartRepository
}

func NewClientService(
	userRepo repository.UserRepository,
	clientRepo repository.ClientRepository,
	cartRepo repository.CartRepository,
) ClientService {
	return clientService{
		userRepo:   userRepo,
		clientRepo: clientRepo,
		cartRepo:   cartRepo,
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
		Role:     "client",
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	client := &domain.Client{
		ID:          uuid.New(),
		UserID:      user.ID,
		CPF:         newClientDTO.CPF,
		UF:          newClientDTO.UF,
		CEP:         newClientDTO.CEP,
		City:        newClientDTO.City,
		Address:     newClientDTO.Address,
		PhoneNumber: newClientDTO.PhoneNumber,
	}

	err = s.clientRepo.Create(client)
	if err != nil {
		return err
	}

	cart := &domain.Cart{
		ID:     uuid.New(),
		UserID: user.ID,
	}

	err = s.cartRepo.Create(cart)
	if err != nil {
		return err
	}

	return nil
}

func (s clientService) GetClientByID(id string) (*contract.GetClientResponseDTO, error) {
	clientUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid client ID")
	}

	client, err := s.clientRepo.FindByID(clientUUID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(client.UserID)
	if err != nil {
		return nil, err
	}

	clientResponseDTO := &contract.GetClientResponseDTO{
		ID:          client.ID,
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		CPF:         client.CPF,
		UF:          client.UF,
		CEP:         client.CEP,
		City:        client.City,
		Address:     client.Address,
		PhoneNumber: client.PhoneNumber,
	}

	return clientResponseDTO, nil
}

func (s clientService) GetClientByUserID(id string) (*contract.GetClientResponseDTO, error) {
	userUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	client, err := s.clientRepo.FindByUserID(userUUID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(client.UserID)
	if err != nil {
		return nil, err
	}

	clientResponseDTO := &contract.GetClientResponseDTO{
		ID:          client.ID,
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		CPF:         client.CPF,
		UF:          client.UF,
		CEP:         client.CEP,
		City:        client.City,
		Address:     client.Address,
		PhoneNumber: client.PhoneNumber,
	}

	return clientResponseDTO, nil
}
