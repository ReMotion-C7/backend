package repository

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
)

func AddExercise(dto request.CreateEditExerciseDto, imageUrl string, videoUrl string) error {

	database := config.GetDatabase()

	exercise := model.Exercise{
		Name:        dto.Name,
		TypeID:      uint(dto.TypeId),
		Description: dto.Description,
		Muscle:      dto.Muscle,
		Image:       imageUrl,
		Video:       videoUrl,
	}

	err := database.Create(&exercise).Error
	if err != nil {
		return err
	}

	return nil

}

func RetrieveExercises(mode int) (interface{}, error) {

	database := config.GetDatabase()

	var exercises []model.Exercise

	err := database.Preload("Type").Find(&exercises).Error
	if err != nil {
		return []response.ExerciseDto{}, err
	}

	if mode == constant.ExerciseDefault {

		var dto []response.ExerciseDto
		for _, e := range exercises {
			dto = append(dto, response.ExerciseDto{
				Id:          int(e.ID),
				Name:        e.Name,
				Type:        e.Type.Name,
				Description: e.Description,
				Muscle:      e.Muscle,
				Image:       e.Image,
			})
		}
		return dto, nil

	}

	var dto []response.ExerciseModalDto
	for _, e := range exercises {
		dto = append(dto, response.ExerciseModalDto{
			Id:     int(e.ID),
			Name:   e.Name,
			Type:   e.Type.Name,
			Muscle: e.Muscle,
			Image:  e.Image,
		})
	}
	return dto, nil

}

func RetrieveExercisesModal() ([]response.ExerciseDto, error) {

	database := config.GetDatabase()

	var exercises []model.Exercise

	err := database.Preload("Type").Find(&exercises).Error
	if err != nil {
		return []response.ExerciseDto{}, err
	}

	var dto []response.ExerciseDto
	for _, e := range exercises {
		dto = append(dto, response.ExerciseDto{
			Id:          int(e.ID),
			Name:        e.Name,
			Type:        e.Type.Name,
			Description: e.Description,
			Muscle:      e.Muscle,
			Image:       e.Image,
		})
	}

	return dto, nil

}

func FindExercise(id int) (response.ExerciseDetailForFisioDto, error) {

	database := config.GetDatabase()
	var exercise model.Exercise

	err := database.Preload("Type").Where("id = ?", id).Find(&exercise).Error
	if err != nil {
		return response.ExerciseDetailForFisioDto{}, err
	}

	return response.ExerciseDetailForFisioDto{
		Id:          int(exercise.ID),
		Name:        exercise.Name,
		Type:        exercise.Type.Name,
		Description: exercise.Description,
		Muscle:      exercise.Muscle,
		Image:       exercise.Image,
		Video:       exercise.Video,
	}, nil

}
