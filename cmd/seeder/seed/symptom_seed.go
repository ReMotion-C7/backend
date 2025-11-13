package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"math/rand/v2"

	"gorm.io/gorm"
)

func SeedSymptoms(db *gorm.DB) {

	var patients []model.Patient
	err := db.Find(&patients).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	symptomNameList := []string{
		"Nyeri lutut",
		"Bengkak di lutut",
		"Kaku saat membengkokkan lutut",
		"Sulit berjalan normal",
		"Rasa tidak stabil pada lutut",
	}

	symptoms := []model.Symptom{}

	for _, patient := range patients {

		symptomTotal := rand.IntN(3) + 1

		for _ = range make([]struct{}, symptomTotal) {

			symptomName := symptomNameList[rand.IntN(len(symptomNameList))]

			symptom := model.Symptom{
				Name:      symptomName,
				PatientID: patient.ID,
			}

			symptoms = append(symptoms, symptom)

		}

	}

	err = db.Create(&symptoms).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
