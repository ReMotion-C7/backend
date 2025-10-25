package service

import (
	"ReMotion-C7/app/controller/helper"
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginService(c *fiber.Ctx, loginDto request.LoginDto) (response.UserDto, error) {

	user, err := helper.FindUserByIdentifier(loginDto.Identifier)
	if err != nil {
		return response.UserDto{}, err
	}

	err = utils.CheckPassword(loginDto.Password, user.Password)
	if err != nil {
		return response.UserDto{}, err
	}

	userDto := response.UserDto{
		Id:     int(user.ID),
		Name:   user.Name,
		RoleId: int(user.RoleID),
	}

	return userDto, nil

}

func RegisterService(c *fiber.Ctx, registerDto request.RegisterDto) (response.UserDto, error) {

	user, err := helper.AddUser(registerDto)
	if err != nil {
		return response.UserDto{}, err
	}

	userDto := response.UserDto{
		Id:     int(user.ID),
		Name:   user.Name,
		RoleId: int(user.RoleID),
	}

	return userDto, nil

}
