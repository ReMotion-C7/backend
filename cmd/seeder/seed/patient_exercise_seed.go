package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
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

		if patient.ID == 1 || patient.ID == 2 {

			for i := 0; i < 3; i++ {

				patientExercise := model.PatientExercise{
					MethodID:   utils.PointNumber(rand.IntN(2) + 1),
					Set:        (rand.IntN(3) + 1),
					RepOrTime:  (rand.IntN(3) + 1) * 10,
					PatientID:  patient.ID,
					ExerciseID: uint(i + 1),
				}

				patientExercises = append(patientExercises, patientExercise)

			}

		} else if patient.ID == 3 || patient.ID == 4 {

			for i := 0; i < 2; i++ {

				patientExercise := model.PatientExercise{
					MethodID:   utils.PointNumber(rand.IntN(2) + 1),
					Set:        (rand.IntN(3) + 1),
					RepOrTime:  (rand.IntN(3) + 1) * 10,
					PatientID:  patient.ID,
					ExerciseID: uint(i + 4),
				}

				patientExercises = append(patientExercises, patientExercise)

			}

		}

	}

	err = db.Create(&patientExercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
