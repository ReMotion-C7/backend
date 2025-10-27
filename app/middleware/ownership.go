package middleware

import (
	"ReMotion-C7/app/repository"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func PatientOwnership(c *fiber.Ctx) error {

	fisioId, patientId, err := utils.ConvertToNum2Var(c, "id", "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = repository.FindPatientById(fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)

	}

	return c.Next()
}
