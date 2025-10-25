package controller

import (
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func GetPatients(c *fiber.Ctx) error {

	patients, err := service.GetPatientsService(c)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatients, patients)

}

func GetPatientDetail(c *fiber.Ctx) error {

	fisioIdStr := c.Params("id")
	fisioId, err := utils.ConvertToNum(fisioIdStr)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patientIdStr := c.Params("patient_id")
	patientId, err := utils.ConvertToNum(patientIdStr)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patient, err := service.GetPatientDetail(c, fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatientDetail, patient)

}
