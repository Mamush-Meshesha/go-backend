package repositories

import (
	"todo/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error){
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}