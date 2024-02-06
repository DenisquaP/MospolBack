package generator

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtGenerator(login, password, secretKey string) (string, error) {
	bSecretKey := []byte(secretKey)
	payload := jwt.MapClaims{
		"login":    login,
		"password": password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(bSecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
