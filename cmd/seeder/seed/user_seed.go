package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"log"
	"math/rand/v2"
	"time"

	"github.com/bxcodec/faker/v3"
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

func SeedGenders(db *gorm.DB) {

	genders := []model.Gender{
		{Name: "Laki-laki"},
		{Name: "Wanita"},
		{Name: "Tidak ingin memberitahukan"},
	}

	err := db.Create(&genders).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}

func SeedUsers(db *gorm.DB) {

	var roles []model.Role
	var genders []model.Gender

	err := db.Find(&roles).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	err = db.Find(&genders).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	users := []model.User{}

	for i := 0; i < 10; i++ {

		hashPassword, err := utils.HashPassword("password123")
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		dateOfBirthStr := faker.Date()
		date, err := time.Parse("2006-01-02", dateOfBirthStr)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		roleID := uint(1)
		if i >= 5 {
			roleID = 2
		}

		genderId := uint(rand.IntN(2) + 1)

		user := model.User{
			Name:        faker.Name(),
			Email:       faker.Email(),
			Password:    hashPassword,
			PhoneNumber: faker.Phonenumber(),
			DateOfBirth: date,
			RoleID:      &roleID,
			GenderID:    &genderId,
		}

		users = append(users, user)
	}

	err = db.Create(&users).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}

func SeedPatients(db *gorm.DB) {

	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		log.Fatalf(constant.ErrFetchDataWhileSeeding)
	}

	patients := []model.Patient{}

	for i := 0; i < 5; i++ {

		therapyStartDate := faker.Date()
		date, err := time.Parse("2006-01-02", therapyStartDate)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		patient := model.Patient{
			Phase:            (rand.IntN(2) + 1),
			TherapyStartDate: date,
			UserID:           uint(i + 6),
			FisiotherapyID:   uint(i + 1),
		}

		patients = append(patients, patient)

	}

	err = db.Create(&patients).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}

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
