package controller

import (
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func ParseAndValidateBody(c *fiber.Ctx, target interface{}) error {

	if err := c.BodyParser(target); err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	if err := utils.GetValidator().Struct(target); err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	return nil

}
