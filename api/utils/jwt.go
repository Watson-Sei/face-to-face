package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string `json:"user_id"`
	Roles  string `json:"roles"`
	jwt.StandardClaims
}

// JWTトークンを発行する関数
func GenerateJWT(user_id string, role string) (string, error) {

	// トークンを署名するための秘密鍵
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := Claims{
		UserID: user_id,
		Roles:  role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "face-to-face",
			Subject:   user_id,
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
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

// JWTをデコードする関数
func DecodeJWT(token string) (Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// 署名を検証する
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return Claims{}, err
	}
	return *claims, nil
}

// JWTの検証をする
func VerifyJWT(claims Claims, role string) error {
	// トークンの有効期限を検証する
	exp, ok := claims.StandardClaims.ExpiresAt, true
	if ok && time.Unix(int64(exp), 0).Before(time.Now()) {
		return fmt.Errorf("token is expired")
	}

	// トークンのissuerを検証する
	iss, ok := claims.StandardClaims.Issuer, true
	if ok && iss != "face-to-face" {
		return fmt.Errorf("invalid iss")
	}

	// トークンのsubjectを検証する
	sub, ok := claims.StandardClaims.Subject, true
	if ok && sub != claims.UserID {
		return fmt.Errorf("invalid sub")
	}

	// トークンのroleを検証する
	roles, ok := claims.Roles, true
	if ok && roles != role {
		return fmt.Errorf("invalid role")
	}

	return nil
}
