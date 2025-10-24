package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}
