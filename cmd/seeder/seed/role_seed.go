package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {

	roles := []model.Role{
		{Name: "Fisioterapi"},
		{Name: "Pasien"},
	}

	err := db.Create(&roles).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
