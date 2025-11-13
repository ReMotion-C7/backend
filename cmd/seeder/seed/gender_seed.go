package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedGenders(db *gorm.DB) {

	genders := []model.Gender{
		{Name: "Laki-laki"},
		{Name: "Wanita"},
	}

	err := db.Create(&genders).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
