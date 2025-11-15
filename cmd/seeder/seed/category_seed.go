package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {

	categories := []model.Category{
		{Name: "Keseimbangan"},
		{Name: "Penguatan"},
	}

	err := db.Create(&categories).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
