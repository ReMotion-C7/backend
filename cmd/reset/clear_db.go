package reset

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func ClearDatabase(db *gorm.DB) {

	tables := []interface{}{
		&model.PatientExercise{},
		&model.Symptom{},
		&model.Patient{},
		&model.Exercise{},
		&model.User{},
		&model.Type{},
		&model.Gender{},
		&model.Role{},
	}

	for _, table := range tables {

		err := db.Migrator().DropTable(table)
		if err != nil {
			log.Fatalf(constant.ErrResetTable)
		}

	}

}
