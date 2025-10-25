package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/helper"
	"ReMotion-C7/constant"
	"ReMotion-C7/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateExerciseService(c *fiber.Ctx, createExerciseDto request.CreateEditExerciseDto) error {

	imageUrl, err := utils.UploadFileToStorage(c, "image")
	if err != nil {
		return err
	}
	videoUrl, err := utils.UploadFileToStorage(c, "video")
	if err != nil {
		return err
	}

	err = helper.AddExercise(createExerciseDto, imageUrl, videoUrl)
	if err != nil {
		return err
	}

	return nil

}

func AssignExerciseService(assignExerciseDto request.AssignExerciseToPatientDto, patientId int) error {

	err := helper.AddExerciseToPatient(assignExerciseDto, patientId)
	if err != nil {
		return err
	}

	return nil

}

func EditPatientExerciseService(editPatientExerciseDto request.EditPatientExerciseDto, patientId int, exerciseId int) error {

	err := helper.EditPatientExercise(editPatientExerciseDto, patientId, exerciseId)
	if err != nil {
		return err
	}

	return nil

}

func GetExercisesService(c *fiber.Ctx) ([]response.ExerciseDto, error) {

	exercises, err := helper.RetrieveExercises()
	if err != nil {
		return []response.ExerciseDto{}, err
	}

	return exercises, nil

}

func GetExerciseDetail(c *fiber.Ctx, id int) (response.ExerciseDetailForFisioDto, error) {

	exercise, err := helper.FindExercise(id)
	if err != nil {
		return response.ExerciseDetailForFisioDto{}, err
	}

	if exercise.Id == 0 {
		return response.ExerciseDetailForFisioDto{}, fmt.Errorf(constant.ErrExerciseNotFound)
	}

	return exercise, nil

}
