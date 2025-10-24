package model

type Symptom struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}
