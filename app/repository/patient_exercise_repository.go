package repository

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func AddExerciseToPatient(dto request.AssignExerciseToPatientDto, id int) error {

	database := config.GetDatabase()

	patientId := uint(id)
	exerciseId := uint(dto.ExerciseId)

	patientExercise := model.PatientExercise{
		PatientID:  patientId,
		ExerciseID: exerciseId,
		MethodID:   utils.PointNumber(dto.MethodId),
		Set:        dto.Set,
		RepOrTime:  dto.RepOrTime,
	}

	err := database.Create(&patientExercise).Error
	if err != nil {
		return err
	}

	return nil

}

func RetrievePatientExercises(mode int, id int) (interface{}, error) {

	database := config.GetDatabase()

	var patientExercises []model.PatientExercise

	err := database.Preload("Method").
		Preload("Exercise").
		Where(`patient_id = ?`, id).Find(&patientExercises).Error

	if err != nil {
		return nil, err
	}

	if mode == constant.PatientExerciseDefault {

		var dto []response.ExerciseForPatientDto
		for _, e := range patientExercises {

			exerciseId := int(e.ExerciseID)

			dto = append(dto, response.ExerciseForPatientDto{
				Id:        exerciseId,
				Name:      e.Exercise.Name,
				Method:    e.Method.Name,
				Muscle:    e.Exercise.Muscle,
				Video:     e.Exercise.Video,
				Set:       e.Set,
				RepOrTime: e.RepOrTime,
			})
		}

		return dto, nil

	}

	var dto []response.PatientSessionDto

	for _, e := range patientExercises {

		exerciseId := int(e.ExerciseID)

		dto = append(dto, response.PatientSessionDto{
			Id:        exerciseId,
			Name:      e.Exercise.Name,
			Method:    e.Method.Name,
			Muscle:    e.Exercise.Muscle,
			Image:     e.Exercise.Image,
			Set:       e.Set,
			RepOrTime: e.RepOrTime,
		})
	}

	return dto, nil

}

func RetrievePatientDetailExercise(patientId int, exerciseId int) (response.ExerciseDetailForPatientDto, error) {

	database := config.GetDatabase()

	var patientExercise model.PatientExercise

	err := database.Preload("Method").
		Preload("Exercise").
		Where(`patient_id = ? AND exercise_id = ?`, patientId, exerciseId).
		First(&patientExercise).Error
	if err != nil {
		return response.ExerciseDetailForPatientDto{}, fmt.Errorf(constant.ErrExerciseNotFound)
	}

	if patientExercise.ID == 0 {
		return response.ExerciseDetailForPatientDto{}, fmt.Errorf(constant.ErrExerciseNotFound)
	}

	return response.ExerciseDetailForPatientDto{
		Id:          exerciseId,
		Name:        patientExercise.Exercise.Name,
		Description: patientExercise.Exercise.Description,
		Method:      patientExercise.Method.Name,
		Muscle:      patientExercise.Exercise.Muscle,
		Video:       patientExercise.Exercise.Video,
		Set:         patientExercise.Set,
		RepOrTime:   patientExercise.RepOrTime,
	}, nil
}

func EditPatientExercise(dto request.EditPatientExerciseDto, patientId int, exerciseId int) error {

	database := config.GetDatabase()

	var patientExercise model.PatientExercise

	if err := database.
		Where("patient_id = ? AND exercise_id = ?", patientId, exerciseId).
		First(&patientExercise).Error; err != nil {
		return fmt.Errorf(constant.ErrExerciseNotFound)
	}

	err := database.Model(&patientExercise).
		Updates(map[string]interface{}{
			"method_id":   dto.MethodId,
			"set":         dto.Set,
			"rep_or_time": dto.RepOrTime,
		}).Error
	if err != nil {
		return err
	}

	return nil

}

func DeletePatientExercise(patientId int, exerciseId int) error {

	database := config.GetDatabase()

	var patientExercise model.PatientExercise

	err := database.Where("patient_id = ? AND exercise_id = ?", patientId, exerciseId).
		First(&patientExercise).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf(constant.ErrExerciseNotFound)
		}
		return err
	}

	err = database.Unscoped().Delete(&patientExercise).Error
	if err != nil {
		return err
	}

	return nil

}
