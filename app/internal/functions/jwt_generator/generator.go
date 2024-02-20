package generator

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Login    string
	Password string
}

func JwtGenerator(login, password string) (string, error) {
	secretKey, err := LoadKey()
	if err != nil {
		log.Fatal(err)
	}

	bSecretKey := []byte(secretKey)
	claims := UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		Login:            login,
		Password:         password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(bSecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
