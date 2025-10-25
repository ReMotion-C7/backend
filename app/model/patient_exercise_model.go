package model

import "gorm.io/gorm"

type PatientExercise struct {
	gorm.Model
	Set        int      `gorm:"column:set"`
	RepOrTime  int      `gorm:"column:rep_or_time"`
	PatientID  uint     `gorm:"column:patient_id"`
	Patient    Patient  `gorm:"foreignKey:PatientID"`
	ExerciseID uint     `gorm:"column:exercise_id"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseID"`
}
