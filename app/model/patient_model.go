package model

type Patient struct {
	ID     uint `gorm:"primaryKey"`
	UserID int  `gorm:"column:userId"`
	User   User `gorm:"foreignKey:UserID"`
}
