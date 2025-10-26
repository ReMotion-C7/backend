package controller

import (
	"ReMotion-C7/app/dto/request"
	"ReMotion-C7/app/service"
	"ReMotion-C7/constant"
	"ReMotion-C7/output"
	"ReMotion-C7/utils"

	"github.com/gofiber/fiber/v2"
)

func AddPatient(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var dto request.AddPatientDto
	err = ParseAndValidateBody(c, &dto)
	if err != nil {
		return err
	}

	err = service.AddPatientService(c, dto, id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAddPatient, nil)

}

func EditPatient(c *fiber.Ctx) error {

	fisioId, patientId, err := utils.ConvertToNum2Var(c, "id", "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var dto request.EditPatientDto
	err = ParseAndValidateBody(c, &dto)
	if err != nil {
		return err
	}

	err = service.EditPatientService(c, dto, fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditPatient, nil)

}

func GetPatients(c *fiber.Ctx) error {

	patientName := c.Query("name")

	patients, err := service.GetPatientsService(c, patientName)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatients, patients)

}

func GetPatientPhase(c *fiber.Ctx) error {

	id, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	phase, err := service.GetPatientPhaseService(id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatients, phase)

}

func GetPatientDetail(c *fiber.Ctx) error {

	fisioId, patientId, err := utils.ConvertToNum2Var(c, "id", "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patient, err := service.GetPatientDetail(c, fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatientDetail, patient)

}
