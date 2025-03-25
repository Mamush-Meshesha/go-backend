package services

import (
	"errors"
	"log"
	"todo/crypto"
	"todo/models"
	"todo/repositories"
)

type AuthService struct {
	userRepo     *repositories.UserRepository
	emailService *EmailService
}

func NewAuthService(userRepo *repositories.UserRepository, emailService *EmailService) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		emailService: emailService,
	}
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.emailService.SendWelcomeEmail(user.Email, user.Email); err != nil {
		log.Printf("Failed to send welcome email: %v", err)
	}
	

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if err := s.emailService.SendWelcomeEmail(email, email); err != nil {
		log.Printf("fialed to send welcome email: %v", err)
	}
	return createdUser, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !crypto.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return crypto.GenerateToken(user.ID)
}
