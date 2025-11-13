package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"math/rand/v2"
	"time"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func SeedPatients(db *gorm.DB) {

	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	patients := []model.Patient{}

	for i := 0; i < 4; i++ {

		therapyStartDate := faker.Date()
		date, err := time.Parse("2006-01-02", therapyStartDate)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		var fisioId uint

		if i <= 1 {
			fisioId = 1
		} else {
			fisioId = 2
		}

		patient := model.Patient{
			Phase:            (rand.IntN(2) + 1),
			TherapyStartDate: date,
			UserID:           uint(i + 3),
			FisiotherapyID:   fisioId,
		}

		patients = append(patients, patient)

	}

	err = db.Create(&patients).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
