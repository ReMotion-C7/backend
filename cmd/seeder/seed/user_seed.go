package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"log"
	"math/rand/v2"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {

	var roles []model.Role
	var genders []model.Gender

	err := db.Find(&roles).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	err = db.Find(&genders).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	users := []model.User{}

	for i := 0; i < 6; i++ {

		hashPassword, err := utils.HashPassword("password123")
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		dateOfBirthStr := faker.Date()
		date, err := time.Parse("2006-01-02", dateOfBirthStr)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		roleID := uint(1)
		if i >= 2 {
			roleID = 2
		}

		user := model.User{
			Name:        faker.Name(),
			Email:       faker.Email(),
			Password:    hashPassword,
			PhoneNumber: faker.Phonenumber(),
			DateOfBirth: date,
			RoleID:      &roleID,
			GenderID:    utils.PointNumber(rand.IntN(2) + 1),
		}

		users = append(users, user)
	}

	err = db.Create(&users).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
