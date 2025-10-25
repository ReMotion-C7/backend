package repository

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/model"
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func AddExerciseToPatient(dto request.AssignExerciseToPatientDto, patientId int) error {

	database := config.GetDatabase()

	patientExercise := model.PatientExercise{
		PatientID:  uint(patientId),
		ExerciseID: uint(dto.ExerciseId),
		Set:        dto.Set,
		RepOrTime:  dto.RepOrTime,
	}

	err := database.Create(&patientExercise).Error
	if err != nil {
		return err
	}

	return nil

}

func EditPatientExercise(dto request.EditPatientExerciseDto, patientId int, exerciseId int) error {

	database := config.GetDatabase()

	var patientExercise model.PatientExercise

	if err := database.
		Where("patient_id = ? AND exercise_id = ?", patientId, exerciseId).
		First(&patientExercise).Error; err != nil {
		return fmt.Errorf(constant.ErrPatientExerciseNotFound)
	}

	err := database.Model(&patientExercise).
		Updates(map[string]interface{}{
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
			return fmt.Errorf(constant.ErrPatientExerciseNotFound)
		}
		return err
	}

	err = database.Delete(&patientExercise).Error
	if err != nil {
		return err
	}

	return nil

}
