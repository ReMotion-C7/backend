package helper

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
)

func AddExercise(createExerciseDto request.CreateEditExerciseDto, imageUrl string, videoUrl string) error {

	database := config.GetDatabase()

	exercise := model.Exercise{
		Name:        createExerciseDto.Name,
		TypeID:      uint(createExerciseDto.TypeId),
		Description: createExerciseDto.Description,
		Muscle:      createExerciseDto.Muscle,
		Image:       imageUrl,
		Video:       videoUrl,
	}

	err := database.Create(&exercise).Error
	if err != nil {
		return err
	}

	return nil

}

func AddExerciseToPatient(assignExerciseDto request.AssignExerciseToPatientDto, patientId int) error {

	database := config.GetDatabase()

	patientExercise := model.PatientExercise{
		PatientID:  uint(patientId),
		ExerciseID: uint(assignExerciseDto.ExerciseId),
		Set:        assignExerciseDto.Set,
		RepOrTime:  assignExerciseDto.RepOrTime,
	}

	err := database.Create(&patientExercise).Error
	if err != nil {
		return err
	}

	return nil

}

func RetrieveExercises() ([]response.ExerciseDto, error) {

	database := config.GetDatabase()

	var exercises []model.Exercise

	err := database.Preload("Type").Find(&exercises).Error
	if err != nil {
		return []response.ExerciseDto{}, err
	}

	var exercisesDto []response.ExerciseDto
	for _, e := range exercises {
		exercisesDto = append(exercisesDto, response.ExerciseDto{
			Id:          int(e.ID),
			Name:        e.Name,
			Type:        e.Type.Name,
			Description: e.Description,
			Muscle:      e.Muscle,
			Image:       e.Image,
		})
	}

	return exercisesDto, nil

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
