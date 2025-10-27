package utils

import (
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/config"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

func ParseToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	token := c.Get("Authorization")

	tokenString := strings.TrimPrefix(token, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)
	fmt.Println(token)

	claims := jwt.MapClaims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, GetSecretKey)
	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	fmt.Println(claims)

	return claims, nil

}

func GetSecretKey(token *jwt.Token) (interface{}, error) {

	secretKey := config.LoadEnvConfig("SECRET_KEY")

	return []byte(secretKey), nil

}
