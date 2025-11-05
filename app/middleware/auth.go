package middleware

import (
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {

	expectedRoleId := 2

	_, err := GetAuthorizedUser(c, expectedRoleId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return c.Next()

}

func FisioMiddlewareWithoutId(c *fiber.Ctx) error {

	expectedRoleId := 1

	_, err := GetAuthorizedUser(c, expectedRoleId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return c.Next()

}

func FisioMiddleware(c *fiber.Ctx) error {

	expectedRoleId := 1

	fisioId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	userId, err := GetAuthorizedUser(c, expectedRoleId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusUnauthorized, err.Error(), nil)
	}

	if fisioId != userId {
		return output.GetOutput(c, constant.StatusError, fiber.StatusUnauthorized, constant.ErrInvalidFisio, nil)
	}

	return c.Next()

}

func GetAuthorizedUser(c *fiber.Ctx, expectedRoleId int) (int, error) {

	claims, err := utils.ParseToken(c)
	if err != nil {
		return 0, fmt.Errorf(constant.ErrInvalidTokenError)
	}

	roleId := claims["roleId"].(float64)
	if int(roleId) != expectedRoleId {
		return 0, fmt.Errorf(constant.ErrPermissionDenied)
	}

	userId := claims["id"].(float64)

	return int(userId), nil

}
