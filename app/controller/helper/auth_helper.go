package helper

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/utils"
	"time"
)

func FindUserByIdentifier(value string) (model.User, error) {

	database := config.GetDatabase()
	var user model.User

	err := database.Where(`email = ? OR phone_number = ?`, value, value).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}

func AddUser(registerDto request.RegisterDto) (model.User, error) {

	database := config.GetDatabase()

	hashPassword, err := utils.HashPassword(registerDto.Password)
	if err != nil {
		return model.User{}, err
	}

	dob, err := time.Parse("2006-01-02", registerDto.DateOfBirth)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Email:       registerDto.Email,
		PhoneNumber: registerDto.PhoneNumber,
		Password:    hashPassword,
		DateOfBirth: dob,
		RoleID:      2,
		GenderID:    uint(registerDto.GenderId),
	}

	err = database.Create(user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}
