package main

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/cmd/seeder/seed"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"fmt"
	"log"

	"gorm.io/gorm"
)

const (
	Green = "\033[32m"
	Blue  = "\033[34m"
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func main() {

	config.LoadEnv()

	config.ConnectDatabase()
	database := config.GetDatabase()

	fmt.Println(Blue + "🚀 Starting database migration and seeding..." + Reset)

	autoMigrateAndSeed(database, &model.Category{}, "Category", func() { seed.SeedCategories(database) })
	autoMigrateAndSeed(database, &model.Method{}, "Method", func() { seed.SeedMethods(database) })
	autoMigrateAndSeed(database, &model.Phase{}, "Phase", func() { seed.SeedPhases(database) })
	autoMigrateAndSeed(database, &model.Role{}, "Role", func() { seed.SeedRoles(database) })
	autoMigrateAndSeed(database, &model.Gender{}, "Gender", func() { seed.SeedGenders(database) })
	autoMigrateAndSeed(database, &model.Type{}, "Type", func() { seed.SeedTypes(database) })
	autoMigrateAndSeed(database, &model.User{}, "User", func() { seed.SeedUsers(database) })
	autoMigrateAndSeed(database, &model.Exercise{}, "Exercise", func() { seed.SeedExercises(database) })
	autoMigrateAndSeed(database, &model.Patient{}, "Patient", func() { seed.SeedPatients(database) })
	autoMigrateAndSeed(database, &model.Progress{}, "Progress", func() { seed.SeedProgress(database) })
	autoMigrateAndSeed(database, &model.Symptom{}, "Symptom", func() { seed.SeedSymptoms(database) })
	autoMigrateAndSeed(database, &model.PatientExercise{}, "PatientExercise", func() { seed.SeedPatientExercises(database) })

	fmt.Println(Green + "✅ All tables migrated and seeded successfully!" + Reset)

}

func autoMigrateAndSeed(db *gorm.DB, model interface{}, name string, seedFunc func()) {

	if err := db.AutoMigrate(model); err != nil {
		log.Printf(Red+"%s %s: %v"+Reset, constant.ErrMigrateDatabase, name, err)
		return
	}

	log.Printf(Green+"%s %s"+Reset, constant.SuccessMigrateDatabase, name)

	seedFunc()

	log.Printf(Green+"%s %s"+Reset, constant.SuccessSeedDatabase, name)

}
