package model

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Phase            int               `gorm:"column:phase"`
	TherapyStartDate time.Time         `gorm:"column:therapyStartDate"`
	UserID           uint              `gorm:"column:userId"`
	User             User              `gorm:"foreignKey:UserID"`
	Symptoms         []Symptom         `gorm:"foreignKey:PatientID"`
	PatientExercises []PatientExercise `gorm:"foreignKey:PatientID"`
}
