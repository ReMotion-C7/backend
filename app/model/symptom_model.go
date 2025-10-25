package model

import "gorm.io/gorm"

type Symptom struct {
	gorm.Model
	Name      string  `gorm:"column:name"`
	PatientID uint    `gorm:"column:patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID"`
}
