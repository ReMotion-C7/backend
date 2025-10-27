package middleware

import (
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {

	expectedRoleId := 2

	err := GetAuthorizedUser(c, expectedRoleId)
	if err != nil {
		return err
	}

	return c.Next()

}

func FisioMiddleware(c *fiber.Ctx) error {

	expectedRoleId := 1

	err := GetAuthorizedUser(c, expectedRoleId)
	if err != nil {
		return err
	}

	return c.Next()

}

func GetAuthorizedUser(c *fiber.Ctx, expectedRoleId int) error {

	claims, err := utils.ParseToken(c)
	if err != nil {
		return err
	}

	roleId, ok := claims["roleId"].(int)
	if roleId != expectedRoleId || !ok {
		return fmt.Errorf(constant.ErrPermissionDenied)
	}

	return nil

}
