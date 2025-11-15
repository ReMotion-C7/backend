package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func SeedProgress(db *gorm.DB) {

	var progresses []model.Progress

	for i := 0; i < 4; i++ {

		date := faker.Date()
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		progress := model.Progress{
			Date:      parsedDate,
			PatientID: uint(i + 1),
		}

		progresses = append(progresses, progress)

	}

	err := db.Create(&progresses).Error
	if err != nil {
		log.Println(constant.ErrSeedingDatabase)
	}

}
