package utils

import (
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user response.UserDto) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.Id,
		"roleId": user.RoleId,
		"exp":    expirationTime,
	})

	secretKey := config.LoadEnvConfig("SECRET_KEY")

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
