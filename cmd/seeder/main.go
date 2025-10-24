package seeder

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/cmd/seeder/seed"
	"ReMotion-C7/config"
)

func main() {

	config.ConnectDatabase()

	database := config.GetDatabase()

	database.AutoMigrate(&model.Role{})
	seed.SeedRoles(database)

	database.AutoMigrate(&model.Gender{})
	seed.SeedGenders(database)

	database.AutoMigrate(&model.Type{})
	seed.SeedTypes(database)

	database.AutoMigrate(&model.User{})
	seed.SeedUsers(database)

	database.AutoMigrate(&model.Exercise{})
	seed.SeedExercises(database)

	database.AutoMigrate(&model.Patient{})
	seed.SeedPatients(database)

	database.AutoMigrate(&model.Symptom{})
	seed.SeedSymptoms(database)

	database.AutoMigrate(&model.PatientExercise{})
	seed.SeedPatientExercises(database)
	
}
