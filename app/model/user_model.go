package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	PhoneNumber string    `gorm:"column:phoneNumber"`
	Password    string    `gorm:"column:password"`
	DateOfBirth time.Time `gorm:"column:dateOfBirth"`
	RoleID      int       `gorm:"column:roleId"`
	Role        Role      `gorm:"foreignKey:RoleID"`
	GenderID    int       `gorm:"column:genderId"`
	Gender      Gender    `gorm:"foreignKey:GenderID"`
}
