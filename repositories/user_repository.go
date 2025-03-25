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

func (r *UserRepository) DeleteUser(userID uint) error {
    // Start a transaction
    tx := r.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // First permanently delete all user's todos
    if err := tx.Unscoped().Where("user_id = ?", userID).Delete(&models.Todo{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Then permanently delete the user
    if err := tx.Unscoped().Delete(&models.User{}, userID).Error; err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error
    return users, err
}