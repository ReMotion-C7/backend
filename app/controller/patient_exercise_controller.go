package controller

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func AssignExercise(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var dto request.AssignExerciseToPatientDto
	err = ParseAndValidateBody(c, &dto)
	if err != nil {
		return err
	}

	err = service.AssignExerciseService(dto, id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAssignExercise, nil)

}

func GetPatientExercises(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	exercises, err := service.GetPatientExercisesService(id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchExercises, exercises)

}

func EditPatientExercise(c *fiber.Ctx) error {

	patientId, exerciseId, err := utils.ConvertToNum2Var(c, "patient_id", "exercise_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var dto request.EditPatientExerciseDto
	err = ParseAndValidateBody(c, &dto)
	if err != nil {
		return err
	}

	err = service.EditPatientExerciseService(dto, patientId, exerciseId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditPatientExercise, nil)

}

func DeletePatientExercise(c *fiber.Ctx) error {

	patientId, exerciseId, err := utils.ConvertToNum2Var(c, "patient_id", "exercise_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = service.DeletePatientExerciseService(patientId, exerciseId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeletePatientExercise, nil)
}
