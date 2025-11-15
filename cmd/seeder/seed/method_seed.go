package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedMethods(db *gorm.DB) {

	methods := []model.Method{
		{Name: "Repetisi"},
		{Name: "Waktu"},
	}

	err := db.Create(&methods).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
