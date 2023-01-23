package auth

import (
	"errors"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	user_id := "1"
	role := "guest"

	// run method
	tokenString, err := GenerateJWT(user_id, role)

	// assert
	assert.NoError(t, err)

	claims := &jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	assert.NoError(t, err)
	assert.Equal(t, user_id, (*claims)["user_id"].(string))
}

func TestValidateJWT(t *testing.T) {
	user_id := "1"
	role := "guest"

	token, _ := GenerateJWT(user_id, role)

	// run method
	claims, err := ValidateJWT(token, role)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, user_id, claims.(*jwt.Token).Claims.(jwt.MapClaims)["user_id"])
	assert.Equal(t, role, claims.(*jwt.Token).Claims.(jwt.MapClaims)["role"])
}
