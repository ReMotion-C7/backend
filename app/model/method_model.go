package model

import "gorm.io/gorm"

type Method struct {
	gorm.Model
	Name string `gorm:"column:name"`
}
