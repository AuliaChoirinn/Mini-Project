package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"lab-loan/constant"
	"time"
)

func CreateTokenAdmin(username string, password string, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT_ADMIN))

}

func CreateTokenUser(username string, password string, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT_USER))

}
