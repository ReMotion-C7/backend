package model

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Muscle      string `gorm:"column:muscle"`
	Image       string `gorm:"column:image"`
	Video       string `gorm:"column:video"`
	TypeID      int    `gorm:"column:typeId"`
	Type        Type   `gorm:"foreignKey:TypeID"`
}
