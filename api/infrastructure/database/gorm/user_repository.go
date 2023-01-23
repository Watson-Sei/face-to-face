package gorm

import (
	"github.com/Watson-Sei/face-to-face/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	List() ([]models.User, error)
	Get(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	Delete(id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// List returns all users
func (r *userRepo) List() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get returns a user
func (r *userRepo) Get(id string) (models.User, error) {
	var user models.User
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Create creates a user
func (r *userRepo) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Delete deletes a user
func (r *userRepo) Delete(id string) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
