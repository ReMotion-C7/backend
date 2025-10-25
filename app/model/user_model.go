package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Password    string    `gorm:"column:password"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	RoleID      uint      `gorm:"column:role_id"`
	Role        Role      `gorm:"foreignKey:RoleID"`
	GenderID    uint      `gorm:"column:gender_id"`
	Gender      Gender    `gorm:"foreignKey:GenderID"`
}
