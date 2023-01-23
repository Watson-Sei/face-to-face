package usecases

import (
	"log"
	"strconv"

	"github.com/Watson-Sei/face-to-face/domain/models"
	"github.com/Watson-Sei/face-to-face/domain/repositories"
	"github.com/Watson-Sei/face-to-face/infrastructure/auth"
	"github.com/Watson-Sei/face-to-face/infrastructure/auth/oauth"
)

type AuthUsecase interface {
	Login(code string) (string, error)
}

type authUsecase struct {
	UserRepository repositories.UserRepository
}

func NewAuthUsecase(ur repositories.UserRepository) AuthUsecase {
	return &authUsecase{UserRepository: ur}
}

func (a *authUsecase) Login(code string) (string, error) {
	accessToken, err := oauth.RequestToken(code)
	if err != nil {
		return "", err
	}

	userInfo, err := oauth.RequestUserInfo(accessToken)
	if err != nil {
		return "", err
	}

	user, err := a.UserRepository.Create(models.User{
		Name:  userInfo.Name,
		Email: userInfo.Email,
	})
	if err != nil {
		return "", err
	}

	id := strconv.Itoa(int(user.ID))
	log.Println(id)
	token, err := auth.GenerateJWT(id, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
