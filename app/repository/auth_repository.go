package repository

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

func AddUser(dto request.RegisterDto) (model.User, error) {

	database := config.GetDatabase()

	hashPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return model.User{}, err
	}

	dob, err := time.Parse("2006-01-02", dto.DateOfBirth)
	if err != nil {
		return model.User{}, err
	}

	roleId := uint(2)
	genderId := uint(dto.GenderId)

	user := model.User{
		Name:        dto.Name,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		Password:    hashPassword,
		DateOfBirth: dob,
		RoleID:      &roleId,
		GenderID:    &genderId,
	}

	err = database.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}
