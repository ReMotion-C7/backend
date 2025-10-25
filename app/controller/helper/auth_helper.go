package helper

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
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
