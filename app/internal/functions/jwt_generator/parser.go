package generator

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func Parser(jwtEncr string) (login, password string) {
	var userClaim UserClaim
	secretKey, err := LoadKey()
	if err != nil {
		log.Fatal(err)
	}

	token, err := jwt.ParseWithClaims(jwtEncr, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if !token.Valid {
		log.Fatal("invalid token")
	}

	login, password = userClaim.Login, userClaim.Password
	return
}
