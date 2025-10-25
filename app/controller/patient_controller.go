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

	var addPatientDto request.AddPatientDto
	err = c.BodyParser(&addPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(addPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	err = service.AddPatientService(c, addPatientDto, id)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAddPatient, nil)

}

func EditPatient(c *fiber.Ctx) error {

	fisioId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patientId, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	var editPatientDto request.EditPatientDto
	err = c.BodyParser(&editPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	err = utils.GetValidator().Struct(editPatientDto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusBadRequest, constant.ErrAllInputMustBeFilled, nil)
	}

	err = service.EditPatientService(c, editPatientDto, fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditPatient, nil)

}

func GetPatients(c *fiber.Ctx) error {

	patientName := c.Query("name")

	if patientName == "" {
		patients, err := service.GetPatientsService(c)
		if err != nil {
			return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
		}
		return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatients, patients)
	}

	patients, err := service.SearchPatientService(c, patientName)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessSearchPatients, patients)

}

func GetPatientDetail(c *fiber.Ctx) error {

	fisioId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patientId, err := utils.ConvertToNum(c, "patient_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	patient, err := service.GetPatientDetail(c, fisioId, patientId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchPatientDetail, patient)

}
