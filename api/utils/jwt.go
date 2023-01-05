package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTトークンを発行する関数
func GenerateJWT(username string, role string) (string, error) {

	// トークンを署名するための秘密鍵
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	// トークンの有効期限
	expirationTime := time.Now().Add(1 * time.Hour)

	// トークンに載せるクレームを定義
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    role,
		Subject:   username,
	}

	// トークンを署名する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// トークンを検証する関数
func ValidateJWT(tokenString string) (bool, error) {

	// トークンを署名するための秘密鍵
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	// トークンを解析する
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	// トークンが有効かどうかを判定する
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims["iss"], claims["sub"])
		return true, nil
	} else {
		return false, nil
	}
}
