package controller

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"
	"fmt"

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
		fmt.Println(err)
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessCreateExercise, nil)

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
