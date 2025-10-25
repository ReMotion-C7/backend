package service

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/helper"
	"ReMotion-C7/utils"

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
