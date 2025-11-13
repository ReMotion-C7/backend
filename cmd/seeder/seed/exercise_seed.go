package seed

import (
	"ReMotion-C7/app/model"
	"ReMotion-C7/constant"
	"log"

	"gorm.io/gorm"
)

func SeedExercises(db *gorm.DB) {

	typeId1 := uint(1)
	typeId2 := uint(2)

	exercises := []model.Exercise{
		{
			Name:        "Lunges",
			Description: "Latihan kaki untuk memperkuat otot paha dan bokong dengan melangkah maju lalu menekuk lutut.",
			Muscle:      "Otot Paha",
			Image:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Lunges%20Landscape%20Image.png",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Lunges%20Landscape.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "Squat",
			Description: "Latihan dasar kaki yang melatih otot paha, bokong, dan punggung bawah.",
			Muscle:      "Otot Paha",
			Image:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Squat%20Landscape%20Image.png",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/Squat%20Landscape.mp4",
			TypeID:      &typeId1,
		},
		{
			Name:        "One Leg Balance",
			Description: "Latihan keseimbangan dengan berdiri di satu kaki untuk melatih stabilitas dan otot kaki.",
			Muscle:      "Otot Kaki",
			Image:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/One%20Leg%20Balance%20Image.png",
			Video:       "https://tjyoilicubnsdpujursp.supabase.co/storage/v1/object/public/ReMotion/One%20Leg%20Balance%20Landscape.mp4",
			TypeID:      &typeId2,
		},
	}

	err := db.Create(&exercises).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
