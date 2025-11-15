package model

import (
	"time"

	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	Date      time.Time `gorm:"column:date"`
	PatientID uint      `gorm:"column:patient_id"`
	Patient   Patient   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE;"`
}
