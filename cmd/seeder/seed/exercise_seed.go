package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"
	"math/rand/v2"

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

func SeedExercises(db *gorm.DB) {

	typeId1 := uint(1)
	typeId2 := uint(2)

	exercises := []model.Exercise{
		{
			Name:        "Lunges",
			Description: "Latihan kaki untuk memperkuat otot paha dan bokong dengan melangkah maju lalu menekuk lutut.",
			Muscle:      "Otot Paha",
			Image:       "https://via.placeholder.com/150",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Lunges%20Landscape%20Image.png",
			TypeID:      &typeId1,
		},
		{
			Name:        "Squat",
			Description: "Latihan dasar kaki yang melatih otot paha, bokong, dan punggung bawah.",
			Muscle:      "Otot Paha",
			Image:       "https://via.placeholder.com/150",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Squat%20Landscape%20Image.png",
			TypeID:      &typeId1,
		},
		{
			Name:        "One Leg Balance",
			Description: "Latihan keseimbangan dengan berdiri di satu kaki untuk melatih stabilitas dan otot kaki.",
			Muscle:      "Otot Kaki",
			Image:       "https://via.placeholder.com/150",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/One%20Leg%20Balance%20Image.png",
			TypeID:      &typeId2,
		},
	}

	err := db.Create(&exercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}

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

		exerciseTotal := rand.IntN(2) + 2

		for _ = range make([]struct{}, exerciseTotal) {

			exerciseId := uint(rand.IntN(len(exercises)) + 1)

			patientExercise := model.PatientExercise{
				Set:        (rand.IntN(3) + 1),
				RepOrTime:  (rand.IntN(3) + 1) * 10,
				PatientID:  patient.ID,
				ExerciseID: exerciseId,
			}

			patientExercises = append(patientExercises, patientExercise)

		}

	}

	err = db.Create(&patientExercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
