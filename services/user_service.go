package services

import (
	"todo/models"
	"todo/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) DeleteUser(userID uint) error {
	return s.userRepo.DeleteUser(userID)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.userRepo.GetAllUsers()
}