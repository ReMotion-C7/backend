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
	PatientUser      User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	FisiotherapyID   uint              `gorm:"column:fisiotherapy_id"`
	FisiotherapyUser User              `gorm:"foreignKey:FisiotherapyID;constraint:OnDelete:SET NULL;"`
	Symptoms         []Symptom         `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE;"`
	PatientExercises []PatientExercise `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE;"`
}
