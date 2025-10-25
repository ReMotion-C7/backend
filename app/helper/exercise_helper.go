package helper

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
)

func AddExercise(createExerciseDto request.CreateEditExerciseDto, imageUrl string, videoUrl string) error {

	database := config.GetDatabase()

	exercise := model.Exercise{
		Name: createExerciseDto.Name,
		TypeID: uint(createExerciseDto.TypeId),
		Description: createExerciseDto.Description,
		Muscle: createExerciseDto.Muscle,
		Image: imageUrl,
		Video: videoUrl,
	}

	err := database.Create(&exercise).Error
	if err != nil {
		return err
	}

	return nil

}