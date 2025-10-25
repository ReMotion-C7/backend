package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/repository"
)

func AssignExerciseService(dto request.AssignExerciseToPatientDto, patientId int) error {

	err := repository.AddExerciseToPatient(dto, patientId)
	if err != nil {
		return err
	}

	return nil

}

func EditPatientExerciseService(dto request.EditPatientExerciseDto, patientId int, exerciseId int) error {

	err := repository.EditPatientExercise(dto, patientId, exerciseId)
	if err != nil {
		return err
	}

	return nil

}

func DeletePatientExerciseService(patientId int, exerciseId int) error {

	err := repository.DeletePatientExercise(patientId, exerciseId)
	if err != nil {
		return err
	}

	return nil
}
