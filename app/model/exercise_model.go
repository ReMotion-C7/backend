package model

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name             string            `gorm:"column:name"`
	Description      string            `gorm:"column:description"`
	Muscle           string            `gorm:"column:muscle"`
	Image            string            `gorm:"column:image"`
	Video            string            `gorm:"column:video"`
	TypeID           *uint             `gorm:"column:type_id"`
	Type             Type              `gorm:"foreignKey:TypeID;constraint:OnDelete:SET NULL;"`
	PatientExercises []PatientExercise `gorm:"foreignKey:ExerciseID;constraint:OnDelete:CASCADE;"`
}
