package model

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}
