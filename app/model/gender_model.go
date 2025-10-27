package model

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	Name  string `gorm:"column:name"`
}
