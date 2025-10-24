package model

import "gorm.io/gorm"

type Symptom struct {
	gorm.Model
	Name      string  `gorm:"column:name"`
	PatientID uint    `gorm:"column:patientId"`
	Patient   Patient `gorm:"foreignKey:PatientID"`
}
