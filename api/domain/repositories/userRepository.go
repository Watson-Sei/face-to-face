package repositories

import (
	"errors"

	"github.com/Watson-Sei/face-to-face/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	List() ([]models.User, error)                 // 全てのユーザー情報を返す
	Get(id string) (models.User, error)           // 指定したIDのユーザー情報を返す
	Create(user models.User) (models.User, error) // ユーザー情報を作成する
	Delete(id string) error                       // 指定したIDのユーザー情報を削除する
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) List() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Get(id string) (models.User, error) {
	var user models.User
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepo) Create(userInfo models.User) (models.User, error) {
	if err := r.db.Where("email = ?", userInfo.Email).First(&userInfo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		userInfo.Role = "guest"
		if err := r.db.Create(&userInfo).Error; err != nil {
			return models.User{}, err
		}
		if err := r.db.Where("email = ?", userInfo.Email).First(&userInfo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, err
		}
	} else {
		if err := r.db.Model(&userInfo).Updates(&models.User{
			Name:  userInfo.Name,
			Email: userInfo.Email,
		}).Error; err != nil {
			return models.User{}, err
		}
	}
	return userInfo, nil
}

func (r *userRepo) Delete(id string) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
