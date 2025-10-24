package model

import "gorm.io/gorm"

type PatientExercise struct {
	gorm.Model
	Set        int      `gorm:"column:set"`
	RepOrTime  int      `gorm:"column:repOrTime"`
	PatientID  uint     `gorm:"column:patientId"`
	Patient    Patient  `gorm:"foreignKey:PatientID"`
	ExerciseID uint     `gorm:"column:exerciseId"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseID"`
}
