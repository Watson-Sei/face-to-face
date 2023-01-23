package usecases

import (
	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/Watson-Sei/face-to-face/domain/repositories"
)

type UserUsecase interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
}

type userUsecase struct {
	UserRepository repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: ur}
}

func (u *userUsecase) GetUsers() ([]models.User, error) {
	return u.UserRepository.List()
}

func (u *userUsecase) GetUser(id string) (models.User, error) {
	return u.UserRepository.Get(id)
}

func (u *userUsecase) CreateUser(user models.User) (models.User, error) {
	return u.UserRepository.Create(user)
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.UserRepository.Delete(id)
}
