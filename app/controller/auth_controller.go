package controller

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var loginDto request.LoginDto
	err := c.BodyParser(&loginDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(loginDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	user, err := service.LoginService(c, loginDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, constant.ErrInvalidCredentials, nil)
	}

	accessToken, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessLogin, response.AuthDataDto{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   "24h",
		User:        user,
	})

}

func Register(c *fiber.Ctx) error {

	var registerDto request.RegisterDto
	err := c.BodyParser(&registerDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(registerDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	user, err := service.RegisterService(c, registerDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	accessToken, err := utils.GenerateJWT(user)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessRegister, response.AuthDataDto{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   "24h",
		User:        user,
	})

}
