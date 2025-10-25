package helper

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"fmt"
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

func EditPatient(editPatientDto request.EditPatientDto, fisioId int, patientId int) error {

	database := config.GetDatabase()

	var patient model.Patient

	err := database.
		Where(`id = ? AND fisiotherapy_id = ?`, patientId, fisioId).
		First(&patient).Error
	if err != nil {
		return fmt.Errorf(constant.ErrPatientNotFound)
	}

	err = database.Model(&patient).Update("phase", editPatientDto.Phase).Error
	if err != nil {
		return err
	}

	err = database.Where(`patient_id = ?`, patientId).Delete(&model.Symptom{}).Error
	if err != nil {
		return fmt.Errorf(constant.ErrPatientNotFound)
	}

	var newSymptoms []model.Symptom
	for _, s := range editPatientDto.Symptoms {
		newSymptoms = append(newSymptoms, model.Symptom{
			Name:      s,
			PatientID: patient.ID,
		})
	}

	if err := database.Create(&newSymptoms).Error; err != nil {
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
