package model

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Phase            int               `gorm:"column:phase"`
	TherapyStartDate time.Time         `gorm:"column:therapy_start_date"`
	UserID           uint              `gorm:"column:user_id"`
	PatientUser      User              `gorm:"foreignKey:UserID"`
	FisiotherapyID   uint              `gorm:"column:fisiotherapy_id"`
	FisiotherapyUser User              `gorm:"foreignKey:FisiotherapyID"`
	Symptoms         []Symptom         `gorm:"foreignKey:PatientID"`
	PatientExercises []PatientExercise `gorm:"foreignKey:PatientID"`
}
