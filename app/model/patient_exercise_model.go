package model

import "gorm.io/gorm"

type PatientExercise struct {
	gorm.Model
	Set        int      `gorm:"column:set"`
	RepOrTime  int      `gorm:"column:rep_or_time"`
	MethodID   *uint    `gorm:"column:method_id"`
	Method     Method   `gorm:"foreignKey:MethodID;constraint:OnDelete:SET NULL;"`
	PatientID  uint     `gorm:"column:patient_id"`
	Patient    Patient  `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE;"`
	ExerciseID uint     `gorm:"column:exercise_id"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseID;constraint:OnDelete:CASCADE;"`
}
