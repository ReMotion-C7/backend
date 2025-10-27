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
			Name:        "Straight Leg Raise",
			Description: "Berbaring telentang dan angkat kaki lurus setinggi 30 cm, lalu turunkan perlahan.",
			Muscle:      "Quadriceps",
			Image:       "straight_leg_raise.jpg",
			Video:       "straight_leg_raise.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Heel Slide",
			Description: "Geser tumit mendekati bokong untuk meningkatkan fleksi lutut.",
			Muscle:      "Hamstring",
			Image:       "heel_slide.jpg",
			Video:       "heel_slide.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Seated Knee Extension",
			Description: "Duduk di kursi, luruskan lutut hingga sejajar, tahan sebentar lalu turunkan.",
			Muscle:      "Quadriceps",
			Image:       "seated_knee_extension.jpg",
			Video:       "seated_knee_extension.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Side-Lying Hip Abduction",
			Description: "Berbaring menyamping, angkat kaki atas ke atas, lalu turunkan perlahan.",
			Muscle:      "Gluteus Medius",
			Image:       "side_lying_hip_abduction.jpg",
			Video:       "side_lying_hip_abduction.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Mini Squat",
			Description: "Berdiri dengan kaki selebar bahu, tekuk lutut 30°, lalu kembali berdiri.",
			Muscle:      "Quadriceps, Gluteus",
			Image:       "mini_squat.jpg",
			Video:       "mini_squat.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Wall Sit",
			Description: "Sandarkan punggung ke dinding, tekuk lutut 90°, tahan posisi beberapa detik.",
			Muscle:      "Quadriceps, Gluteus",
			Image:       "wall_sit.jpg",
			Video:       "wall_sit.mp4",
			TypeID:      &typeId2,
		},
		{
			Name:        "Bridging",
			Description: "Berbaring telentang, tekuk lutut, angkat panggul ke atas, tahan beberapa detik.",
			Muscle:      "Gluteus, Hamstring",
			Image:       "bridging.jpg",
			Video:       "bridging.mp4",
			TypeID:      &typeId2,
		},
		{
			Name:        "Plank",
			Description: "Tahan posisi plank dengan siku di lantai untuk melatih stabilitas core.",
			Muscle:      "Core, Gluteus",
			Image:       "plank.jpg",
			Video:       "plank.mp4",
			TypeID:      &typeId2,
		},
		{
			Name:        "Single Leg Balance",
			Description: "Berdiri di satu kaki, jaga keseimbangan selama beberapa detik.",
			Muscle:      "Quadriceps, Calves",
			Image:       "single_leg_balance.jpg",
			Video:       "single_leg_balance.mp4",
			TypeID:      &typeId2,
		},
		{
			Name:        "Step-Up Hold",
			Description: "Naik ke tangga kecil, tahan posisi selama beberapa detik, lalu turun perlahan.",
			Muscle:      "Quadriceps, Gluteus",
			Image:       "step_up_hold.jpg",
			Video:       "step_up_hold.mp4",
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
