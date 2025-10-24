package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	PhoneNumber string    `gorm:"column:phoneNumber"`
	Password    string    `gorm:"column:password"`
	DateOfBirth time.Time `gorm:"column:dateOfBirth"`
	RoleID      uint      `gorm:"column:roleId"`
	Role        Role      `gorm:"foreignKey:RoleID"`
	GenderID    uint      `gorm:"column:genderId"`
	Gender      Gender    `gorm:"foreignKey:GenderID"`
}
