package model

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}
