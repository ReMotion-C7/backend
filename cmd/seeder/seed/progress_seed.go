package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedProgress(db *gorm.DB) {

    var progresses []model.Progress

    dates := []time.Time{
        time.Date(2025, 11, 17, 0, 0, 0, 0, time.UTC),
        time.Date(2025, 11, 15, 0, 0, 0, 0, time.UTC),
        time.Date(2025, 11, 13, 0, 0, 0, 0, time.UTC),
    }

    for patientID := 1; patientID <= 4; patientID++ {
        for j := 0; j < 3; j++ {
            progress := model.Progress{
                Date:      dates[j],
                PatientID: uint(patientID),
            }

            progresses = append(progresses, progress)
        }
    }

    if err := db.Create(&progresses).Error; err != nil {
        log.Println(constant.ErrSeedingDatabase)
    }

}
