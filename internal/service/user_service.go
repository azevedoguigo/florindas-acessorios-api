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

type UserService interface {
	CreateUser(newUserDTO *contract.NewUserDTO) error
	Login(loginDTO *contract.LoginDTO) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(newUserDTO *contract.NewUserDTO) error {
	if err := pkg.ValidateStruct(newUserDTO); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		ID:       uuid.New(),
		Name:     newUserDTO.Name,
		Email:    newUserDTO.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(loginDTO *contract.LoginDTO) (string, error) {
	user, err := s.userRepo.FindByEmail(loginDTO.Email)
	if err != nil {
		return "", errors.New("email not registred")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	return pkg.GenerateJWT(user.ID)
}
