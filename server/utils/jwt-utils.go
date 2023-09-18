package utils

import (
	"errors"
	"os"
	"start/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	models.User
	jwt.StandardClaims
}

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwt(user models.User) (string, error) {

	claims := &CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "authservice",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signed, nil
}

func CheckJwt(token string) (CustomClaims, error) {

	customClaims := &CustomClaims{}

	jwtToken, err := jwt.ParseWithClaims(token, customClaims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error parsing jwt token")
		}

		return secret, nil
	})

	if err != nil || !jwtToken.Valid {
		return CustomClaims{}, err
	}

	if claims, ok := jwtToken.Claims.(*CustomClaims); ok {
		return *claims, nil
	} else {
		return CustomClaims{}, errors.New("error parsing jwt token")
	}
}
