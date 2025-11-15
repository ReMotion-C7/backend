package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/dto/response"
	"ReMotion-C7/app/repository"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExerciseService(c *fiber.Ctx, dto request.CreateEditExerciseDto) error {

	imageUrl, err := utils.UploadFileToStorage(c, "image")
	if err != nil {
		return err
	}
	videoUrl, err := utils.UploadFileToStorage(c, "video")
	if err != nil {
		return err
	}

	err = repository.AddExercise(dto, imageUrl, videoUrl)
	if err != nil {
		return err
	}

	return nil

}

func GetExercisesService(c *fiber.Ctx, mode int) (interface{}, error) {

	// name := c.Query("name")

	// if name == "" {

	// }

	// exercises, err := repository.FindExercisesByName(mode, name)
	// if err != nil {
	// 	return nil, err
	// }

	exercises, err := repository.RetrieveExercises(mode)
	if err != nil {
		return nil, err
	}

	return exercises, nil

}

func GetExerciseDetailService(id int) (response.ExerciseDetailForFisioDto, error) {

	exercise, err := repository.FindExercise(id)
	if err != nil {
		return response.ExerciseDetailForFisioDto{}, err
	}

	return exercise, nil

}

func DeleteExerciseService(id int) error {

	err := repository.DeleteExercise(id)
	if err != nil {
		return err
	}

	return nil

}
