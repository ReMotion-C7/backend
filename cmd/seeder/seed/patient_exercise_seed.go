package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"math/rand/v2"

	"gorm.io/gorm"
)

func SeedPatientExercises(db *gorm.DB) {

	var patients []model.Patient
	var exercises []model.Exercise

	err := db.Find(&patients).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	err = db.Find(&exercises).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	patientExercises := []model.PatientExercise{}

	for _, patient := range patients {

		for _, e := range exercises {

			patientExercise := model.PatientExercise{
				Set:        (rand.IntN(3) + 1),
				RepOrTime:  (rand.IntN(3) + 1) * 10,
				PatientID:  patient.ID,
				ExerciseID: e.ID,
			}

			patientExercises = append(patientExercises, patientExercise)

		}

	}

	err = db.Create(&patientExercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
