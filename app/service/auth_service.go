package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/helper"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoginService(c *fiber.Ctx, dto request.LoginDto) (response.UserDto, error) {

	user, err := helper.FindUserByIdentifier(dto.Identifier)
	if err != nil {
		return response.UserDto{}, err
	}

	err = utils.CheckPassword(dto.Password, user.Password)
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

func RegisterService(c *fiber.Ctx, dto request.RegisterDto) (response.UserDto, error) {

	_, err := helper.FindUserByIdentifier(dto.Email)
	if err == nil {
		return response.UserDto{}, fmt.Errorf(constant.ErrEmailAlreadyRegistered)
	}

	user, err := helper.AddUser(dto)
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
