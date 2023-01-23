package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user_id string, role string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id": user_id,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJWT(tokenString string, role string) (interface{}, error) {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	exp, ok := claims.Claims.(jwt.MapClaims)["exp"].(float64), true
	if ok && time.Unix(int64(exp), 0).Before(time.Now()) {
		return nil, fmt.Errorf("token is expired")
	}

	roles, ok := claims.Claims.(jwt.MapClaims)["role"].(string), true
	if ok && roles != role {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
