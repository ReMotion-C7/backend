package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedTypes(db *gorm.DB) {

	types := []model.Type{
		{Name: "Repetisi"},
		{Name: "Waktu"},
	}

	err := db.Create(&types).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
