package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedPhases(db *gorm.DB) {

	phases := []model.Phase{
		{Name: "Fase 1 (Post-Op)"},
		{Name: "Fase 2 (Post-Op)"},
		{Name: "Fase 3 (Post-Op)"},
		{Name: "Fase 4 (Post-Op)"},
		{Name: "Pre-Op"},
		{Name: "Non-Op"},
	}

	err := db.Create(&phases).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}
}
