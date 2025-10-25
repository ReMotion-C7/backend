package controller

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExercise(c *fiber.Ctx) error {

	var createExerciseDto request.CreateEditExerciseDto
	err := c.BodyParser(&createExerciseDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(createExerciseDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	err = service.CreateExerciseService(c, createExerciseDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessCreateExercise, nil)

}

func AssignExercise(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var assignExerciseToPatientDto request.AssignExerciseToPatientDto
	err = c.BodyParser(&assignExerciseToPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(assignExerciseToPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	err = service.AssignExerciseService(assignExerciseToPatientDto, id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAssignExercise, nil)

}

func EditPatientExercise(c *fiber.Ctx) error {

	patientId, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	exerciseId, err := utils.ConvertToNum(c, "exercise_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var editPatientExerciseDto request.EditPatientExerciseDto
	err = ParseAndValidateBody(c, &editPatientExerciseDto)
	if err != nil {
		return err
	}

	err = service.EditPatientExerciseService(editPatientExerciseDto, patientId, exerciseId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditPatientExercise, nil)

}

func DeletePatientExercise(c *fiber.Ctx) error {

	patientId, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	exerciseId, err := utils.ConvertToNum(c, "exercise_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = service.DeletePatientExerciseService(patientId, exerciseId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeletePatientExercise, nil)
}

func GetExercises(c *fiber.Ctx) error {

	exercises, err := service.GetExercisesService(c)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchExercises, exercises)

}

func GetExerciseDetail(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	exercise, err := service.GetExerciseDetail(c, id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)

	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchExerciseDetail, exercise)

}
