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

	var dto request.CreateEditExerciseDto
	err := ParseAndValidateBody(c, &dto)
	if err != nil {
		return err
	}

	err = service.CreateExerciseService(c, dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessCreateExercise, nil)

}

func GetExercises(c *fiber.Ctx) error {

	exercises, err := service.GetExercisesService(c, constant.ExerciseDefault)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchExercises, exercises)

}

func GetExercisesModal(c *fiber.Ctx) error {

	exercises, err := service.GetExercisesService(c, constant.ExerciseModal)
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

	exercise, err := service.GetExerciseDetailService(id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)

	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchExerciseDetail, exercise)

}

func DeleteExercise(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	err = service.DeleteExerciseService(id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)

	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeleteExercise, nil)

}
