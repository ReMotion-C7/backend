package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"log"

	"gorm.io/gorm"
)

func SeedExercises(db *gorm.DB) {

	exercises := []model.Exercise{
		{
			Name:        "Lunges",
			Description: "Latihan kaki untuk memperkuat otot paha dan bokong dengan melangkah maju lalu menekuk lutut.",
			Muscle:      "Otot Paha",
			Image:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/Lunges%20Landscape%20Image.png",
			Video:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/Lunges%20Landscape.mp4",
			TypeID:      utils.PointNumber(2),
			CategoryID:  utils.PointNumber(2),
		},
		{
			Name:        "Squat",
			Description: "Latihan dasar kaki yang melatih otot paha, bokong, dan punggung bawah.",
			Muscle:      "Otot Paha",
			Image:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/Squat%20Landscape%20Image.png",
			Video:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/Squat%20Landscape.mp4",
			TypeID:      utils.PointNumber(2),
			CategoryID:  utils.PointNumber(2),
		},
		{
			Name:        "One Leg Balance",
			Description: "Latihan keseimbangan dengan berdiri di satu kaki untuk melatih stabilitas dan otot kaki.",
			Muscle:      "Otot Kaki",
			Image:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/One%20Leg%20Balance%20Image.png",
			Video:       "https://uajktgblgndryqoprbdd.supabase.co/storage/v1/object/public/ReMotion/One%20Leg%20Balance%20Landscape.mp4",
			TypeID:      utils.PointNumber(2),
			CategoryID:  utils.PointNumber(1),
		},
	}

	err := db.Create(&exercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
