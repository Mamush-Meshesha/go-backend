package services

import (
	"errors"
	"todo/crypto"
	"todo/models"
	"todo/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(email, password string) (*models.User, error){
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email: email,
		Password: hashedPassword,
	}
	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !crypto.CheckPasswordHash(password, user.Password){
		return "", errors.New("invalid credentials")
	}

	return crypto.GenerateToken(user.ID)
}